package controller

import (
	"encoding/json"
	"fmt"

	"github.com/faridlan/client-go/model"
	"github.com/gofiber/fiber/v2"
)

type ProfileController interface {
	FindProfile(ctx *fiber.Ctx) error
	CreateProfile(ctx *fiber.Ctx) error
	CreateProfileRender(ctx *fiber.Ctx) error
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

func (controller *ProfileControllerImpl) CreateProfile(ctx *fiber.Ctx) error {

	profile := &model.Profile{
		Name:        ctx.FormValue("name"),
		Description: ctx.FormValue("description"),
		Email:       ctx.FormValue("email"),
		MediaSocial: model.MediaSocial{
			LinkedIn:  ctx.FormValue("linked_in"),
			Instagram: ctx.FormValue("instagram"),
			GitHub:    ctx.FormValue("github"),
		},
		About: ctx.FormValue("about"),
	}

	statusCode, data, err := model.CreateProfile(profile)
	if err != nil {
		return err
	}

	fmt.Println("Response from the server:", string(data))

	if statusCode != fiber.StatusOK {
		return ctx.Status(statusCode).SendString("failed to send data")
	}

	return ctx.SendString("form submission recceived")

}

func (controller *ProfileControllerImpl) CreateProfileRender(ctx *fiber.Ctx) error {

	return ctx.Render("profile_create", nil)

}
