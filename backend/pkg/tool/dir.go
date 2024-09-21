package tool

import (
	"os"
)

func CreateFolder(path string) bool {
	err := os.MkdirAll(path, os.ModePerm)

	return err == nil
}
