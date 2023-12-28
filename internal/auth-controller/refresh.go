package authcontroller

import "github.com/gofiber/fiber/v2"

func (ctrl *Controller) Refresh(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).SendString("Unimplemented")
}
