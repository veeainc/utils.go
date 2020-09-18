package files

import (
	"io/ioutil"
	"os"
)

// Checks if a file exists and is not a directory.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Copy a file from source to dest.
func CopyFile(source, dest string) error {
	input, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dest, input, 0644)
	if err != nil {
		return err
	}

	return nil
}
