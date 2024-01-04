package httpserver

import (
	"database/sql"

	accountcontroller "github.com/alexrehtide/sebastian/internal/account-controller"
	accountprovider "github.com/alexrehtide/sebastian/internal/account-provider"
	accountrolestorage "github.com/alexrehtide/sebastian/internal/account-role-storage"
	accountservice "github.com/alexrehtide/sebastian/internal/account-service"
	accountstorage "github.com/alexrehtide/sebastian/internal/account-storage"
	authcontroller "github.com/alexrehtide/sebastian/internal/auth-controller"
	authmiddleware "github.com/alexrehtide/sebastian/internal/auth-middleware"
	authservice "github.com/alexrehtide/sebastian/internal/auth-service"
	logstorage "github.com/alexrehtide/sebastian/internal/log-storage"
	loginattemptstorage "github.com/alexrehtide/sebastian/internal/login-attempt-storage"
	logrushooker "github.com/alexrehtide/sebastian/internal/logrus-hooker"
	rbaccontroller "github.com/alexrehtide/sebastian/internal/rbac-controller"
	rbacmiddleware "github.com/alexrehtide/sebastian/internal/rbac-middleware"
	rbacservice "github.com/alexrehtide/sebastian/internal/rbac-service"
	sessionprovider "github.com/alexrehtide/sebastian/internal/session-provider"
	sessionservice "github.com/alexrehtide/sebastian/internal/session-service"
	sessionstorage "github.com/alexrehtide/sebastian/internal/session-storage"
	totpcontroller "github.com/alexrehtide/sebastian/internal/totp-controller"
	totpservice "github.com/alexrehtide/sebastian/internal/totp-service"
	"github.com/alexrehtide/sebastian/model"
	"github.com/alexrehtide/sebastian/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func New(sqlDB *sql.DB, log *logrus.Logger) *Server {
	sqlxDB := sqlx.NewDb(sqlDB, "postgres")

	accountRoleStorage := accountrolestorage.New(sqlxDB)
	accountStorage := accountstorage.New(sqlxDB)
	sessionStorage := sessionstorage.New(sqlxDB)
	logStorage := logstorage.New(sqlxDB)

	validate := validator.New()

	accountService := accountservice.New(accountStorage, log, validate)
	loginAttemptStorage := loginattemptstorage.New(sqlxDB)
	sessionService := sessionservice.New(log, sessionStorage, validate)
	authService := authservice.New(accountService, loginAttemptStorage, sessionService, validate)
	rbacService := rbacservice.New(accountRoleStorage, validate)
	totpService := totpservice.New()

	accountProvider := accountprovider.New()
	sessionProvider := sessionprovider.New()

	authMiddleware := authmiddleware.New(accountProvider, accountService, sessionProvider, sessionService)
	rbacMiddleware := rbacmiddleware.New(accountProvider, rbacService)
	accountController := accountcontroller.New(accountService)
	authController := authcontroller.New(accountProvider, authService, rbacService)
	rbacController := rbaccontroller.New(rbacService)
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
		g := app.Group("/rbac")
		g.Post("/add_account_role", wp(model.RBACAddAccountRole), rbacController.AddAccountRole)
		g.Post("/read_account_roles", wp(model.RBACReadAccountRoles), rbacController.ReadAccountRoles)
		g.Post("/remove_account_role", wp(model.RBACRemoveAccountRole), rbacController.RemoveAccountRole)
	}

	{
		g := app.Group("/totp")
		g.Post("/generate", wp(model.TOTPGenerate), totpController.Generate)
		g.Post("/validate", wp(model.TOTPValidate), totpController.Validate)
	}

	return &Server{
		app: app,
	}
}

type Server struct {
	app *fiber.App
}

func (s *Server) Listen(addr string) error {
	return s.app.Listen(addr)
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
