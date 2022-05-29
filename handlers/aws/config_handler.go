package aws

import (
	"net/http"
)

type Profile struct {
	Multiple bool     `json:"multiple"`
	List     []string `json:"list"`
}

func (handler *AWSHandler) ConfigProfilesHandler(w http.ResponseWriter, r *http.Request) {
	var profile Profile
	if handler.multiple {
		sections := readLocalCredentials(w)
		profile = Profile{
			Multiple: handler.multiple,
			List:     sections.List(),
		}
		respondWithJSON(w, http.StatusOK, profile)
	} else {
		profile = Profile{
			Multiple: handler.multiple,
			List:     []string{"default"},
		}
		respondWithJSON(w, http.StatusOK, profile)
	}
}
