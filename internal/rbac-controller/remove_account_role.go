package rbaccontroller

import "github.com/gofiber/fiber/v2"

func (ctrl *Controller) RemoveAccountRole(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).SendString("Unimplemented")
}
