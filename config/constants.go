package config

import (
	"os"
)

type Paths struct {
	Mangas string
	Ouput  string
}

func InitilizeConstants() Paths {
	manga := os.Getenv("MANGAS_PATH")
	output := os.Getenv("OUTPUT_PATH")

	paths := Paths{
		Mangas: manga,
		Ouput:  output,
	}

	return paths
}
