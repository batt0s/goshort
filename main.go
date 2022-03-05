package main

import (
	shortener "github.com/batt0s/goshort/shortener"
)

func main() {
	app := shortener.App{}
	app.Init()
	app.Run("8080")
}
