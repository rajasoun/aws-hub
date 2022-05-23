package aws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
)

func readCredentialForProfile(r *http.Request, handler *AWSHandler, w http.ResponseWriter) (string, aws.Config) {
	profile := r.Header.Get("profile")
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Default AWSConfig loaded successfuly")
	}
	if handler.multiple {
		cfg, err = external.LoadDefaultAWSConfig(external.WithSharedConfigProfile(profile))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't read "+profile+" profile")
		}
	}
	return profile, cfg
}

func handleUsersResponse(handler *AWSHandler, cfg aws.Config, w http.ResponseWriter, key string) {
	response, err := handler.aws.DescribeIAMUsers(cfg)
	msg := "iam:ListUsers is missing"
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, msg)
	} else {
		handler.cache.Set(key, response)
		respondWithJSON(w, http.StatusOK, response)
	}
}

func handleUserResponse(handler *AWSHandler, cfg aws.Config, w http.ResponseWriter, key string) {
	response, err := handler.aws.DescribeIAMUser(cfg)
	msg := "iam:GetUser is missing"
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, msg)
	} else {
		handler.cache.Set(key, response)
		respondWithJSON(w, http.StatusOK, response)
	}
}

func handleOrgResponse(handler *AWSHandler, cfg aws.Config, w http.ResponseWriter, key string) {
	response, err := handler.aws.DescribeOrganization(cfg)
	msg := "organizations:DescribeOrganization is missing"
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, msg)
	} else {
		handler.cache.Set(key, response)
		respondWithJSON(w, http.StatusOK, response)
	}
}

func (handler *AWSHandler) IAMUsersHandler(w http.ResponseWriter, r *http.Request) {
	profile, cfg := readCredentialForProfile(r, handler, w)
	key := fmt.Sprintf("aws.%s.iam.users", profile)
	response, foundInCache := handler.cache.Get(key)
	if foundInCache {
		respondWithJSON(w, http.StatusOK, response)
	} else {
		handleUsersResponse(handler, cfg, w, key)
	}
}

func (handler *AWSHandler) IAMUserHandler(w http.ResponseWriter, r *http.Request) {
	profile, cfg := readCredentialForProfile(r, handler, w)
	key := fmt.Sprintf("aws.%s.iam.user", profile)

	response, foundInCache := handler.cache.Get(key)
	if foundInCache {
		respondWithJSON(w, http.StatusOK, response)
	} else {
		handleUserResponse(handler, cfg, w, key)
	}
}

func (handler *AWSHandler) DescribeOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	profile, cfg := readCredentialForProfile(r, handler, w)
	key := fmt.Sprintf("aws.%s.iam.organization", profile)

	response, foundInCache := handler.cache.Get(key)
	if foundInCache {
		respondWithJSON(w, http.StatusOK, response)
	} else {
		handleOrgResponse(handler, cfg, w, key)
	}
}
