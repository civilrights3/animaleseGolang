package filesystem

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func SaveFile(path string, data []byte, overwrite bool) error {
	_, err := os.Stat(path)
	switch {
	case err == nil:
		if overwrite {
			err = os.Remove(path)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("file %v already exists but overwrite set to false", path)
		}
	case errors.Is(err, os.ErrNotExist):
		break
	default:
		return err
	}

	err = os.MkdirAll(filepath.Dir(path), 0666)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0666)
}

func ReadFile(path string) ([]byte, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return b, nil
}
