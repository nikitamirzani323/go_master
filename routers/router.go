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
	app.Post("/api/savecompany", controller.Companysave)

	app.Post("/api/allpasaran", controller.Pasaran)
	return app
}
