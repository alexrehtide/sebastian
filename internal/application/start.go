package application

import (
	"context"
	"fmt"
	"time"

	accountcontroller "github.com/alexrehtide/sebastian/internal/account-controller"
	accountprovider "github.com/alexrehtide/sebastian/internal/account-provider"
	accountrolestorage "github.com/alexrehtide/sebastian/internal/account-role-storage"
	accountservice "github.com/alexrehtide/sebastian/internal/account-service"
	accountstorage "github.com/alexrehtide/sebastian/internal/account-storage"
	authcontroller "github.com/alexrehtide/sebastian/internal/auth-controller"
	authmiddleware "github.com/alexrehtide/sebastian/internal/auth-middleware"
	authservice "github.com/alexrehtide/sebastian/internal/auth-service"
	eventmiddleware "github.com/alexrehtide/sebastian/internal/event-middleware"
	eventservice "github.com/alexrehtide/sebastian/internal/event-service"
	garbageservice "github.com/alexrehtide/sebastian/internal/garbage-service"
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
	"github.com/alexrehtide/sebastian/pkg/validator"

	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func (a *Application) Start(ctx context.Context) error {
	err := a.ConfigService.Load()
	if err != nil {
		return fmt.Errorf("application.Application.Start: %w", err)
	}

	logger := logrus.New()
	logger.SetFormatter(new(logrus.TextFormatter))
	if a.ConfigService.Debug() {
		logger.SetLevel(logrus.DebugLevel)
	}

	sqlDB, err := a.dbConnection()
	if err != nil {
		return fmt.Errorf("application.Application.Start: %w", err)
	}
	defer sqlDB.Close()
	sqlxDB := sqlx.NewDb(sqlDB, "postgres")

	trm, err := manager.New(trmsqlx.NewDefaultFactory(sqlxDB))
	if err != nil {
		return fmt.Errorf("application.Application.Start: %w", err)
	}

	validate := validator.New()

	// storages
	accountRoleStorage := accountrolestorage.New(sqlxDB, trmsqlx.DefaultCtxGetter)
	accountStorage := accountstorage.New(sqlxDB, trmsqlx.DefaultCtxGetter)
	loginAttemptStorage := loginattemptstorage.New(sqlxDB, trmsqlx.DefaultCtxGetter)
	logStorage := logstorage.New(sqlxDB, trmsqlx.DefaultCtxGetter)
	passwordResettingStorage := passwordresettingstorage.New(sqlxDB, trmsqlx.DefaultCtxGetter)
	registrationFormStorage := registrationstorage.New(sqlxDB, trmsqlx.DefaultCtxGetter)
	remoteAccountStorage := remoteaccountstorage.New(sqlxDB, trmsqlx.DefaultCtxGetter)
	sessionStorage := sessionstorage.New(sqlxDB, trmsqlx.DefaultCtxGetter)

	// services
	accountService := accountservice.New(accountStorage, logger, validate)
	eventService := eventservice.New(logger)
	loginAttemptService := loginattemptservice.New(loginAttemptStorage)
	mailService := mailservice.New(a.ConfigService)
	rbacService := rbacservice.New(accountRoleStorage, validate)
	remoteAccountService := remoteaccountservice.New(a.ConfigService, remoteAccountStorage, "random state") // TODO: change state
	sessionService := sessionservice.New(a.ConfigService, logger, sessionStorage, validate)
	totpService := totpservice.New()

	authService := authservice.New(accountService, sessionService, validate)
	passwordResettingService := passwordresettingservice.New(accountService, passwordResettingStorage)
	registrationFormService := registrationservice.New(accountService, rbacService, registrationFormStorage, trm)

	garbageService := garbageservice.New(sessionService)

	// providers
	accountProvider := accountprovider.New()
	sessionProvider := sessionprovider.New()

	// middleware
	authMiddleware := authmiddleware.New(accountProvider, accountService, sessionProvider, sessionService)
	eventMiddleware := eventmiddleware.New(eventService)
	rbacMiddleware := rbacmiddleware.New(accountProvider, rbacService)

	// controllers
	accountController := accountcontroller.New(accountService)
	authController := authcontroller.New(accountProvider, authService, loginAttemptService, rbacService)
	passwordResettingController := passwordresettingcontroller.New(mailService, passwordResettingService)
	rbacController := rbaccontroller.New(rbacService)
	registrationController := registrationcontroller.New(mailService, registrationFormService, sessionService)
	remoteAccountController := remoteaccountcontroller.New(accountService, rbacService, remoteAccountService, sessionService)
	totpController := totpcontroller.New(accountProvider, totpService)

	logger.Hooks.Add(logrushooker.New(logStorage, sessionProvider))

	app := fiber.New()
	app.Use(authMiddleware.Authorize)
	if a.ConfigService.Debug() {
		app.Use(eventMiddleware.Handle)
	}

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
		for {
			select {
			case <-time.After(5 * time.Minute):
				if err := garbageService.Clean(gCtx); err != nil {
					return err
				}
			case <-gCtx.Done():
				return nil
			}
		}
	})
	g.Go(func() error {
		<-gCtx.Done()
		return app.Shutdown()
	})

	return g.Wait()
}
