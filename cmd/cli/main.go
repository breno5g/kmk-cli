package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/breno5g/kmk-cli/config"
	"github.com/breno5g/kmk-cli/internal/entity"
	"github.com/urfave/cli/v2"
)

func kmkInit() {
	logger := config.GetLogger("cli")
	db := config.GetDB()

	app := &cli.App{
		Name:  "kmk-cli",
		Usage: "Command line interface to download manga from Komikku",
		Commands: []*cli.Command{
			{
				Name:    "Get mangas",
				Aliases: []string{"gm"},
				Usage:   "Get all mangas with id and status",
				Action: func(ctx *cli.Context) error {
					var mangas entity.Manga
					res, err := mangas.GetAllMangas(db, logger)
					if err != nil {
						logger.Errorf(fmt.Sprintf("Error getting all mangas: %v", err))
						return err
					}

					fmt.Println("Mangas")
					for _, manga := range res {
						formatedManga := fmt.Sprintf("%d - %s - %s", manga.ID, manga.Name.String, manga.Status.String)
						fmt.Println(formatedManga)
					}
					return nil
				},
			},
			{
				Name:    "Get Chapters",
				Aliases: []string{"gc"},
				Usage:   "Get all manga chapters",
				Action: func(ctx *cli.Context) error {
					mangaId, err := strconv.Atoi(ctx.Args().First())
					if err != nil {
						logger.Errorf(fmt.Sprintf("Please pass a valid manga id: %v", err))
						return err
					}
					var chapters entity.Chapters
					res, err := chapters.GetChaptersByManga(mangaId, db, logger)
					if err != nil {
						logger.Error(err)
						return nil
					}

					fmt.Println("Chapters")
					for _, chapter := range res {
						formatedDate := fmt.Sprintf("%d/%02d/%d", chapter.Date.Time.Year(), int(chapter.Date.Time.Month()), chapter.Date.Time.Day())

						formatedChapter := fmt.Sprintf("%d - %s - %s", chapter.ID.Int32, chapter.Title.String, formatedDate)
						fmt.Println(formatedChapter)
					}
					return nil
				},
			},
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

}
