package main

import (
	"context"
	"os"
	"os/signal"

	shortener "github.com/batt0s/goshort/shortener"
)

func main() {
	app := shortener.App{}
	app.Init()

	// Gracefully shutdown
	shutdown := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)

		// Interrupt signal sent from terminal
		signal.Notify(sigint, os.Interrupt)

		<-sigint

		// when recieved a interrupt signal
		shortener.CustomLogger.Info("Interrupt signal recieved. Gracefully stopping.")
		err := app.Server.Shutdown(context.Background())
		if err != nil {
			shortener.CustomLogger.Error("HTTP Server Shutdown Error :", err.Error())
		}
		shortener.CustomLogger.Info("Stopped.")
		close(shutdown)
	}()

	app.Run()

	<-shutdown
}
