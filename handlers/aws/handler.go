package aws

import (
	"encoding/json"
	"net/http"

	"github.com/rajasoun/aws-hub/services/aws"
	"github.com/rajasoun/aws-hub/services/cache"
)

type AWSHandler struct {
	cache    cache.Cache
	multiple bool
	aws      aws.AWS
}

func NewAWSHandler(cache cache.Cache, multiple bool) *AWSHandler {
	awsHandler := AWSHandler{
		cache:    cache,
		multiple: multiple,
		aws:      aws.AWS{},
	}
	return &awsHandler
}

func (handler *AWSHandler) GetAWSHandler() aws.AWS {
	return handler.aws
}

func (handler *AWSHandler) HasMultipleEnvs() bool {
	return handler.multiple
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
