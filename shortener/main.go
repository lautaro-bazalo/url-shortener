package main

import "shortener/internal/application"

func main() {
	app := application.NewApp()

	app.Run()
}
