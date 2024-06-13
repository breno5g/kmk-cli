package main

import (
	"fmt"

	"github.com/breno5g/kmk-cli/config"
)

func main() {
	err := config.Init()
	logger := config.GetLogger("main")
	if err != nil {
		logger.Error(fmt.Sprintf("error initializing config: %v", err))
		return
	}

	// var mangas entity.Manga
	// res, err := mangas.GetAllMangas(config.GetDB(), logger)
	// if err != nil {
	// 	logger.Error(fmt.Sprintf("error getting all mangas: %v", err))
	// 	return
	// }

	// for _, manga := range res {
	// 	formatedManga := fmt.Sprintf("Id: %d, Name: %s, Slug: %s", manga.ID, manga.Name.String, manga.Slug)
	// 	fmt.Println(formatedManga)
	// }
}
