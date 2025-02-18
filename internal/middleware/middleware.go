package middleware

import (
	"encoding/json"
	"net/http"
)

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
