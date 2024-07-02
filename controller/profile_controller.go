package controller

import (
	"encoding/json"
	"fmt"

	"github.com/faridlan/client-go/model"
	"github.com/gofiber/fiber/v2"
)

type ProfileController interface {
	FindProfile(ctx *fiber.Ctx) error
}

type ProfileControllerImpl struct {
}

func NewProfileController() ProfileController {
	return &ProfileControllerImpl{}
}

func (controller *ProfileControllerImpl) FindProfile(ctx *fiber.Ctx) error {

	id := ctx.Params("profileId")

	data, profileResponse, err := model.FindProfile(id)
	if err != nil {
		return err
	}

	jsonErr := json.Unmarshal(data, &profileResponse)
	if jsonErr != nil {
		return fmt.Errorf("failed to unmarshal response : %v", jsonErr)
	}

	return ctx.Render("profile", profileResponse)

}
