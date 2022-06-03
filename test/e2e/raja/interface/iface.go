package iface

import "os"

type OpenFile func(name string, flag int, perm os.FileMode) (*os.File, error)
