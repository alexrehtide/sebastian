package accountcontroller

import "github.com/gofiber/fiber/v2"

type ReadByIDInput struct {
	ID uint `params:"id"`
}

func (ctrl *Controller) ReadByID(c *fiber.Ctx) error {
	var input ReadByIDInput
	if err := c.ParamsParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	output, err := ctrl.AccountService.ReadByID(c.Context(), input.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(output)
}
