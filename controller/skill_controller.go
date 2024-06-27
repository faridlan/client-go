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

	// // Manually encode form data
	// form := url.Values{}
	// form.Add("name", skill.Name)
	// formData := form.Encode()

	// // Create a new POST request using Fiber's HTTP client
	// agent := fiber.Post("http://localhost:9090/api/skills")
	// agent.Set(fiber.HeaderContentType, fiber.MIMEApplicationForm)
	// agent.Body([]byte(formData))

	// // Send the request and get the response
	// statusCode, data, err := agent.Bytes()
	// if err != nil {
	// 	return fmt.Errorf("failed to make request: %v", err)
	// }

	// Log the response
	fmt.Println("Response from the server:", string(data))

	// Check if response status is not OK
	if statusCode != fiber.StatusOK {
		return ctx.Status(statusCode).SendString("Failed to send data")
	}

	// Return success response to the client
	return ctx.SendString("Form submission received")

	// skill := &model.Skill{
	// 	Name: ctx.FormValue("name"),
	// }

	// err := model.CreateSkill(skill)
	// if err != nil {
	// 	return err
	// }

	// return ctx.JSON(fiber.Map{
	// 	"status":  "success",
	// 	"message": "Form submission received",
	// })

}

func (controller *SkillControllerImpl) DeleteSKill(ctx *fiber.Ctx) error {

	id := ctx.Params("skillId")

	err := model.DeleteSkill(id)
	if err != nil {
		return err
	}

	return nil

}

func (controller *SkillControllerImpl) UpdateSkill(ctx *fiber.Ctx) error {

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
