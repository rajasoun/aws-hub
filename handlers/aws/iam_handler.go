package aws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws/external"
)

func (handler *AWSHandler) IAMUsersHandler(w http.ResponseWriter, r *http.Request) {
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

	key := fmt.Sprintf("aws.%s.iam.users", profile)

	response, found := handler.cache.Get(key)
	if found {
		respondWithJSON(w, http.StatusOK, response)
	} else {
		response, err := handler.aws.DescribeIAMUsers(cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "iam:ListUsers is missing")
		} else {
			handler.cache.Set(key, response)
			respondWithJSON(w, http.StatusOK, response)
		}
	}
}

func (handler *AWSHandler) IAMUserHandler(w http.ResponseWriter, r *http.Request) {
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

	key := fmt.Sprintf("aws.%s.iam.user", profile)

	response, found := handler.cache.Get(key)
	if found {
		respondWithJSON(w, http.StatusOK, response)
	} else {
		response, err := handler.aws.DescribeIAMUser(cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "iam:GetUser is missing")
		} else {
			handler.cache.Set(key, response)
			respondWithJSON(w, http.StatusOK, response)
		}
	}
}

func (handler *AWSHandler) DescribeOrganizationHandler(w http.ResponseWriter, r *http.Request) {
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

	key := fmt.Sprintf("aws.%s.iam.organization", profile)

	response, found := handler.cache.Get(key)
	if found {
		respondWithJSON(w, http.StatusOK, response)
	} else {
		response, err := handler.aws.DescribeOrganization(cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "organizations:DescribeOrganization is missing")
		} else {
			handler.cache.Set(key, response)
			respondWithJSON(w, http.StatusOK, response)
		}
	}
}
