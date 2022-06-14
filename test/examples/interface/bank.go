package spike


type BankAccountDetails interface {
	account()
}

type acc_details struct{
	acc_holder_name string
	acc_number int
	acc_type string
	acc_balance float32
}


func (ad *acc_details) account() (string, int, string,float32) {
	return ad.acc_holder_name,ad.acc_number,ad.acc_type,ad.acc_balance
}
