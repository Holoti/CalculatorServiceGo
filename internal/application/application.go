package application

import (
	"fmt"
	"net/http"

	"github.com/holoti/CalculatorServiceGo/internal/handlers"
	"github.com/holoti/CalculatorServiceGo/internal/middleware"
)

type Config struct {
	Port int
}

type Application struct {
	Cfg Config
}

func New(config Config) *Application {
	return &Application{
		Cfg: config,
	}
}

func (a *Application) Run() {
	http.Handle("/api/v1/calculate", middleware.ValidateMethod(handlers.CalcHandler))
	http.ListenAndServe(fmt.Sprintf("localhost:%d", a.Cfg.Port), nil)
}
