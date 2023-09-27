package route

import (
	handlers "github.com/Rest_Savings/apps/hadlers"
	"github.com/Rest_Savings/apps/services"
	"github.com/Rest_Savings/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, services *services.Services) {

	route := app.Group("/api/")
	handlers := &handlers.Handler{
		Ser: services,
	}

	route.Post("/daftar", handlers.CreateAccount)
	route.Get("/saldo/:no_rekening", handlers.GetBalaceAccount)
	route.Get("/mutasi/:no_rekening", handlers.GetListTransaction)
	auten := &utils.AuthorizeUser{
		Ser: services,
	}
	routeAuth := app.Group("/api/", auten.CheckRek)
	routeAuth.Post("/tabung", handlers.CreateTractionAdds)
	routeAuth.Post("/tarik", handlers.CreateTractionWidraw)

}
