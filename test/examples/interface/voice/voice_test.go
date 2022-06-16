package voice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVoice(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		lang string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Check English",
			args: args{
				lang: "English",
			},
			want: "Hello",
		},
		{
			name: "Check Spanish",
			args: args{
				lang: "Spanish",
			},
			want: "Hola",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			voice := NewVoice(tt.args.lang)
			got := voice.SayHello()
			assert.Equal(tt.want, got, "voice.SayHello() = %v, want %v", got, tt.want)
		})
	}
}
