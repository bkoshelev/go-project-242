package code

import (
	"errors"
	"fmt"
	"os"
)

func GetSize(path string) (int64, error) {
	fileInfo, err := os.Lstat(path)

	if err != nil {
		return 0, errors.New("path does not exist")
	}

	if fileInfo.IsDir() == false {
		return fileInfo.Size(), nil
	}

	dirEntries, err := os.ReadDir(path)

	if err != nil {
		return 0, errors.New("directory reading error")
	}

	var dirSize int64

	for _, entry := range dirEntries {
		if entry.IsDir() == true {
			continue
		}

		info, err := entry.Info()

		if err != nil {
			return 0, fmt.Errorf("directory reading error: %w", err)
		}

		dirSize += info.Size()
	}

	return dirSize, nil
}

func GetPathSize(path string) (string, error) {
	size, e := GetSize(path)

	if e != nil {
		return "", fmt.Errorf("GetSize error: %w", e)
	}

	return fmt.Sprintf("%d\t%s", size, path), nil
}
