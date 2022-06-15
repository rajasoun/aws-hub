package arg

import (
	"testing"

	"github.com/rajasoun/aws-hub/service/cache"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestCliContext(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	tests := []struct {
		name                   string
		ctx                    *CliContext
		wantPort               int
		wantDuration           int
		wantIsMultipleProfiles bool
		wantCacheType          string
	}{
		{
			name:                   "Default Cli Context",
			ctx:                    NewDefaultCliContext(),
			wantPort:               3000,
			wantDuration:           30,
			wantIsMultipleProfiles: false,
			wantCacheType:          "InMemoryCache",
		},
		{
			name:                   "Custom Cli Context",
			ctx:                    NewCustomCliContext(9000, 30, &cache.Redis{}, true),
			wantPort:               9000,
			wantDuration:           30,
			wantIsMultipleProfiles: true,
			wantCacheType:          "RedisCache",
		},
		{
			name:                   "Default Cli Context",
			ctx:                    NewCliContext(&cli.Context{}),
			wantPort:               3000,
			wantDuration:           30,
			wantIsMultipleProfiles: false,
			wantCacheType:          "InMemoryCache",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("Check Port", func(t *testing.T) {
				gotPort := tt.ctx.port
				assert.Equal(tt.wantPort, gotPort,
					"got tt.ctx.Port() = %v, want = %v", gotPort, tt.wantPort)
			})

			t.Run("Check Duration", func(t *testing.T) {
				gotDuration := tt.ctx.duration
				assert.Equal(tt.wantDuration, gotDuration,
					"got tt.ctx.Duration() = %v, want = %v", gotDuration, tt.wantDuration)
			})

			t.Run("Check Profile", func(t *testing.T) {
				gotIsMultipleProfiles := tt.ctx.isMultipleAWSProfile
				assert.Equal(tt.wantIsMultipleProfiles, gotIsMultipleProfiles,
					"got tt.ctx.IsMultipleAwsProfiles() = %v, want = %v", gotIsMultipleProfiles, tt.wantIsMultipleProfiles)
			})

			t.Run("Check Cache", func(t *testing.T) {
				gotCacheType := tt.ctx.cache.Type()
				assert.Equal(tt.wantCacheType, gotCacheType,
					"got tt.ctx.Cache().Type() = %v, want = %v ", gotCacheType, tt.wantCacheType)

			})
			t.Run("Check Cache", func(t *testing.T) {
				tt.ctx.GetCacheHandler("redis")
				gotCacheType := tt.ctx.cache.Type()
				assert.Equal(tt.wantCacheType, gotCacheType,
					"got cliContext.Cache().Type() = %v, want = %v ", gotCacheType, tt.wantCacheType)

			})
		})
	}
}
