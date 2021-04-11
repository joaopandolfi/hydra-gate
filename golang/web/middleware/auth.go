package middleware

import (
	"net/http"

	"hydra_gate/utils/security"

	"hydra_gate/utils/logger"

	"github.com/gorilla/mux"
)

// TokenHandler -
// @handler
// Intercept all transactions and check if is authenticated by token
func TokenHandler(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Headers", "*")

			Response(w, "", 200)
			return
		}
		url := r.URL.String()

		token := GetHeader(r, "token")
		userID := GetHeader(r, "id")

		t, err := security.CheckJwtToken(token)

		if !t.Authorized || err != nil || t.ID != userID {
			logger.Debug("[TokenHandler]", "Auth Error", url)
			Response(w, "No cookies for you", http.StatusForbidden)
			return
		}

		InjectHeader(r, "_xid", t.ID)

		logger.Debug("[TokenHandler]", "Authenticated", url)
		next.ServeHTTP(w, r)
	})
}

// AuthTokenedProtection - Chain Logged handler to protect connections
// @middleware
// Uses session stored value `logged` to make a best gin of the world
// If is not connected, check token
func AuthTokenedProtection(f http.HandlerFunc) http.HandlerFunc {
	return Chain(f, TokenHandler)
}

// HandleToken -
func HandleToken(r *mux.Router, path string, f http.HandlerFunc, methods ...string) {
	r.HandleFunc(path, AuthTokenedProtection(f)).Methods(methods...)
}
