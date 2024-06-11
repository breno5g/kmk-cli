package main

import (
	"fmt"

	"github.com/breno5g/kmk-cli/config"
	"github.com/breno5g/kmk-cli/internal/entity"
)

func main() {
	err := config.Init()
	logger := config.GetLogger("main")
	if err != nil {
		logger.Error(fmt.Sprintf("error initializing config: %v", err))
		return
	}

	var manga entity.Chapters
	manga.GetMangas(config.GetDB(), logger)
}
