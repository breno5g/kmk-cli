package entity

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"

	"github.com/breno5g/kmk-cli/config"
	"github.com/breno5g/kmk-cli/internal/helpers"
	"github.com/breno5g/kmk-cli/pkg/errors"
)

type Chapters struct {
	ID                   sql.NullInt32  `json:"id,omitempty"`
	Manga_id             sql.NullInt32  `json:"manga_id,omitempty"`
	Slug                 sql.NullString `json:"slug,omitempty"`
	Url                  sql.NullString `json:"url" `
	Title                sql.NullString `json:"title,omitempty"`
	Scanlators           []uint8        `json:"scanlators,omitempty"`
	Pages                []uint8        `json:"pages,omitempty"`
	Date                 sql.NullTime   `json:"date,omitempty"`
	Rank                 sql.NullInt32  `json:"rank,omitempty"`
	Downloaded           sql.NullInt32  `json:"downloaded,omitempty"`
	Recent               sql.NullInt32  `json:"recent,omitempty"`
	Read_Progress        sql.NullString `json:"read_progress,omitempty"`
	Read                 sql.NullInt32  `json:"read,omitempty"`
	Last_Page_Read_Index sql.NullInt32  `json:"last_page_read_index,omitempty"`
	Last_read            sql.NullString `json:"last_read,omitempty"`
}

func (c *Chapters) GetAllChapters(db *sql.DB, logger *config.Logger) ([]Chapters, error) {
	// Get all mangas from database
	query := "SELECT * FROM chapters"
	rows, err := db.Query(query)
	if errors.ValidError(err) {
		logger.Error(fmt.Sprintf("error querying database: %v", err))
		return nil, err
	}

	defer rows.Close()

	var mangas []Chapters
	for rows.Next() {
		var manga Chapters

		err = rows.Scan(
			&manga.ID,
			&manga.Manga_id,
			&manga.Slug,
			&manga.Url,
			&manga.Title,
			&manga.Scanlators,
			&manga.Pages,
			&manga.Date,
			&manga.Rank,
			&manga.Downloaded,
			&manga.Recent,
			&manga.Read_Progress,
			&manga.Read,
			&manga.Last_Page_Read_Index,
			&manga.Last_read,
		)

		if errors.ValidError(err) {
			logger.Error(fmt.Sprintf("error scanning rows: %v", err))
			return nil, err
		}

		mangas = append(mangas, manga)
	}

	return mangas, nil
}

func (c *Chapters) GetChaptersByManga(id int, db *sql.DB, logger *config.Logger, firstChapter, lastChapter int) ([]Chapters, error) {
	query := "SELECT * FROM chapters WHERE manga_id = ?"
	rows, err := db.Query(query, id)
	if errors.ValidError(err) {
		logger.Error(fmt.Sprintf("error querying database: %v", err))
		return nil, err
	}

	defer rows.Close()

	var chapters []Chapters
	for rows.Next() {
		var chapter Chapters

		err = rows.Scan(
			&chapter.ID,
			&chapter.Manga_id,
			&chapter.Slug,
			&chapter.Url,
			&chapter.Title,
			&chapter.Scanlators,
			&chapter.Pages,
			&chapter.Date,
			&chapter.Rank,
			&chapter.Downloaded,
			&chapter.Recent,
			&chapter.Read_Progress,
			&chapter.Read,
			&chapter.Last_Page_Read_Index,
			&chapter.Last_read,
		)

		if errors.ValidError(err) {
			logger.Error(fmt.Sprintf("error scanning rows: %v", err))
			return nil, err
		}

		chapters = append(chapters, chapter)
	}

	if firstChapter > 0 && lastChapter > 0 {
		if len(chapters) != 0 {
			var holder []Chapters
			for _, chapter := range chapters {
				re := regexp.MustCompile(`#(\d+)`)

				matches := re.FindAllStringSubmatch(chapter.Title.String, -1)

				if len(matches) > 0 {
					for _, match := range matches {
						pageNumber, err := strconv.Atoi(match[1])
						if errors.ValidError(err) {
							logger.Error(fmt.Sprintf("error converting page number to integer: %v", err))
							return nil, err
						}

						if pageNumber >= firstChapter && pageNumber <= lastChapter {
							holder = append(holder, chapter)
						}
					}
				}
			}

			if len(holder) == 0 {
				return nil, fmt.Errorf("no chapters found for manga with id %d", id)
			}

			return holder, nil
		}
	}

	if len(chapters) == 0 {
		return nil, fmt.Errorf("no chapters found for manga with id %d", id)
	}

	return chapters, nil
}

func (c *Chapters) GetChapterBySlug(slug string, db *sql.DB, logger *config.Logger) (Chapters, error) {
	query := "SELECT * FROM chapters WHERE slug = ?"
	row := db.QueryRow(query, slug)

	var chapter Chapters
	err := row.Scan(
		&chapter.ID,
		&chapter.Manga_id,
		&chapter.Slug,
		&chapter.Url,
		&chapter.Title,
		&chapter.Scanlators,
		&chapter.Pages,
		&chapter.Date,
		&chapter.Rank,
		&chapter.Downloaded,
		&chapter.Recent,
		&chapter.Read_Progress,
		&chapter.Read,
		&chapter.Last_Page_Read_Index,
		&chapter.Last_read,
	)

	if errors.ValidError(err) {
		logger.Error(fmt.Sprintf("error scanning row: %v", err))
		return chapter, err
	}

	return chapter, nil
}

func (c *Chapters) GetBySlug(slug string, chapters []Chapters) (Chapters, error) {
	for _, chapter := range chapters {
		if chapter.Slug.String == slug {
			return chapter, nil
		}
	}

	return Chapters{}, fmt.Errorf("chapter with slug %s not found", slug)
}

func (c *Chapters) GetAllSlugs(chapters []Chapters) []string {
	var names []string
	for _, chapter := range chapters {
		names = append(names, chapter.Slug.String)
	}

	return names
}

func (c *Chapters) Download(manga Manga, chapters []Chapters, logger *config.Logger) error {
	mangaPath := fmt.Sprintf("%s/%s", config.GetPaths().Mangas, manga.Name.String)
	dirs, err := helpers.GetDirsInside(mangaPath)
	if errors.ValidError(err) {
		logger.Error(fmt.Sprintf("error getting directories inside %s: %v", mangaPath, err))
		return err
	}

	previousDownloadedChaptersPath := fmt.Sprintf("%s/%s", config.GetPaths().Ouput, manga.Name.String)
	var previousDownloadedChapters []string
	if helpers.CheckIfDirExists(previousDownloadedChaptersPath) {
		previousDownloadedChapters, err = helpers.GetDirsInside(previousDownloadedChaptersPath)

		if errors.ValidError(err) {
			logger.Error(fmt.Sprintf("error getting directories inside %s: %v", previousDownloadedChaptersPath, err))
			return err
		}
	} else {
		helpers.CreateDirectory(previousDownloadedChaptersPath)
	}

	if len(dirs) == 0 {
		logger.Info(fmt.Sprintf("no chapters to download for manga %s", manga.Name.String))
		return nil
	}

	sortedDirs := helpers.SortDirsByChapters(dirs, c.GetAllSlugs(chapters))

	for _, dir := range sortedDirs {
		if helpers.Contains(previousDownloadedChapters, dir) {
			logger.Info(fmt.Sprintf("chapter %s already downloaded", dir))
			continue
		}

		chapter, err := c.GetBySlug(dir, chapters)
		if errors.ValidError(err) {
			logger.Error(fmt.Sprintf("error getting chapter by slug %s: %v", dir, err))
			return err
		}

		outputPath := previousDownloadedChaptersPath
		if helpers.CheckIfDirExists(fmt.Sprintf("%s/%s", outputPath, chapter.Title.String)) {
			logger.Info(fmt.Sprintf("chapter %s already downloaded", chapter.Title.String))
			continue
		} else {
			err = helpers.CreateDirectory(fmt.Sprintf("%s/%s", outputPath, chapter.Title.String))
			if errors.ValidError(err) {
				logger.Error(fmt.Sprintf("error creating directory %s: %v", chapter.Title.String, err))
				return err
			}
		}

		files, err := helpers.GetDirContent(fmt.Sprintf("%s/%s", mangaPath, dir))
		if errors.ValidError(err) {
			logger.Error(fmt.Sprintf("error getting directory content: %v", err))
			return err
		}

		logger.Infof("moving %s \n", chapter.Title.String)
		for _, file := range files {
			err = helpers.MoveDirContent(fmt.Sprintf("%s/%s/%s", mangaPath, dir, file), fmt.Sprintf("%s/%s/%s", outputPath, chapter.Title.String, file))
			if errors.ValidError(err) {
				logger.Error(fmt.Sprintf("error moving directory content: %v", err))
				return err
			}
		}
	}

	return nil
}
