package main

import (
	shortener "github.com/batt0s/goshort/shortener"
)

func main() {
	app := shortener.App{}
	// port := os.Getenv("PORT")
	app.Init()
	app.Run("8080")
}
