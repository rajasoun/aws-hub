package api

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func TestGetAliasesAPI_Execute(t *testing.T) {
	type args struct {
		client *iam.Client
	}
	tests := []struct {
		name    string
		api     GetAliasesAPI
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := GetAliasesAPI{}
			got, err := api.Execute(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAliasesAPI.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAliasesAPI.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
