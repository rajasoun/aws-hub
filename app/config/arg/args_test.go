package arg

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func Test_hub_config_arg_NewCliContext(t *testing.T) {
	assert := assert.New(t)
	cliContext := NewCliContext(&cli.Context{})
	t.Parallel()

	tests := []struct {
		name                   string
		wantPort               int
		wantDuration           int
		wantIsMultipleProfiles bool
		wantCacheType          string
	}{
		{"Default", 3000, 30, false, "InMemoryCache"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("Check Port", func(t *testing.T) {
				gotPort := cliContext.GetPort()
				assert.Equal(tt.wantPort, gotPort,
					"got cliContext.Port() = %v, want = %v", gotPort, tt.wantPort)
			})

			t.Run("Check Duration", func(t *testing.T) {
				gotDuration := cliContext.GetDuration()
				assert.Equal(tt.wantDuration, gotDuration,
					"got cliContext.Duration() = %v, want = %v", gotDuration, tt.wantDuration)
			})

			t.Run("Check Profile", func(t *testing.T) {
				gotIsMultipleProfiles := cliContext.GetAwsProfileType()
				assert.Equal(tt.wantIsMultipleProfiles, gotIsMultipleProfiles,
					"got cliContext.IsMultipleAwsProfiles() = %v, want = %v", gotIsMultipleProfiles, tt.wantIsMultipleProfiles)
			})

			t.Run("Check InMemory", func(t *testing.T) {
				gotCacheType := cliContext.GetCache().Type()
				assert.Equal(tt.wantCacheType, gotCacheType,
					"got cliContext.Cache().Type() = %v, want = %v ", gotCacheType, tt.wantCacheType)

			})
			t.Run("Check Cache", func(t *testing.T) {
				cliContext.cache = cliContext.GetCacheHandler("redis")
				gotCacheType := cliContext.GetCache().Type()
				assert.Equal(tt.wantCacheType, gotCacheType,
					"got cliContext.Cache().Type() = %v, want = %v ", gotCacheType, tt.wantCacheType)

			})
		})
	}
}