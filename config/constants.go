package config

import "os"

type Paths struct {
	Mangas string
}

func InitilizeConstants() Paths {
	var paths Paths
	path := os.Getenv("MANGAS_PATH")
	paths.Mangas = path
	return paths
}
