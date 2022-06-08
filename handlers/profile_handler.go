package handlers

import (
	"net/http"

	"github.com/rajasoun/aws-hub/provider/credential"
)

type Profile struct {
	Multiple bool     `json:"multiple"`
	List     []string `json:"list"`
}

func (handler *AWSHandler) GetSections(w http.ResponseWriter) []string {
	cl := credential.CredentialLoader{}
	sections, err := cl.GetSections()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't parse credentials file")
	}
	return sections.List()
}

func (handler *AWSHandler) ConfigProfilesHandler(w http.ResponseWriter, r *http.Request) {
	var profile Profile
	var sectionList []string
	if handler.multiple {
		sectionList = handler.GetSections(w)
	} else {
		sectionList = []string{"default"}
	}
	profile = Profile{
		Multiple: handler.multiple,
		List:     sectionList,
	}
	respondWithJSON(w, http.StatusOK, profile)
}
