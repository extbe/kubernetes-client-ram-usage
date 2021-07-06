package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	openapi_v2 "github.com/googleapis/gnostic/openapiv2"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func main() {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(pprof.New())

	app.Post("/charts/install", func(ctx *fiber.Ctx) error {
		r, err := installChart()
		if err != nil {
			return err
		}

		return ctx.JSON(r)
	})

	panic(app.Listen(":8080"))
}

func installChart() (*openapi_v2.Document, error) {
	opts := genericclioptions.NewConfigFlags(true)
	client, err := opts.ToDiscoveryClient()
	if err != nil {
		return nil, err
	}

	return client.OpenAPISchema()
}
