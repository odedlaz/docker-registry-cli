package path

import "os"

// FileMode returns the file mode bits of a file
func FileMode(path string) (os.FileMode, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	// file mode bits
	return fileInfo.Mode(), nil
}
