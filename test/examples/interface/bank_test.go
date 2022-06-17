package spike

import (
	"os"
	"reflect"
	"testing"
)

func Test_acc_details_getaccountnamenumber(t *testing.T) {
	type fields struct {
		acc_holder_name string
		acc_number      int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  int
	}{
		// TODO: Add test cases.
		{name: "Account Name & Number",
			fields: fields{
				acc_holder_name: "Manu V H",
				acc_number:      543210,
			},
			want:  "Manu V H",
			want1: 543210,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ad := &acc_details{
				acc_holder_name: tt.fields.acc_holder_name,
				acc_number:      tt.fields.acc_number,
			}
			got, got1 := ad.getaccountnamenumber()
			if got != tt.want {
				t.Errorf("Account Name: got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Account Number got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_acc_details_getaccounttypebalance(t *testing.T) {
	type fields struct {
		acc_type    string
		acc_balance float32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  float32
	}{
		// TODO: Add test cases.
		{name: "Account Type & Balance",
			fields: fields{
				acc_type:    "Savings",
				acc_balance: 5678.99,
			},
			want:  "Savings",
			want1: 5678.99,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ad := &acc_details{
				acc_type:    tt.fields.acc_type,
				acc_balance: tt.fields.acc_balance,
			}
			got, got1 := ad.getaccounttypebalance()
			if got != tt.want {
				t.Errorf("Account Type: got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Account Balance() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAccoutDetails(t *testing.T) {
	tests := []struct {
		name string
		want *acc_details
	}{
		// TODO: Add test cases.
		{name: "Account Details",
			want: &acc_details{
				acc_holder_name: "Manu V H",
				acc_number:      543210,
				acc_type:        "Savings",
				acc_balance:     5678.95,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AccoutDetails(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccoutDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatement(t *testing.T) {

	t.Run("Generate Statement.txt", func(t *testing.T) {
		filename := Statement().Name()
		checkFile, err := os.Stat(filename)
		actual := filename
		expected := checkFile.Name()
		if actual != expected {
			t.Errorf("actual: %v,  expected: %v, error: %v", actual, expected, err)
		}
	})
}

func TestUsers(t *testing.T) {
	tests := []struct {
		name string
		want users
	}{
		// TODO: Add test cases.
		{name: "Check value and memory address ",
			want: users{
				user1: "Manu",
				user2: "Ajit",
				user3: "Pratim",
				user4: "Rohini",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Users(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Users() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_entries_details(t *testing.T) {
	type fields struct {
		Customer_Name string
		Account_Type  string
		Status        string
		AccountNumber int
		Balance       float32
	}
	tests := []struct {
		name   string
		fields fields
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Generates Json File",
			fields: fields{
				Customer_Name: "Manu V H",
				Account_Type:  "Current",
				Status:        "Active",
				AccountNumber: 0012345,
				Balance:       56789.00,
			},
			want: "Error writing json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := entries{
				Customer_Name: tt.fields.Customer_Name,
				Account_Type:  tt.fields.Account_Type,
				Status:        tt.fields.Status,
				AccountNumber: tt.fields.AccountNumber,
				Balance:       tt.fields.Balance,
			}
			e.details()
		})
	}
}
