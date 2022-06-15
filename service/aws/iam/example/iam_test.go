package example

/**
* manual Mocking of AWS APIs
* Technique : Interface Substitution
 */

const testUserName = "test@example.com"
const testAlias = "aws-test-account-alias"
const testARN = "arn:aws:iam::000123456789:user/test@example.com"
const testUserID = "ABCDEFGHIJKLMNOPQRST"

var testUsers = []string{"test1@example.com", "test2@example.com"}

/**
* Mock via manual creation - Just For Reference
* Technique : Interface Substitution
 */

type MockReciever struct {
	wantErr error
}
