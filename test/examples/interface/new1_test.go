package spike

import (
	"reflect"
	"testing"
)

func Test_perimeter_Perimeter(t *testing.T) {
	type fields struct {
		width  float64
		height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
		{
			name: "perimeter",
			fields: fields{
				width:  12,
				height: 12,
			},
			want: 48,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &perimeter{
				width:  tt.fields.width,
				height: tt.fields.height,
			}
			if got := p.Perimeter(); got != tt.want {
				t.Errorf("perimeter.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newaudio(t *testing.T) {
	type args struct {
		lang string
	}
	tests := []struct {
		name string
		args args
		want Audio
	}{
		// TODO: Add test cases.
		{
			name: "Check for English",
			args: args{
				lang: "Hello",
			},
			want: English{},
		},
		{
			name: "check for spanish",
			args: args{
				lang: "Hola",
			},
			want: Spanish{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newaudio(tt.args.lang); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newaudio() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayerDetails(t *testing.T) {
	tests := []struct {
		name string
		want *player
	}{
		// TODO: Add test cases.
		{
			name: "define the player name ",
			want: &player{
				name: "ajit kumar",
				age:  25,
				perimeter: perimeter{
					width:  10,
					height: 20,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PlayerDetails(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlayerDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}
