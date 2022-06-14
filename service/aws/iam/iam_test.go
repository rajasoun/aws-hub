package iam

import "github.com/stretchr/testify/mock"

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
