package controllers

import (
	"net/http"

	"github.com/jrsmth97/go-rest/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "HELLO WORLD")
}
