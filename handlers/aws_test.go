package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDefaultAWSHandler(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	tests := []struct {
		name          string
		multiple      bool
		wantCacheType string
	}{
		{
			name:          "Check New Default AWS Handler with no profile",
			multiple:      false,
			wantCacheType: "InMemoryCache",
		},
		{
			name:          "Check New Default AWS Handler with multiple profile",
			multiple:      true,
			wantCacheType: "InMemoryCache",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDefaultAWSHandler(tt.multiple)
			assert.Equal(got.cache.Type(), tt.wantCacheType,
				"NewDefaultAWSHandler() = %v, want %v", got, tt.wantCacheType)
		})
	}
}
