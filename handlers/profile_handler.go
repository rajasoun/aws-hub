package handlers

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/rajasoun/aws-hub/provider/credential"
)

type Profile struct {
	Multiple bool     `json:"multiple"`
	List     []string `json:"list"`
}

func (handler *AWSHandler) GetSections(w http.ResponseWriter, credentialFile string) []string {
	cl := credential.CredentialLoader{}
	sections, err := cl.GetSections(credentialFile)
	if err != nil {
		awsWrapper := AWSWrapper{writer: w}
		awsWrapper.RespondWithErrorJSON(err, "Couldn't parse credentials file")
	}
	return sections.List()
}

func (handler *AWSHandler) ListProfilesHandler(w http.ResponseWriter, r *http.Request) {
	var profile Profile
	var sectionList []string
	credentialFile := config.DefaultSharedCredentialsFilename()
	if handler.multiple {
		sectionList = handler.GetSections(w, credentialFile)
	} else {
		sectionList = []string{"default"}
	}
	profile = Profile{
		Multiple: handler.multiple,
		List:     sectionList,
	}
	awsWrapper := AWSWrapper{request: r, writer: w}
	awsWrapper.RespondWithJSON(http.StatusOK, profile)
}
