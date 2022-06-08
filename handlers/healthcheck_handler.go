package handlers

import (
	"net/http"
)

func (handler *AWSHandler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["http-server-alive"] = "Ok"
	respondWithJSON(w, http.StatusOK, response)
}
