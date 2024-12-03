package larvy

import "os"

// CreateDirIfNotExist creates directories in the given path
func (l *Larvy) CreateDirIfNotExist(path string) error {
	const mode = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, mode)
		if err != nil {
			return err
		}
	}
	return nil
}
