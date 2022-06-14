package function

import "testing"

func TestEmployee(t *testing.T) {
	type args struct {
		name string
		age  int
		id   int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
		want2 int
	}{
		// TODO: Add test cases.
		{
			name: "sakmoto",
			args: args{
				name: "manu",
				age:  20,
				id:   007,
			},
			want:  "manu",
			want1: 20,
			want2: 007,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := Employee(tt.args.name, tt.args.age, tt.args.id)
			if got != tt.want {
				t.Errorf("Employee() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Employee() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("Employee() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
