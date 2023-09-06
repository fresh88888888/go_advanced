package main

import (
	"fmt"

	"umbrella.github.com/advanced_go/advanced_7/web-framewok/app"
)

func main() {
	apps := app.New()
	apps.Use(func(ctx *app.Context) {
		ctx.AddHeader("X-Info", "Hello")
	})
	apps.Get("/", func(ctx *app.Context) {
		ctx.Send("Hello World")
	})
	apps.Post("/add/user", func(ctx *app.Context) {
		name, _ := ctx.Query("name")
		if name == "" {
			ctx.Send("What's you name again?")
		} else {
			ctx.Send(fmt.Sprintf("Got Username: %s", name))
		}
	})

	apps.Run()
}
