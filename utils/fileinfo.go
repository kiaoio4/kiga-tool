package utils

import (
	"os"
	"path/filepath"
)

func GetAllFilesList(dirpath string) ([]string, error) {
	var dirlist []string
	direrr := filepath.Walk(dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if path == dirpath {
				return nil
			}
			dirlist = append(dirlist, path)
			return nil
		})
	return dirlist, direrr
}
