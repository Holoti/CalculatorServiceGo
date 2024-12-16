package main

import "github.com/holoti/CalculatorServiceGo/internal/application"

func main() {
	config := application.Config{
		Port: 8080,
	}
	app := application.New(config)
	app.Run()
}
