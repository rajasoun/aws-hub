package main

import "testing"

func Test_main(t *testing.T) {
	type fields struct {
		Name string
		ID   int
	}
	tests := []struct {
		Name   string
		fields fields
		want   string
		want1  int
	}{
		// TODO: Add test cases.
		{Name: "Employee",
			fields: fields{
				Name: "Rajesh",
				ID:   002,
			},
			want:  "Rajesh",
			want1: 002,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			main()
		})
	}
}
