package health

import (
	"net/http"

	"hydra_gate/config"
	"hydra_gate/web/controllers"
	"hydra_gate/web/middleware"
	"hydra_gate/web/server"

	"hydra_gate/utils/logger"
)

// --- Health ---

type controller struct {
	s *server.Server
}

// New Health controller
func New() controllers.Controller {
	return &controller{
		s: nil,
	}
}

// Health route
func (c *controller) health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	middleware.Response(w, true, http.StatusOK)
}

// Config database
func (c *controller) config(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	hash := middleware.GetHeader(r, "hash")
	logger.Debug("HEADER", hash, config.Get().Server.Security.Resethash)
	if hash != config.Get().Server.Security.Resethash {
		middleware.Response(w, "Not cookies for you", http.StatusForbidden)
		return
	}
	token := middleware.GetHeader(r, "token")

	var result []string
	errs := prepare(token)
	for _, e := range errs {
		if e != nil {
			result = append(result, e.Error())
		}
	}
	middleware.Response(w, result, http.StatusOK)
}

// ResetDatabase route
func (c *controller) resetDatabase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	hash := middleware.GetHeader(r, "hash")
	if hash != config.Get().Server.Security.Resethash {
		middleware.Response(w, "Invalid Hash", http.StatusForbidden)
		return
	}

	var errors []error
	//_, err := mongo.GetSession().GetCollection("files").RemoveAll(nil)
	//if err != nil {
	//	errors = append(errors, err)
	//}

	//Config
	//config()

	if len(errors) > 0 {
		middleware.Response(w, errors, http.StatusInternalServerError)
	} else {
		logger.Info("[ResetService] - Crypt-files database RESETED")
		middleware.Response(w, "Reseted", http.StatusOK)
	}
}

func prepare(token string) []error {
	return nil
}
