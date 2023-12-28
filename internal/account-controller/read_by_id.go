package accountcontroller

import (
	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

type ReadByIDInput struct {
	ID uint `json:"id"`
}

type ReadByIDOutput struct {
	model.Account
}

func (ctrl *Controller) ReadByID(c *fiber.Ctx) error {
	var input ReadByIDInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	output, err := ctrl.AccountService.ReadByID(c.UserContext(), input.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(ReadByIDOutput{output})
}
