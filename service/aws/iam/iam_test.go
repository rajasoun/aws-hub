package iam

import "github.com/stretchr/testify/mock"

/**
* Mock using testify Framework
* Technique : Interface Substitution
 */

const testUserName = "test@example.com"
const testAlias = "aws-test-account-alias"
const testARN = "arn:aws:iam::000123456789:user/test@example.com"
const testUserID = "ABCDEFGHIJKLMNOPQRST"

var testUsers = []string{"test1@example.com", "test2@example.com"}

// To mock AWS operations
type MockClient struct {
	mock.Mock
}
