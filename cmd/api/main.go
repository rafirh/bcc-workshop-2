package main

import (
	"bcc-workshop-2/config"
	"bcc-workshop-2/internal/app"
)

func main() {
	cfg := config.New()

	application := app.New(cfg)
	application.Run()
}
