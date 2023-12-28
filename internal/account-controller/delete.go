package accountcontroller

import "github.com/gofiber/fiber/v2"

type DeleteInput struct {
	ID uint `json:"id"`
}

func (ctrl *Controller) Delete(c *fiber.Ctx) error {
	var input DeleteInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := ctrl.AccountService.Delete(c.Context(), input.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
