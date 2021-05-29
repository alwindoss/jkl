package initdb

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alwindoss/jkl"
	"github.com/spf13/afero"
)

func TestCreateCacheDir(t *testing.T) {
	osFS := afero.NewMemMapFs()
	err := createCacheDir(osFS)
	if err != nil {
		t.Errorf("expected to create cache dir successfully but found error: %v", err)
		t.FailNow()
	}
	info, err := osFS.Stat(jkl.CacheDir)
	if err != nil {
		t.Errorf("expected no error when getting the fileinfo but got: %v", err)
		t.FailNow()
	}
	if !info.IsDir() {
		t.Errorf("expected the create cache to be a dir but it was not: %v", err)
		t.FailNow()
	}

}

func TestCreateCacheDirWhenCacheDirExistsAlready(t *testing.T) {
	osFs := afero.NewMemMapFs()
	err := osFs.Mkdir(jkl.CacheDir, 0700)
	if err != nil {
		t.Errorf("unable to setup the test: %v", err)
		t.FailNow()
	}
	err = createCacheDir(osFs)
	if err == nil {
		t.Errorf("expected createCacheDir to return error but the err was nil")
		t.FailNow()
	}
	if !errors.Is(err, afero.ErrFileExists) && !errors.Is(err, jkl.ErrCacheDirAlreadyExists) {
		t.Errorf("expected the error to be %v but it was %v", afero.ErrFileExists, err)
		t.FailNow()
	}
}

func TestCreateDefaultDBEnv(t *testing.T) {
	memFS := afero.NewMemMapFs()
	err := createDefaultDBEnv(memFS, jkl.DefaultDBEnvironment)
	if err != nil {
		t.Errorf("expected no error but found: %v", err)
		t.FailNow()
	}
	// path := filepath.Join(jkl.DefaultDBEnvironment, "objects")
	for i := 0; i < 256; i++ {
		checkPath := fmt.Sprintf("%s/%02x", jkl.DefaultDBEnvironment, i)
		f, err := memFS.Open(checkPath)
		if err != nil {
			t.Errorf("expected no error when opening the path but found: %v", err)
			t.FailNow()
		}
		fi, err := f.Stat()
		if err != nil {
			t.Errorf("expected no error when doing f.Stat but found: %v", err)
			t.FailNow()
		}
		if !fi.IsDir() {
			t.Errorf("expected this to be a file but it was found to be a directory")
			t.FailNow()
		}

	}
}
