package controller

import (
	"encoding/json"
	"fmt"

	"github.com/faridlan/client-go/model"
	"github.com/gofiber/fiber/v2"
)

type SkillController interface {
	SKillRender(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	CreateRender(ctx *fiber.Ctx) error
	CreateSKill(ctx *fiber.Ctx) error
	DeleteSKill(ctx *fiber.Ctx) error
	UpdateSkillRender(ctx *fiber.Ctx) error
	UpdateSkill(ctx *fiber.Ctx) error
}

type SkillControllerImpl struct {
}

func NewSkillController() SkillController {
	return &SkillControllerImpl{}
}

func (controller *SkillControllerImpl) SKillRender(ctx *fiber.Ctx) error {

	skillResponse, err := model.FindAllSkill()
	if err != nil {
		return err
	}

	return ctx.Render("index", skillResponse)

}

func (controller *SkillControllerImpl) FindById(ctx *fiber.Ctx) error {

	id := ctx.Params("skillId")

	data, skillResponse, err := model.FindAllSkilByIdl(id)
	if err != nil {
		return err
	}

	jsonErr := json.Unmarshal(data, &skillResponse)
	if jsonErr != nil {
		return fmt.Errorf("failed to unmarshal response: %v", jsonErr)
	}

	return ctx.Render("detail", skillResponse)

}

func (controller *SkillControllerImpl) CreateRender(ctx *fiber.Ctx) error {

	return ctx.Render("skill_create", nil)

}

func (controller *SkillControllerImpl) CreateSKill(ctx *fiber.Ctx) error {

	skill := &model.Skill{
		Name: ctx.FormValue("name"),
	}

	statusCode, data, err := model.CreateSkill(skill)
	if err != nil {
		return err
	}

	// Log the response
	fmt.Println("Response from the server:", string(data))

	// Check if response status is not OK
	if statusCode != fiber.StatusOK {
		return ctx.Status(statusCode).SendString("Failed to send data")
	}

	// Return success response to the client
	return ctx.SendString("Form submission received")

}

func (controller *SkillControllerImpl) DeleteSKill(ctx *fiber.Ctx) error {

	id := ctx.Params("skillId")

	err := model.DeleteSkill(id)
	if err != nil {
		return err
	}

	return nil

}

func (controller *SkillControllerImpl) UpdateSkillRender(ctx *fiber.Ctx) error {

	id := ctx.Params("skillId")

	data, skillResponse, err := model.FindAllSkilByIdl(id)
	if err != nil {
		return err
	}

	jsonErr := json.Unmarshal(data, &skillResponse)
	if jsonErr != nil {
		return fmt.Errorf("failed to unmarshal response: %v", jsonErr)
	}

	return ctx.Render("update", skillResponse)

}

func (controller *SkillControllerImpl) UpdateSkill(ctx *fiber.Ctx) error {

	// id := ctx.Params("skillId")

	fmt.Println(ctx.FormValue("id"))
	fmt.Println(ctx.FormValue("name"))

	// id := ctx.Params("skillId")

	skill := &model.Skill{
		// ID:   ctx.FormValue("id"),
		Name: ctx.FormValue("name"),
	}

	statusCode, data, err := model.UpdateSkill(skill)
	if err != nil {
		return err
	}

	fmt.Println("Response from the server:", string(data))

	// Check if response status is not OK
	if statusCode != fiber.StatusOK {
		return ctx.Status(statusCode).SendString("Failed to send data")
	}

	// Return success response to the client
	return ctx.SendString("Form submission received")

}
