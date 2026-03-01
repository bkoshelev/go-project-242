package code

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
	PB = 1024 * TB
	EB = 1024 * PB
)

func GetSize(path string, all bool, recursive bool) (int64, error) {
	fileInfo, err := os.Lstat(path)

	if err != nil {
		return 0, errors.New("path does not exist")
	}

	if !fileInfo.IsDir() {
		return fileInfo.Size(), nil
	}

	dirEntries, err := os.ReadDir(path)

	if err != nil {
		return 0, errors.New("directory reading error")
	}

	var dirSize int64

	for _, entry := range dirEntries {
		if entry.IsDir() && !recursive {
			continue
		}

		if entry.IsDir() && recursive {
			internalDirSize, err := GetSize(filepath.Join(path, entry.Name()), all, recursive)

			if err != nil {
				return 0, errors.New("directory reading error")
			}

			dirSize += internalDirSize
		} else {
			info, err := entry.Info()

			if err != nil {
				return 0, fmt.Errorf("directory reading error: %w", err)
			}

			if !all && strings.HasPrefix(info.Name(), ".") {
				continue
			}

			dirSize += info.Size()
		}

	}

	return dirSize, nil
}

func FormatSize(size int64, human bool) string {
	if size <= 0 {
		return "0B"
	}

	if !human || float64(size) < KB {
		return fmt.Sprintf("%dB", size)
	}

	if float64(size) < MB {
		kbSize := float64(size) / KB
		return fmt.Sprintf("%.1fKB", kbSize)
	}

	if float64(size) < GB {
		mbSize := float64(size) / MB
		return fmt.Sprintf("%.1fMB", mbSize)
	}

	if float64(size) < TB {
		gbSize := float64(size) / GB
		return fmt.Sprintf("%.1fGB", gbSize)
	}

	if float64(size) < PB {
		tbSize := float64(size) / TB
		return fmt.Sprintf("%.1fTB", tbSize)
	}

	if float64(size) < EB {
		pbSize := float64(size) / PB
		return fmt.Sprintf("%.1fPB", pbSize)
	}

	ebSize := float64(size) / EB
	return fmt.Sprintf("%.1fEB", ebSize)
}

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, e := GetSize(path, all, recursive)

	if e != nil {
		return "", fmt.Errorf("GetSize error: %w", e)
	}

	fmtSize := FormatSize(size, human)

	return fmtSize, nil
}
