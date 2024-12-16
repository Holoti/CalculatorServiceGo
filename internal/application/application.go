package application

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/holoti/CalculatorServiceGo/pkg/calculate"
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

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	expression := r.URL.Query().Get("expression")
	result, err := calculate.Calculate(expression)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Expression is not valid",
		})
		return
	}
	rounded := math.Round(result*math.Pow10(6)) / math.Pow10(6)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"result": fmt.Sprint(rounded),
	})
}

func Logging(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expression := r.URL.Query().Get("expression")
		log.Printf("expression: %s", expression)
		next.ServeHTTP(w, r)
	})
}

func (a *Application) Run() {
	http.Handle("/api/v1/calculate", Logging(CalculateHandler))
	http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", a.Cfg.Port), nil)
}
