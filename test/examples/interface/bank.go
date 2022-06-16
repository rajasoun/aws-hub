package spike

import (
	"fmt"
	"log"
	"os"
)

type BankAccountDetails interface {
	account()
}

type acc_details struct {
	acc_holder_name string
	acc_number      int
	acc_type        string
	acc_balance     float32
}

func (ad *acc_details) getaccountnamenumber() (string, int) {
	accname := ad.acc_holder_name
	accnumber := ad.acc_number
	fmt.Printf("Account Holder Name: %s\n", accname)
	fmt.Printf("Account Number: %d\n", accnumber)
	return ad.acc_holder_name, ad.acc_number
}

func (ad *acc_details) getaccounttypebalance() (string, float32) {
	acctype := ad.acc_type
	accbalance := ad.acc_balance
	fmt.Printf("Account Type: %s\n", acctype)
	fmt.Printf("Account Balance: %f\n", accbalance)
	return ad.acc_type, ad.acc_balance
}

func AccoutDetails() *acc_details {
	AccDetails := acc_details{
		acc_holder_name: "Manu V H",
		acc_number:      543210,
		acc_type:        "Savings",
		acc_balance:     5678.95,
	}
	return &AccDetails
}

func Statement() {
	file, err := os.Create("Statement.txt")
	if err != nil {
		log.Printf("Error generating Statement: %v", err)
	} else {
		log.Printf("Statement generated successfully: %v", file)
	}
}
