package remoteaccountcontroller

import (
	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

type ExchangeInput struct {
	Platform model.Platform `json:"platform"`
	State    string         `json:"state"`
	Code     string         `json:"code"`
}

type ExchangeOutput struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (ctrl *Controller) Authenticate(c *fiber.Ctx) error {
	var input ExchangeInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	token, err := ctrl.RemoteAccountService.Exchange(c.UserContext(), input.Platform, input.Code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	remoteAcc, err := ctrl.RemoteAccountService.Authorize(c.UserContext(), input.Platform, token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if remoteAcc.AccountID == 0 {
		accId, err := ctrl.AccountService.CreateWithUsername(c.UserContext(), remoteAcc.RemoteEmail)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		err = ctrl.RBACService.AddAccountRole(c.UserContext(), accId, model.User)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		remoteAcc.AccountID = accId
	}

	err = ctrl.RemoteAccountService.UpdateAccountID(c.UserContext(), remoteAcc.ID, remoteAcc.AccountID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	sessionID, err := ctrl.SessionService.CreateWithAccountID(c.UserContext(), remoteAcc.AccountID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	session, err := ctrl.SessionService.ReadByID(c.UserContext(), sessionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(ExchangeOutput{
		AccessToken:  session.AccessToken,
		RefreshToken: session.RefreshToken,
	})
}
