package entity

import (
	"database/sql"
	"fmt"

	"github.com/breno5g/kmk-cli/config"
)

type Chapters struct {
	ID                   sql.NullInt32  `json:"id,omitempty"`
	Manga_id             sql.NullInt32  `json:"manga_id,omitempty"`
	Slug                 sql.NullString `json:"slug,omitempty"`
	Url                  sql.NullString `json:"url" `
	Title                sql.NullString `json:"title,omitempty"`
	Scanlators           []uint8        `json:"scanlators,omitempty"`
	Pages                []uint8        `json:"pages,omitempty"`
	Date                 sql.NullString `json:"date,omitempty"`
	Rank                 sql.NullInt32  `json:"rank,omitempty"`
	Downloaded           sql.NullInt32  `json:"downloaded,omitempty"`
	Recent               sql.NullInt32  `json:"recent,omitempty"`
	Read_Progress        sql.NullInt32  `json:"read_progress,omitempty"`
	Read                 sql.NullInt32  `json:"read,omitempty"`
	Last_Page_Read_Index sql.NullInt32  `json:"last_page_read_index,omitempty"`
	Last_read            sql.NullString `json:"last_read,omitempty"`
}

func (manga *Chapters) GetMangas(db *sql.DB, logger *config.Logger) {
	// Get all mangas from database
	query := "SELECT * FROM chapters"
	rows, err := db.Query(query)
	if err != nil {
		logger.Error(fmt.Sprintf("error querying database: %v", err))
		return
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

		if err != nil {
			logger.Error(fmt.Sprintf("error scanning rows: %v", err))
			return
		}

		mangas = append(mangas, manga)
	}

	for _, manga := range mangas {
		formatedManga := fmt.Sprintf("ID: %d, Manga_id: %v, Slug: %v, Title: %v, Scans: %v \n", manga.ID.Int32, manga.Manga_id.Int32, manga.Slug.String, manga.Title, manga.Scanlators)
		fmt.Println(formatedManga)
	}
}
