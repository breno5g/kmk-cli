package handlers

import (
	"fmt"
	"os"

	"github.com/breno5g/kmk-cli/config"
	"github.com/breno5g/kmk-cli/pkg/errors"
)

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
