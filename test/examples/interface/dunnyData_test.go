package spike

import "testing"

func TestUserAvilable(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Checking",
			args: args{
				email: "ajithkumar.sinha@srsconsultinginc.com",
			},
			want: UserAvilable("ajithkumar.sinha@srsconsultinginc.com"),
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UserAvilable(tt.args.email); got != tt.want {
				t.Errorf("UserAvilable() = %v, want %v", got, tt.want)
			}
		})
	}
}
