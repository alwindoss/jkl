package jkl

import "path/filepath"

var (
	DefaultDBEnvironment = filepath.Join(CacheDir, "objects")
)

const (
	DBEnvironment = "SHA1_FILE_DIRECTORY"
	CacheDir      = ".dircache"
)
