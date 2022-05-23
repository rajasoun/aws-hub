package aws

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	ini "github.com/rajasoun/go-config-parsers/aws_credentials"
)

func readCredentials(w http.ResponseWriter) ini.Sections {
	sections, err := ini.OpenFile(external.DefaultSharedCredentialsFilename())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't parse credentials file")
	}
	return sections
}

func (handler *AWSHandler) ConfigProfilesHandler(w http.ResponseWriter, r *http.Request) {
	if handler.multiple {
		sections := readCredentials(w)
		respondWithJSON(w, http.StatusOK, sections.List())
	} else {
		respondWithJSON(w, http.StatusOK, []string{})
	}
}
