package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/batt0s/goshort/controllers"
)

func main() {
	appMode := os.Getenv("APP_MODE")
	if appMode == "" {
		appMode = "dev"
		log.Println("[warning] No APP_MODE in env. Defaulting to dev.")
	}

	// Create and init app
	app := controllers.App{}
	app.Init(appMode)

	// Gracefully shutdown
	shutdown := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)

		// Interrupt signal sent from terminal
		signal.Notify(sigint, os.Interrupt)

		<-sigint

		// when recieved a interrupt signal
		log.Println("[info] Interrupt signal recieved. Gracefully stopping.")
		err := app.Server.Shutdown(context.Background())
		if err != nil {
			log.Println("[error] HTTP Server Shutdown Error :", err.Error())
		}
		log.Println("[info] Stopped.")
		close(shutdown)
	}()

	app.Run()

	<-shutdown
}
