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

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req struct {
		Expression string `json:"expression"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("error reading request: %v)", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Internal server error",
		})
		return
	}

	if req.Expression == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Bad request body",
		})
		return
	}

	result, err := calculate.Calc(req.Expression)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]string{
			"error": fmt.Sprintf("Expression is not valid: %v", err),
		})
		return
	}

	rounded := math.Round(result*math.Pow10(6)) / math.Pow10(6)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"result": fmt.Sprint(rounded),
	})
}

func ValidateMethod(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Method not allowed; use POST",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (a *Application) Run() {
	// http.HandleFunc("/api/v1/calculate", CalcHandler)
	http.Handle("/api/v1/calculate", ValidateMethod(CalcHandler))
	http.ListenAndServe(fmt.Sprintf("localhost:%d", a.Cfg.Port), nil)
}
