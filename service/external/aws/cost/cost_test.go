package cost

import "github.com/stretchr/testify/mock"

/**
* Mock using testify Framework
* Technique : Interface Substitution
 */

// To mock AWS operations.
type MockClient struct {
	mock.Mock
}
