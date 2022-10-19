package handlers

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

type AppHandler struct{}

func NewApp() *AppHandler {
	return &AppHandler{}
}

func (handler AppHandler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	log.Debug().Msg("I'm health!")
	w.WriteHeader(http.StatusOK)
}
