package controller

import (
	"encoding/json"
	"fmt"

	"github.com/faridlan/client-go/model"
	"github.com/gofiber/fiber/v2"
)

type SkillController interface {
	SKillRender(ctx *fiber.Ctx) error
}

type SkillControllerImpl struct {
}

func NewSkillController() SkillController {
	return &SkillControllerImpl{}
}

func (controller *SkillControllerImpl) SKillRender(ctx *fiber.Ctx) error {

	url := "http://localhost:9090/api/skills"

	agent := fiber.Get(url)

	_, data, err := agent.Bytes()
	if err != nil {
		return fmt.Errorf("failed to make request: %v", err)
	}

	skillResponse := model.WebResponse{
		Data: []model.Skill{},
	}

	jsonErr := json.Unmarshal(data, &skillResponse)
	if jsonErr != nil {
		return fmt.Errorf("failed to unmarshal response: %v", jsonErr)
	}

	return ctx.Render("index", skillResponse)

}
