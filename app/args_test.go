package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func Test_ArgsAll(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	tests := []struct {
		name                   string
		cliContext             *cli.Context
		wantPort               int
		wantDuration           int
		wantIsMultipleProfiles bool
		wantCacheType          string
	}{
		{"Check Default Arguments", &cli.Context{}, 3000, 30, false, "InMemoryCache"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cliContext := NewCliContext(tt.cliContext)
			gotPort := cliContext.Port()
			assert.Equal(tt.wantPort, gotPort,
				"got cliContext.Port() = %v, want = %v", gotPort, tt.wantPort)
			gotDuration := cliContext.Duration()
			assert.Equal(tt.wantPort, gotDuration,
				"got cliContext.Duration() = %v, want = %v", gotDuration, tt.wantDuration)
			gotIsMultipleProfiles := cliContext.IsMultipleAwsProfiles()
			assert.Equal(tt.wantIsMultipleProfiles, gotIsMultipleProfiles,
				"got cliContext.IsMultipleAwsProfiles() = %v, want = %v", gotIsMultipleProfiles, tt.wantIsMultipleProfiles)
			gotCacheType := cliContext.Cache().Type()
			assert.Equal(tt.wantCacheType, gotCacheType,
				"got cliContext.Cache().Type() = %v, want = %v ", gotCacheType, tt.wantCacheType)
		})
	}
}
