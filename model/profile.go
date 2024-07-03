package model

import (
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

type Profile struct {
	ID          string      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Email       string      `json:"email,omitempty"`
	MediaSocial MediaSocial `json:"media_social,omitempty"`
	About       string      `json:"about,omitempty"`
}

type MediaSocial struct {
	LinkedIn  string `json:"linked_in,omitempty"`
	Instagram string `json:"instagram,omitempty"`
	GitHub    string `json:"git_hub,omitempty"`
}

func FindProfile(id string) ([]byte, *WebResponse, error) {

	url := fmt.Sprintf("http://localhost:9090/api/profiles/%s", id)

	agent := fiber.Get(url)

	_, data, err := agent.Bytes()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to make reqeust : %v", err)
	}

	webResponse := &WebResponse{
		Data: Profile{},
	}

	return data, webResponse, nil

}

func CreateProfile(profile *Profile) (int, []byte, error) {

	form := url.Values{}
	form.Add("name", profile.Name)
	form.Add("description", profile.Description)
	form.Add("email", profile.Email)
	form.Add("linked_in", profile.MediaSocial.LinkedIn)
	form.Add("instagram", profile.MediaSocial.Instagram)
	form.Add("github", profile.MediaSocial.GitHub)
	form.Add("about", profile.About)
	formData := form.Encode()

	agent := fiber.Post("http://localhost:9090/api/profiles")
	agent.Set(fiber.HeaderContentType, fiber.MIMEApplicationForm)
	agent.Body([]byte(formData))

	statusCode, data, err := agent.Bytes()
	if err != nil {
		return 0, nil, fmt.Errorf("failed to make request : %v", err)
	}

	return statusCode, data, nil

}
