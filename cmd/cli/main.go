package main

import (
	"fmt"
	"log"
	"os"

	"github.com/breno5g/kmk-cli/config"
	"github.com/urfave/cli/v2"
)

func kmkInit() {
	app := &cli.App{
		Name:  "kmk-cli",
		Usage: "Command line interface to download manga from Komikku",
		Action: func(*cli.Context) error {
			fmt.Println("Hello, World!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := config.Init()
	logger := config.GetLogger("main")
	if err != nil {
		logger.Error(fmt.Sprintf("error initializing config: %v", err))
		return
	}

	kmkInit()

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
