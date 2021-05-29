package jkl

import (
	"fmt"
)

var (
	ErrCacheDirAlreadyExists = fmt.Errorf("the cache directory %s already exists", CacheDir)
)
