package handlers

import (
	"net/http"
)

func (handler *AWSHandler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["http-server-alive"] = "Ok"
	awsWrapper := AWSWrapper{request: r, writer: w}
	awsWrapper.RespondWithJSON(http.StatusOK, response)
}
