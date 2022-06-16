package iam

import "github.com/stretchr/testify/mock"

const (
	testUserName = "test@example.com"
	testAlias    = "aws-test-account-alias"
	testARN      = "arn:aws:iam::000123456789:user/test@example.com"
	testUserID   = "ABCDEFGHIJKLMNOPQRST"
	testErrMsg   = "simulated error"
)

/**
* Mock using testify Framework
* Technique : Interface Substitution
 */

// To mock AWS operations
type MockClient struct {
	mock.Mock
}

/**
* Mock via manual creation - Just For Reference
* Technique : Interface Substitution
 */

type MockReciever struct {
	wantErr error
}
