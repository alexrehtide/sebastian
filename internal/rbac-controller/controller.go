package rbaccontroller

import "github.com/gofiber/fiber/v2"

type RBACService interface{}

func New(rbacService RBACService) *Controller {
	return &Controller{RBACService: rbacService}
}

type Controller struct {
	RBACService RBACService
}

func (ctrl *Controller) AddAccountRole(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).SendString("Unimplemented")
}

func (ctrl *Controller) ReadAccountRoles(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).SendString("Unimplemented")
}

func (ctrl *Controller) RemoveAccountRole(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).SendString("Unimplemented")
}
