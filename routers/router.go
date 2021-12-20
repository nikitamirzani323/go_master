package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nikitamirzani323/go_mastertoto/controller"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Static("/", "frontend/public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})

	app.Post("/api/init", controller.Init)
	app.Post("/api/login", controller.CheckLogin)
	app.Post("/api/company", controller.Company)
	app.Post("/api/companydetail", controller.Companydetail)
	app.Post("/api/companylistadmin", controller.Companylistadmin)
	app.Post("/api/companylistpasaran", controller.Companylistpasaran)
	app.Post("/api/companypasaranconf", controller.Companypasaranconf)
	app.Post("/api/companypasaranonline", controller.Companypasaranonline)
	app.Post("/api/savecompany", controller.Companysave)
	app.Post("/api/savecompanyadmin", controller.Companysaveadmin)
	app.Post("/api/savecompanyupdatepasaran", controller.Companysaveupdatepasaran)
	app.Post("/api/savecompanyupdatepasaranline", controller.Companysaveupdatepasaranline)
	app.Post("/api/savecompanypasaranonline", controller.Companysavepasaranonline)
	app.Post("/api/deletecompanypasaranonline", controller.Companydeletepasaranonline)
	app.Post("/api/updatecompanypasaran432", controller.Companyupdatepasaran432)
	app.Post("/api/updatecompanypasarancolokbebas", controller.Companyupdatepasarancolokbebas)
	app.Post("/api/updatecompanypasarancolokmacau", controller.Companyupdatepasarancolokmacau)
	app.Post("/api/updatecompanypasarancoloknaga", controller.Companyupdatepasarancoloknaga)
	app.Post("/api/updatecompanypasarancolokjitu", controller.Companyupdatepasarancolokjitu)
	app.Post("/api/updatecompanypasaran5050umum", controller.Companyupdatepasaran5050umum)
	app.Post("/api/updatecompanypasaran5050special", controller.Companyupdatepasaran5050special)
	app.Post("/api/updatecompanypasaran5050kombinasi", controller.Companyupdatepasaran5050kombinasi)
	app.Post("/api/updatecompanypasarankombinasi", controller.Companyupdatepasarankombinasi)
	app.Post("/api/updatecompanypasarandasar", controller.Companyupdatepasarandasar)
	app.Post("/api/updatecompanypasaranshio", controller.Companyupdatepasaranshio)

	app.Post("/api/allpasaran", controller.Pasaran)
	return app
}
