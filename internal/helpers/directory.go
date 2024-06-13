package helpers

import (
	"fmt"
	"os"

	"github.com/breno5g/kmk-cli/config"
	"github.com/breno5g/kmk-cli/pkg/errors"
)

func CheckIfDirExists(path string) bool {
	stat, err := os.Stat(path)
	if errors.ValidError(err) {
		return false
	}

	return stat.IsDir()

}

func GetDirsInside(path string) ([]string, error) {
	logger := config.GetLogger("directory handler")
	var dirs []string

	d, err := os.ReadDir(path)

	if errors.ValidError(err) {
		logger.Error(fmt.Sprintf("error walking through directories: %v", err))
		return nil, err
	}

	for _, dir := range d {
		if dir.IsDir() {
			dirs = append(dirs, dir.Name())
		}
	}

	return dirs, nil
}

func ReverseDirs(dirs []string) []string {
	var reversedDirs []string
	for i := len(dirs) - 1; i >= 0; i-- {
		reversedDirs = append(reversedDirs, dirs[i])
	}

	return reversedDirs
}

func SortDirsByChapters(dirs []string, chapters []string) []string {
	var sortedDirs []string

	for _, chapter := range chapters {
		for _, dir := range dirs {
			if chapter == dir {
				sortedDirs = append(sortedDirs, dir)
			}
		}
	}

	return sortedDirs
}

func CreateDirectory(path string) error {
	logger := config.GetLogger("directory handler")
	err := os.MkdirAll(path, os.ModePerm)

	if errors.ValidError(err) {
		logger.Error(fmt.Sprintf("error creating directory: %v", err))
		return err
	}

	return nil
}
