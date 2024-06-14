package helpers

import (
	"fmt"
	"io"
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

func MoveDirContent(src, dst string) error {
	logger := config.GetLogger("directory handler")
	sourceFile, err := os.Open(src)

	if errors.ValidError(err) {
		logger.Error(fmt.Sprintf("error opening source file: %v", err))
		return err
	}
	defer sourceFile.Close()

	destination, err := os.Create(dst)

	if errors.ValidError(err) {
		logger.Error(fmt.Sprintf("error creating destination file: %v", err))
		return err
	}

	defer destination.Close()

	_, err = io.Copy(destination, sourceFile)

	if errors.ValidError(err) {
		logger.Error(fmt.Sprintf("error moving directory content: %v", err))
		return err
	}

	return nil
}

func GetDirContent(path string) ([]string, error) {
	logger := config.GetLogger("directory handler")

	var files []string

	c, err := os.ReadDir(path)

	if errors.ValidError(err) {
		logger.Error(fmt.Sprintf("error walking through directories: %v", err))
		return nil, err
	}

	for _, file := range c {
		files = append(files, file.Name())
	}

	return files, nil
}
