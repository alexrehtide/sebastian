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
	sessionservice "github.com/alexrehtide/sebastian/internal/session-service"
	sessionstorage "github.com/alexrehtide/sebastian/internal/session-storage"
	"github.com/alexrehtide/sebastian/model"
	"github.com/alexrehtide/sebastian/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type AuthMiddleware interface {
	Authorize(c *fiber.Ctx) error
}

type RBACMiddleware interface {
	WithPermission(permission model.Permission) fiber.Handler
}

type AccountController interface {
	Create(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Read(c *fiber.Ctx) error
	ReadByID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
}

type AuthController interface {
	Authenticate(c *fiber.Ctx) error
	Authorize(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	Refresh(c *fiber.Ctx) error
}

type RBACController interface {
	AddAccountRole(c *fiber.Ctx) error
	ReadAccountRoles(c *fiber.Ctx) error
	RemoveAccountRole(c *fiber.Ctx) error
}

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

	var authMiddleware AuthMiddleware = authmiddleware.New(accountProvider)
	var rbacMiddleware RBACMiddleware = rbacmiddleware.New(accountProvider, rbacService)
	var accountController AccountController = accountcontroller.New(accountService)
	var authController AuthController = authcontroller.New(accountProvider, authService)
	var rbacController RBACController = rbaccontroller.New(rbacService)

	app := fiber.New()
	app.Use(authMiddleware.Authorize)

	accountRoute := app.Group("/account")
	accountRoute.Post("/", rbacMiddleware.WithPermission(model.AccountCreate), accountController.Create)
	accountRoute.Delete("/:id", rbacMiddleware.WithPermission(model.AccountDelete), accountController.Delete)
	accountRoute.Get("/", rbacMiddleware.WithPermission(model.AccountRead), accountController.Read)
	accountRoute.Get("/:id", rbacMiddleware.WithPermission(model.AccountRead), accountController.ReadByID)
	accountRoute.Put("/:id", rbacMiddleware.WithPermission(model.AccountUpdate), accountController.Update)

	authRoute := app.Group("/auth")
	authRoute.Post("/authenticate", authController.Authenticate)
	authRoute.Get("/authorize", rbacMiddleware.WithPermission(model.AuthAuthorize), authController.Authorize)
	authRoute.Get("/logout", rbacMiddleware.WithPermission(model.AuthLogout), authController.Logout)
	authRoute.Post("/refresh", rbacMiddleware.WithPermission(model.AuthRefresh), authController.Refresh)

	rbacRoute := app.Group("/rbac")
	rbacRoute.Post("/:account_id/add_role", rbacMiddleware.WithPermission(model.RBACAddAccountRole), rbacController.AddAccountRole)
	rbacRoute.Get("/:account_id/roles", rbacMiddleware.WithPermission(model.RBACReadAccountRoles), rbacController.ReadAccountRoles)
	rbacRoute.Post("/:account_id/remove_role", rbacMiddleware.WithPermission(model.RBACRemoveAccountRole), rbacController.RemoveAccountRole)

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
