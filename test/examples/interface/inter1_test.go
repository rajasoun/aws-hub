package spike

import "testing"

func TestStudent_Student(t *testing.T) {
	type fields struct {
		Name   string
		Course string
		id     int
	}
	tests := []struct {
		name   string
		fields fields
		want0  string
		want1  string
		want2  int
	}{
		// TODO: Add test cases.
		{name: "Pratim",
			fields: fields{
				Name:   "Pratim Sil",
				Course: "BE",
				id:     736,
			},
			want0: "Pratim Sil",
			want1: "BE",
			want2: 736,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Student{
				Name:   tt.fields.Name,
				Course: tt.fields.Course,
				id:     tt.fields.id,
			}
			got, got1, got2 := e.Student()
			if got != tt.want0 {
				t.Errorf("Student.Student() got = %v, want %v", got, tt.want0)
			}
			if got1 != tt.want1 {
				t.Errorf("Student.Student() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("Student.Student() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
