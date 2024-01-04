package totpcontroller

import (
	controllererror "github.com/alexrehtide/sebastian/internal/controller-error"
	"github.com/gofiber/fiber/v2"
)

type GenerateOutput struct {
	URL string `json:"url"`
}

func (ctrl *Controller) Generate(c *fiber.Ctx) error {
	acc := ctrl.AccountProvider.Inject(c.UserContext())
	if acc == nil {
		return c.Status(fiber.StatusForbidden).SendString(controllererror.ErrPermissionDenied.Error())
	}

	url, err := ctrl.TOTPService.Generate(*acc)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(GenerateOutput{url})
}
