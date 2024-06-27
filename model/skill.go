package model

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

type Skill struct {
	ID   string `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

func FindAllSkill() (*WebResponse, error) {

	url := "http://localhost:9090/api/skills"

	agent := fiber.Get(url)

	_, data, err := agent.Bytes()
	if err != nil {
		return nil, fmt.Errorf("failed to make request : %v", err)
	}

	webResponse := &WebResponse{
		Data: []Skill{},
	}

	jsonErr := json.Unmarshal(data, &webResponse)
	if jsonErr != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", jsonErr)
	}

	return webResponse, nil

}

func FindAllSkilByIdl(id string) ([]byte, *WebResponse, error) {

	url := fmt.Sprintf("http://localhost:9090/api/skills/%s", id)

	agent := fiber.Get(url)

	_, data, err := agent.Bytes()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to make request : %v", err)
	}

	webResponse := &WebResponse{
		Data: Skill{},
	}

	return data, webResponse, nil

}

func CreateSkill(skill *Skill) (int, []byte, error) {

	form := url.Values{}
	form.Add("name", skill.Name)
	formData := form.Encode()

	// Create a new POST request using Fiber's HTTP client
	agent := fiber.Post("http://localhost:9090/api/skills")
	agent.Set(fiber.HeaderContentType, fiber.MIMEApplicationForm)
	agent.Body([]byte(formData))

	// Send the request and get the response
	statusCode, data, err := agent.Bytes()
	if err != nil {
		return 0, nil, fmt.Errorf("failed to make request: %v", err)
	}

	return statusCode, data, nil

	// // Define the URL
	// skillUrl := "http://localhost:9090/api/skills"

	// form := url.Values{}
	// form.Add("name", skill.Name)
	// formData := form.Encode()

	// agent := fiber.Post(skillUrl)
	// agent.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	// agent.Body([]byte(formData))

	// return nil
}

func DeleteSkill(id string) error {

	url := fmt.Sprintf("http://localhost:9090/api/skills/%s", id)

	agent := fiber.Delete(url)

	_, _, err := agent.Bytes()
	if err != nil {
		return fmt.Errorf("failed to make request : %v", err)
	}

	return nil

}

func UpdateSkill(skill *Skill) (int, []byte, error) {

	form := url.Values{}
	form.Add("name", skill.Name)
	formData := form.Encode()

	// Create a new POST request using Fiber's HTTP client
	agent := fiber.Put(fmt.Sprintf("http://localhost:9090/api/skills/%s", skill.ID))
	agent.Set(fiber.HeaderContentType, fiber.MIMEApplicationForm)
	agent.Body([]byte(formData))

	// Send the request and get the response
	statusCode, data, err := agent.Bytes()
	if err != nil {
		return 0, nil, fmt.Errorf("failed to make request: %v", err)
	}

	return statusCode, data, nil

}
