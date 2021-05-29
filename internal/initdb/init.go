package initdb

import (
	"fmt"
	"os"

	"github.com/alwindoss/jkl"
	"github.com/spf13/afero"
)

func createCacheDir(fs afero.Fs) error {
	err := fs.Mkdir(jkl.CacheDir, 0700)
	if err != nil {
		err = fmt.Errorf("%v: %w", jkl.ErrCacheDirAlreadyExists, err)
		return err
	}
	return nil
}

func Run() error {
	initFS := afero.NewOsFs()
	err := createCacheDir(initFS)
	if err != nil {
		fmt.Printf("unabe to create %s\n", jkl.CacheDir)
		os.Exit(1)
	}
	path := os.Getenv(jkl.DBEnvironment)

	if path != "" {
		file, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s set to a bad directory\n", jkl.DBEnvironment)
		}
		fileInfo, err := file.Stat()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s set to a bad directory\n", jkl.DBEnvironment)
		}
		if fileInfo.IsDir() {
			return nil
		}
	}

	sha1Dir := jkl.DefaultDBEnvironment
	err = createDefaultDBEnv(initFS, sha1Dir)
	if err != nil {
		err = fmt.Errorf("unable to createDefaultDBEnv: %w", err)
		return err
	}
	return nil
}

func createDefaultDBEnv(fs afero.Fs, loc string) error {
	path := loc
	for i := 0; i < 256; i++ {
		dirpath := fmt.Sprintf("%s/%02x", path, i)
		err := fs.MkdirAll("./"+dirpath, 0700)
		if err != nil {
			err = fmt.Errorf("error doing fs.MkdirAll: %w", err)
			return err
		}
	}
	return nil
}
