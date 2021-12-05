package controller

import (
	"log"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/go_mastertoto/config"
)

const PATH string = config.PATH_API

type responseinit struct {
	Status int    `json:"status"`
	Token  string `json:"token"`
}

func Init(c *fiber.Ctx) error {
	hostname := c.Hostname()
	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responseinit{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
		}).
		Post(PATH + "api/init")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responseinit)
	return c.JSON(fiber.Map{
		"status": result.Status,
		"token":  result.Token,
		"time":   time.Since(render_page).String(),
	})
}
