package system

import (
	"os"
)

func Exist(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}

func NotExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

