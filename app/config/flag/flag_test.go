package flag

import (
	"testing"

	"github.com/rajasoun/aws-hub/app/config/arg"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func Test_app_flag_GetFlags(t *testing.T) {
	tests := []struct {
		name      string
		flagIndex int
		want      []cli.Flag
	}{
		{
			name: "Check Cache Flag",
			want: []cli.Flag{
				&cli.StringFlag{
					Name:  "cache, c",
					Usage: "Cache Type",
					Value: arg.DefaultCacheType,
				},
			},
			flagIndex: 2,
		},
	}
	for _, tt := range tests {
		assert := assert.New(t)
		t.Parallel()
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want[0].Names()
			got := GetFlags()[tt.flagIndex].Names()
			assert.Equal(want, got, "GetFlags() = %v, want %v", got, want)
		})
	}
}
