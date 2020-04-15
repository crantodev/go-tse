package file

import "os"

// Exists will check if the file is present in the location or not
func Exists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
