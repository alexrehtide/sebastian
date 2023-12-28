package httpserver

import (
	accountcontroller "github.com/alexrehtide/sebastian/internal/account-controller"
	accountprovider "github.com/alexrehtide/sebastian/internal/account-provider"
	accountrolestorage "github.com/alexrehtide/sebastian/internal/account-role-storage"
	accountservice "github.com/alexrehtide/sebastian/internal/account-service"
	accountstorage "github.com/alexrehtide/sebastian/internal/account-storage"
	authcontroller "github.com/alexrehtide/sebastian/internal/auth-controller"
	authmiddleware "github.com/alexrehtide/sebastian/internal/auth-middleware"
	authservice "github.com/alexrehtide/sebastian/internal/auth-service"
	rbaccontroller "github.com/alexrehtide/sebastian/internal/rbac-controller"
	rbacmiddleware "github.com/alexrehtide/sebastian/internal/rbac-middleware"
	rbacservice "github.com/alexrehtide/sebastian/internal/rbac-service"
	sessionprovider "github.com/alexrehtide/sebastian/internal/session-provider"
	sessionservice "github.com/alexrehtide/sebastian/internal/session-service"
	sessionstorage "github.com/alexrehtide/sebastian/internal/session-storage"
	"github.com/alexrehtide/sebastian/model"
	"github.com/alexrehtide/sebastian/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *Server {
	accountRoleStorage := accountrolestorage.New(db)
	accountStorage := accountstorage.New(db)
	sessionStorage := sessionstorage.New(db)

	validate := validator.New()

	accountService := accountservice.New(accountStorage, validate)
	sessionService := sessionservice.New(sessionStorage, validate)
	authService := authservice.New(accountService, sessionService, validate)
	rbacService := rbacservice.New(accountRoleStorage, validate)

	accountProvider := accountprovider.New()
	sessionProvider := sessionprovider.New()

	authMiddleware := authmiddleware.New(accountProvider, accountService, sessionProvider, sessionService)
	rbacMiddleware := rbacmiddleware.New(accountProvider, rbacService)
	accountController := accountcontroller.New(accountService)
	authController := authcontroller.New(accountProvider, authService, rbacService)
	rbacController := rbaccontroller.New(rbacService)

	app := fiber.New()
	app.Use(authMiddleware.Authorize)

	accountRoute := app.Group("/account")
	accountRoute.Post("/create", rbacMiddleware.WithPermission(model.AccountCreate), accountController.Create)
	accountRoute.Post("/delete", rbacMiddleware.WithPermission(model.AccountDelete), accountController.Delete)
	accountRoute.Post("/read", rbacMiddleware.WithPermission(model.AccountRead), accountController.Read)
	accountRoute.Post("/read_by_id", rbacMiddleware.WithPermission(model.AccountRead), accountController.ReadByID)
	accountRoute.Post("/update", rbacMiddleware.WithPermission(model.AccountUpdate), accountController.Update)

	authRoute := app.Group("/auth")
	authRoute.Post("/authenticate", authController.Authenticate)
	authRoute.Post("/authorize", rbacMiddleware.WithPermission(model.AuthAuthorize), authController.Authorize)
	authRoute.Post("/logout", rbacMiddleware.WithPermission(model.AuthLogout), authController.Logout)
	authRoute.Post("/refresh", rbacMiddleware.WithPermission(model.AuthRefresh), authController.Refresh)

	rbacRoute := app.Group("/rbac")
	rbacRoute.Post("/add_account_role", rbacMiddleware.WithPermission(model.RBACAddAccountRole), rbacController.AddAccountRole)
	rbacRoute.Post("/read_account_roles", rbacMiddleware.WithPermission(model.RBACReadAccountRoles), rbacController.ReadAccountRoles)
	rbacRoute.Post("/remove_account_role", rbacMiddleware.WithPermission(model.RBACRemoveAccountRole), rbacController.RemoveAccountRole)

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
