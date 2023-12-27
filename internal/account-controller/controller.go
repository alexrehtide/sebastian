package accountcontroller

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

type AccountService interface {
	Count(ctx context.Context, ops model.ReadAccountOptions) (int, error)
	Create(ctx context.Context, ops model.CreateAccountOptions) (uint, error)
	Delete(ctx context.Context, id uint) error
	Read(ctx context.Context, ops model.ReadAccountOptions, pgOps model.PaginationOptions) ([]model.Account, error)
	ReadByID(ctx context.Context, id uint) (model.Account, error)
	Update(ctx context.Context, id uint, ops model.UpdateAccountOptions) error
}

func New(accountService AccountService) *Controller {
	return &Controller{
		AccountService: accountService,
	}
}

type Controller struct {
	AccountService AccountService
}

type CreateInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ctrl *Controller) Create(c *fiber.Ctx) error {
	var input CreateInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	output, err := ctrl.AccountService.Create(
		c.Context(),
		model.CreateAccountOptions{
			Email:    input.Email,
			Password: input.Password,
		},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(output)
}

type DeleteInput struct {
	ID uint `params:"id"`
}

func (ctrl *Controller) Delete(c *fiber.Ctx) error {
	var input DeleteInput
	if err := c.ParamsParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := ctrl.AccountService.Delete(c.Context(), input.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

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

type UpdateInput struct {
	ID    uint   `params:"id"`
	Email string `json:"email"`
}

func (ctrl *Controller) Update(c *fiber.Ctx) error {
	var input UpdateInput
	if err := c.ParamsParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err := ctrl.AccountService.Update(
		c.Context(),
		input.ID,
		model.UpdateAccountOptions{
			Email: input.Email,
		},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
