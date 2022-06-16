package function

import (
	"testing"
)

func TestStudents_Students(t *testing.T) {
	type args struct {
		Name    string
		Class   string
		Section string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
		want2 string
	}{
		{
			name: "Student 1",
			args: args{
				Name:    "Rohini",
				Class:   "First",
				Section: "Star",
			},
			want:  "Rohini",
			want1: "First",
			want2: "Star",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Students{
				Name:    tt.args.Name,
				Class:   tt.args.Class,
				Section: tt.args.Section,
			}
			got, got1, got2 := a.Students()
			if got != tt.want {
				t.Errorf("Students.Students() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Students.Students() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("Students.Students() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
