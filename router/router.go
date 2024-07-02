package router

import (
	"github.com/faridlan/client-go/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) {

	skillController := controller.NewSkillController()
	profileController := controller.NewProfileController()

	app.Get("/", skillController.SKillRender)
	app.Get("/:skillId", skillController.FindById)
	app.Get("/form/skills", skillController.CreateRender)
	app.Post("/form/skills/x", skillController.CreateSKill)
	app.Delete("/skills/:skillId", skillController.DeleteSKill)
	app.Get("/skills/:skillId", skillController.UpdateSkillRender)
	app.Put("/skills/update", skillController.UpdateSkill)

	app.Get("/profiles/:profileId", profileController.FindProfile)

}
