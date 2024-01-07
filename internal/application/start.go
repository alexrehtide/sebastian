package application

import (
	"context"
	"fmt"

	accountcontroller "github.com/alexrehtide/sebastian/internal/account-controller"
	accountprovider "github.com/alexrehtide/sebastian/internal/account-provider"
	accountrolestorage "github.com/alexrehtide/sebastian/internal/account-role-storage"
	accountservice "github.com/alexrehtide/sebastian/internal/account-service"
	accountstorage "github.com/alexrehtide/sebastian/internal/account-storage"
	authcontroller "github.com/alexrehtide/sebastian/internal/auth-controller"
	authmiddleware "github.com/alexrehtide/sebastian/internal/auth-middleware"
	authservice "github.com/alexrehtide/sebastian/internal/auth-service"
	logstorage "github.com/alexrehtide/sebastian/internal/log-storage"
	loginattemptservice "github.com/alexrehtide/sebastian/internal/login-attempt-service"
	loginattemptstorage "github.com/alexrehtide/sebastian/internal/login-attempt-storage"
	logrushooker "github.com/alexrehtide/sebastian/internal/logrus-hooker"
	mailservice "github.com/alexrehtide/sebastian/internal/mail-service"
	passwordresettingcontroller "github.com/alexrehtide/sebastian/internal/password-resetting-controller"
	passwordresettingservice "github.com/alexrehtide/sebastian/internal/password-resetting-service"
	passwordresettingstorage "github.com/alexrehtide/sebastian/internal/password-resetting-storage"
	rbaccontroller "github.com/alexrehtide/sebastian/internal/rbac-controller"
	rbacmiddleware "github.com/alexrehtide/sebastian/internal/rbac-middleware"
	rbacservice "github.com/alexrehtide/sebastian/internal/rbac-service"
	registrationcontroller "github.com/alexrehtide/sebastian/internal/registration-controller"
	registrationservice "github.com/alexrehtide/sebastian/internal/registration-service"
	registrationstorage "github.com/alexrehtide/sebastian/internal/registration-storage"
	remoteaccountcontroller "github.com/alexrehtide/sebastian/internal/remote-account-controller"
	remoteaccountservice "github.com/alexrehtide/sebastian/internal/remote-account-service"
	remoteaccountstorage "github.com/alexrehtide/sebastian/internal/remote-account-storage"
	sessionprovider "github.com/alexrehtide/sebastian/internal/session-provider"
	sessionservice "github.com/alexrehtide/sebastian/internal/session-service"
	sessionstorage "github.com/alexrehtide/sebastian/internal/session-storage"
	totpcontroller "github.com/alexrehtide/sebastian/internal/totp-controller"
	totpservice "github.com/alexrehtide/sebastian/internal/totp-service"
	"github.com/alexrehtide/sebastian/model"
	"github.com/alexrehtide/sebastian/pkg/postgres"
	"github.com/alexrehtide/sebastian/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func (a *Application) Start(ctx context.Context) error {
	err := a.ConfigService.Load()
	if err != nil {
		return fmt.Errorf("httpserver.Server.Listen: %w", err)
	}

	log := logrus.New()
	log.SetFormatter(new(logrus.TextFormatter))

	if err != nil {
		return fmt.Errorf("httpserver.Server.Listen: %w", err)
	}

	sqlDB, err := postgres.New(postgres.PostgresOptions{
		User:     a.ConfigService.PostgresUser(),
		Password: a.ConfigService.PostgresPassword(),
		Host:     a.ConfigService.PostgresHost(),
		Port:     a.ConfigService.PostgresPort(),
		DBName:   a.ConfigService.PostgresDBName(),
	})
	if err != nil {
		return fmt.Errorf("httpserver.Server.Listen: %w", err)
	}
	defer sqlDB.Close()

	sqlxDB := sqlx.NewDb(sqlDB, "postgres")
	accountRoleStorage := accountrolestorage.New(sqlxDB)
	accountStorage := accountstorage.New(sqlxDB)
	loginAttemptStorage := loginattemptstorage.New(sqlxDB)
	logStorage := logstorage.New(sqlxDB)
	passwordResettingStorage := passwordresettingstorage.New(sqlxDB)
	registrationFormStorage := registrationstorage.New(sqlxDB)
	remoteAccountStorage := remoteaccountstorage.New(sqlxDB)
	sessionStorage := sessionstorage.New(sqlxDB)

	validate := validator.New()

	accountService := accountservice.New(accountStorage, log, validate)
	sessionService := sessionservice.New(log, sessionStorage, validate)
	authService := authservice.New(accountService, sessionService, validate)
	loginAttemptService := loginattemptservice.New(loginAttemptStorage)
	mailService := mailservice.New(a.ConfigService)
	passwordResettingService := passwordresettingservice.New(accountService, passwordResettingStorage)
	rbacService := rbacservice.New(accountRoleStorage, validate)
	registrationFormService := registrationservice.New(accountService, rbacService, registrationFormStorage)
	remoteAccountService := remoteaccountservice.New(remoteAccountStorage, "random state") // TODO: change state
	totpService := totpservice.New()

	accountProvider := accountprovider.New()
	sessionProvider := sessionprovider.New()

	authMiddleware := authmiddleware.New(accountProvider, accountService, sessionProvider, sessionService)
	rbacMiddleware := rbacmiddleware.New(accountProvider, rbacService)
	accountController := accountcontroller.New(accountService)
	authController := authcontroller.New(accountProvider, authService, loginAttemptService, rbacService)
	passwordResettingController := passwordresettingcontroller.New(mailService, passwordResettingService)
	rbacController := rbaccontroller.New(rbacService)
	registrationController := registrationcontroller.New(mailService, registrationFormService, sessionService)
	remoteAccountController := remoteaccountcontroller.New(accountService, remoteAccountService, rbacService, sessionService)
	totpController := totpcontroller.New(accountProvider, totpService)

	log.Hooks.Add(logrushooker.New(logStorage, sessionProvider))

	app := fiber.New()
	app.Use(authMiddleware.Authorize)

	wp := func(permission model.Permission) fiber.Handler {
		return rbacMiddleware.WithPermission(permission)
	}

	{
		g := app.Group("/account")
		g.Post("/create", wp(model.AccountCreate), accountController.Create)
		g.Post("/delete", wp(model.AccountDelete), accountController.Delete)
		g.Post("/read", wp(model.AccountRead), accountController.Read)
		g.Post("/read_by_id", wp(model.AccountRead), accountController.ReadByID)
		g.Post("/update", wp(model.AccountUpdate), accountController.Update)
	}

	{
		g := app.Group("/auth")
		g.Post("/authenticate", authController.Authenticate)
		g.Post("/authorize", wp(model.AuthAuthorize), authController.Authorize)
		g.Post("/logout", wp(model.AuthLogout), authController.Logout)
		g.Post("/refresh", wp(model.AuthRefresh), authController.Refresh)
	}

	{
		g := app.Group("/password_resetting")
		g.Post("/begin", passwordResettingController.Begin)
		g.Post("/end", passwordResettingController.End)
	}

	{
		g := app.Group("/rbac")
		g.Post("/add_account_role", wp(model.RBACAddAccountRole), rbacController.AddAccountRole)
		g.Post("/read_account_roles", wp(model.RBACReadAccountRoles), rbacController.ReadAccountRoles)
		g.Post("/remove_account_role", wp(model.RBACRemoveAccountRole), rbacController.RemoveAccountRole)
	}

	{
		g := app.Group("/registration")
		g.Post("/begin", registrationController.Begin)
		g.Post("/end", registrationController.End)
	}

	{
		g := app.Group("/remote_account")
		g.Post("/auth_code_url", remoteAccountController.AuthCodeURL)
		g.Post("/authenticate", remoteAccountController.Authenticate)
	}

	{
		g := app.Group("/totp")
		g.Post("/generate", wp(model.TOTPGenerate), totpController.Generate)
		g.Post("/validate", wp(model.TOTPValidate), totpController.Validate)
	}

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return app.Listen(a.ConfigService.HTTPServerAddr())
	})
	g.Go(func() error {
		<-gCtx.Done()
		return app.Shutdown()
	})

	return g.Wait()
}
