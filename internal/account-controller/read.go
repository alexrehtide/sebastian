package accountcontroller

import (
	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

type ReadInput struct {
	Email  string `query:"email"`
	Limit  int    `query:"limit"`
	Offset int    `query:"offset"`
}

func (ctrl *Controller) Read(c *fiber.Ctx) error {
	var input ReadInput
	if err := c.QueryParser(&input); err != nil {
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
