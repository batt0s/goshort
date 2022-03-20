package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	shortener "github.com/batt0s/goshort/shortener"
)

func main() {
	app := shortener.App{}

	shutdown := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)

		// Interrupt signal sent from terminal
		signal.Notify(sigint, os.Interrupt)

		<-sigint

		// when recieved a interrupt signal
		log.Println("Interrupt signal recieved. Gracefully stopping.")
		err := app.Server.Shutdown(context.Background())
		if err != nil {
			log.Println("HTTP Server Shutdown Error :", err.Error())
		}
		log.Println("Stopped.")
		close(shutdown)
	}()

	app.Init()
	app.Run()

	<-shutdown
}
