package spike

//ajit

import "testing"

func TestPerimeter(t *testing.T) {
	type args struct {
		wi float64
		he float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "perimeter",
			args: args{
				wi: 1,
				he: 1,
			},
			want: 4.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Perimeter(tt.args.wi, tt.args.he); got != tt.want {
				t.Errorf("Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
