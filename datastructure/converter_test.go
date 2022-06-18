package ds

import (
	"errors"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/stretchr/testify/assert"
)

func TestStructToJSON(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	tests := []struct {
		name      string
		dsHandler DataStructure
		result    interface{}
		wantErr   bool
	}{
		{
			name:      "Check StructToJSON",
			dsHandler: New(),
			result:    costexplorer.GetCostAndUsageOutput{},
			wantErr:   false,
		},
		{
			name:      "Check StructToJSON for Err JSON",
			dsHandler: New(),
			result:    map[string]interface{}{"foo": make(chan int)},
			wantErr:   true,
		},
		{
			name: "Check StructToJSON for File Creation Error",
			dsHandler: DataStructure{
				fileCreator: func(name string) (*os.File, error) {
					return nil, errors.New("simulated error")
				},
				fileName: "",
			},
			result:  costexplorer.GetCostAndUsageOutput{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.dsHandler.StructToJSON(tt.result)
			if tt.wantErr {
				assert.Error(err, "DataStructure.StructToJSON() error = %v", err)
				return
			}
			assert.NoError(err, "DataStructure.StructToJSON() error = %v", err)
		})
	}
}
