package mock

import (
	"errors"
	"os"

	iface "github.com/rajasoun/aws-hub/test/e2e/raja/interface"
)

type Mock struct{}

func (mock Mock) OpenFile() iface.OpenFile {
	return func(name string, flag int, perm os.FileMode) (*os.File, error) {
		err := errors.New("simulated err")
		return nil, err
	}
}
