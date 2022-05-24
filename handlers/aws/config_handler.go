package aws

import (
	"net/http"
)

func (handler *AWSHandler) ConfigProfilesHandler(w http.ResponseWriter, r *http.Request) {
	if handler.multiple {
		sections := readCredentials(w)
		respondWithJSON(w, http.StatusOK, sections.List())
	} else {
		respondWithJSON(w, http.StatusOK, []string{})
	}
}
