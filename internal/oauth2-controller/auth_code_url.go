package oauth2controller

import (
	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

type AuthenticateInput struct {
	Platform model.Platform `json:"platform"`
}

type AuthenticateOutput struct {
	URL string `json:"url"`
}

func (ctrl *Controller) AuthCodeURL(c *fiber.Ctx) error {
	var input AuthenticateInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	url, err := ctrl.Oauth2Service.AuthCodeURL(input.Platform)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(AuthenticateOutput{url})
}
