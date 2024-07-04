package model

import (
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

type Profile struct {
	ID          string      `json:"id" form:"id"`
	Name        string      `json:"name" form:"name"`
	Description string      `json:"description" form:"description"`
	Email       string      `json:"email" form:"email"`
	MediaSocial MediaSocial `json:"media_social" form:"media_social"`
	About       string      `json:"about" form:"about"`
}

type MediaSocial struct {
	LinkedIn  string `json:"linked_in" form:"linked_in"`
	Instagram string `json:"instagram" form:"instagram"`
	GitHub    string `json:"github" form:"github"`
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

	fmt.Println(webResponse)

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
