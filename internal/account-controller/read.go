package accountcontroller

import (
	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

type ReadInput struct {
	Email  string `json:"email"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

func (ctrl *Controller) Read(c *fiber.Ctx) error {
	var input ReadInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	ops := model.ReadAccountOptions{
		Email: input.Email,
	}
	count, err := ctrl.AccountService.Count(
		c.Context(),
		ops,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	rows, err := ctrl.AccountService.Read(
		c.Context(),
		ops,
		model.PaginationOptions{
			Limit:  input.Limit,
			Offset: input.Offset,
		},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(model.PaginatedOutput[model.Account]{
		Rows:  rows,
		Count: count,
	})
}
