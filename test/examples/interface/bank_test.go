package spike

import "testing"

func Test_acc_details_account(t *testing.T) {
	type fields struct {
		acc_holder_name string
		acc_number      int
		acc_type        string
		acc_balance     float32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  int
		want2  string
		want3  float32
	}{
		// TODO: Add test cases.
		{name : "manu",
		 fields: fields{
			acc_holder_name: "Manu V. H",
			acc_number  :    012345,
			acc_type   :     "Savings Account",
			acc_balance :    5324.75,
		 },
		 want : "Manu V. H",
		 want1: 012345,
		 want2 : "Savings Account",
		 want3 : 5324.75,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ad := &acc_details{
				acc_holder_name: tt.fields.acc_holder_name,
				acc_number:      tt.fields.acc_number,
				acc_type:        tt.fields.acc_type,
				acc_balance:     tt.fields.acc_balance,
			}
			got, got1, got2, got3 := ad.account()
			if got != tt.want {
				t.Errorf("acc_details.account() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("acc_details.account() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("acc_details.account() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("acc_details.account() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}
