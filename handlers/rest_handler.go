package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rajasoun/aws-hub/handlers/api"
	"github.com/rajasoun/aws-hub/service/cache"
)

type RestAPI struct {
	request  *http.Request
	writer   http.ResponseWriter
	cache    cache.Cache
	multiple bool
}

func (restHandler *RestAPI) RespondWithJSON(code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	responseWritter := restHandler.writer
	responseWritter.Header().Set("Content-Type", "application/json")
	responseWritter.WriteHeader(code)
	responseWritter.Write(response)
}

func (restHandler *RestAPI) RespondWithErrorJSON(err error, errMsg string) {
	var errReasons string = "Possible Reasons: Connectivity Failed or Credential Missing or Policy Denied"
	errMessage := errMsg + " : " + errReasons
	if err != nil {
		code := http.StatusInternalServerError
		restHandler.RespondWithJSON(code, map[string]string{"error": errMessage})
	}
}

func (restHandler *RestAPI) InvokeAPI(apiName string, keyCode string, errMsg string) {
	profile := restHandler.request.Header.Get("profile")
	key := fmt.Sprintf(keyCode, profile)
	response, foundInCache := restHandler.cache.Get(key)
	if foundInCache {
		restHandler.RespondWithJSON(http.StatusOK, response)
		return
	} else {
		cfg, _ := api.GetConfig(profile, restHandler.multiple)
		api := api.NewAwsAPI(apiName)
		response, err := api.Execute(cfg)
		if err != nil {
			restHandler.RespondWithErrorJSON(err, errMsg)
		} else {
			restHandler.cache.Set(key, response)
			restHandler.RespondWithJSON(http.StatusOK, response)
		}
	}
}
