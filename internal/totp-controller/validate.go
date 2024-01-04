package totpcontroller

import (
	controllererror "github.com/alexrehtide/sebastian/internal/controller-error"
	"github.com/gofiber/fiber/v2"
)

type ValidateInput struct {
	Code string `json:"code"`
}

func (ctrl *Controller) Validate(c *fiber.Ctx) error {
	acc := ctrl.AccountProvider.Inject(c.UserContext())
	if acc == nil {
		return c.Status(fiber.StatusForbidden).SendString(controllererror.ErrPermissionDenied.Error())
	}

	var input ValidateInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if !ctrl.TOTPService.Validate(*acc, input.Code) {
		return c.SendString("Not Ok")
	}

	return c.SendString("Ok")
}
