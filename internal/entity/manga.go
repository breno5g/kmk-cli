package entity

import (
	"database/sql"
	"fmt"

	"github.com/breno5g/kmk-cli/config"
	"github.com/breno5g/kmk-cli/pkg/errors"
)

type Manga struct {
	ID               int            `json:"id,omitempty"`
	Slug             string         `json:"slug,omitempty"`
	URL              sql.NullString `json:"url,omitempty"`
	Server_Id        sql.NullString `json:"server_id,omitempty"`
	In_Library       bool           `json:"in_library,omitempty"`
	Name             sql.NullString `json:"name,omitempty"`
	Authors          []uint8        `json:"authors,omitempty"`
	Scanlators       []uint8        `json:"scanlators,omitempty"`
	Genres           []uint8        `json:"genres,omitempty"`
	Synopsis         sql.NullString `json:"synopsis,omitempty"`
	Status           sql.NullString `json:"status,omitempty"`
	Background_Color sql.NullString `json:"background_color,omitempty"`
	Border_Crop      sql.NullString `json:"border_crop,omitempty"`
	Landscape_Zoom   sql.NullInt64  `json:"landscape_zoom,omitempty"`
	Page_Numbering   sql.NullString `json:"page_numbering,omitempty"`
	Reading_mode     sql.NullString `json:"reading_mode,omitempty"`
	Scaling          sql.NullString `json:"scaling,omitempty"`
	ScalingFilter    sql.NullString `json:"scaling_filter,omitempty"`
	Sort_Order       sql.NullString `json:"sort_order,omitempty"`
	Last_Read        sql.NullTime   `json:"last_read,omitempty"`
	Last_Update      sql.NullTime   `json:"last_update,omitempty"`
	Tracking         sql.NullString `json:"tracking,omitempty"`
}

func (m *Manga) GetAllMangas(db *sql.DB, logger *config.Logger) ([]Manga, error) {
	// Get all mangas from database
	query := "SELECT * FROM mangas"
	rows, err := db.Query(query)
	if errors.ValidError(err) {
		logger.Error(fmt.Sprintf("error querying database: %v", err))
		return nil, err
	}

	defer rows.Close()

	var mangas []Manga
	for rows.Next() {
		var manga Manga

		err = rows.Scan(
			&manga.ID,
			&manga.Slug,
			&manga.URL,
			&manga.Server_Id,
			&manga.In_Library,
			&manga.Name,
			&manga.Authors,
			&manga.Scanlators,
			&manga.Genres,
			&manga.Synopsis,
			&manga.Status,
			&manga.Background_Color,
			&manga.Border_Crop,
			&manga.Landscape_Zoom,
			&manga.Page_Numbering,
			&manga.Reading_mode,
			&manga.Scaling,
			&manga.ScalingFilter,
			&manga.Sort_Order,
			&manga.Last_Read,
			&manga.Last_Update,
			&manga.Tracking,
		)

		if errors.ValidError(err) {
			logger.Error(fmt.Sprintf("error scanning database: %v", err))
			return nil, err
		}

		mangas = append(mangas, manga)
	}

	return mangas, nil
}

func (m *Manga) GetById(id int, db *sql.DB) (Manga, error) {
	query := "SELECT * FROM mangas WHERE id = ?"
	row := db.QueryRow(query, id)

	var manga Manga
	err := row.Scan(
		&manga.ID,
		&manga.Slug,
		&manga.URL,
		&manga.Server_Id,
		&manga.In_Library,
		&manga.Name,
		&manga.Authors,
		&manga.Scanlators,
		&manga.Genres,
		&manga.Synopsis,
		&manga.Status,
		&manga.Background_Color,
		&manga.Border_Crop,
		&manga.Landscape_Zoom,
		&manga.Page_Numbering,
		&manga.Reading_mode,
		&manga.Scaling,
		&manga.Sort_Order,
		&manga.Last_Read,
		&manga.Last_Update,
	)

	if errors.ValidError(err) {
		return manga, err
	}

	return manga, nil
}
