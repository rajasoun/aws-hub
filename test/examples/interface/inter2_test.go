package spike

import "testing"

func TestCompany_SalaryCalc(t *testing.T) {
	type fields struct {
		Name     string
		Location string
		basicpay int
		pf       int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  string
		want2  int
		want3  int
	}{
		// TODO: Add test cases.
		{name: "xyz",
			fields: fields{
				Name:     "RSS",
				Location: "Chennai",
				basicpay: 10000,
				pf:       1500,
			},
			want:  "RSS",
			want1: "Chennai",
			want2: 10000,
			want3: 1500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Company{
				Name:     tt.fields.Name,
				Location: tt.fields.Location,
				basicpay: tt.fields.basicpay,
				pf:       tt.fields.pf,
			}
			got, got1, got2, got3 := c.SalaryCalc()
			if got != tt.want {
				t.Errorf("Company.SalaryCalc() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Company.SalaryCalc() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("Company.SalaryCalc() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("Company.SalaryCalc() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}
