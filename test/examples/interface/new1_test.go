package spike

import "testing"

func Test_perimeter_Perimeter(t *testing.T) {
	type fields struct {
		wi float64
		he float64
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
				wi: 12,
				he: 12,
			},
			want: 48,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &perimeter{
				wi: tt.fields.wi,
				he: tt.fields.he,
			}
			if got := p.Perimeter(); got != tt.want {
				t.Errorf("perimeter.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
