package aws

import (
	"fmt"
	"net/http"
)

func (handler *AWSHandler) IAMUsersHandler(w http.ResponseWriter, r *http.Request) {
	profile := r.Header.Get("profile")
	cfg := loadAwsConfig(handler.multiple, profile, w)
	key := fmt.Sprintf("aws.%s.iam.users", profile)
	response, foundInCache := handler.cache.Get(key)
	if foundInCache {
		respondWithJSON(w, http.StatusOK, response)
	} else {
		response, err := handler.aws.DescribeIAMUsers(cfg)
		errMsg := "iam:ListUsers - Policy Denied"
		handler.respondWithJSONandSetCache(response, err, w, errMsg, key)
	}
}

func (handler *AWSHandler) IAMUserHandler(w http.ResponseWriter, r *http.Request) {
	profile := r.Header.Get("profile")
	cfg := loadAwsConfig(handler.multiple, profile, w)
	key := fmt.Sprintf("aws.%s.iam.user", profile)

	response, foundInCache := handler.cache.Get(key)
	if foundInCache {
		respondWithJSON(w, http.StatusOK, response)
	} else {
		response, err := handler.aws.DescribeIAMUser(cfg)
		errMsg := "iam:GetUser is missing - Policy Denied"
		handler.respondWithJSONandSetCache(response, err, w, errMsg, key)
	}
}
