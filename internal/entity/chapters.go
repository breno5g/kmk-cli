package entity

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"

	"github.com/breno5g/kmk-cli/config"
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

func (c *Chapters) Download(chapters []Chapters, logger *config.Logger) error {
	for _, chapter := range chapters {
		logger.Info(fmt.Sprintf("Downloading chapter %s", chapter.Title.String))
	}

	return nil
}
