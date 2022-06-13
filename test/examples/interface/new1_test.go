package spike

import "testing"

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
