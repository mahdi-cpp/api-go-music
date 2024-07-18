package model

type Artist struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	OriginalName string `json:"original_name"`
	ProfilePath  string `json:"profile_path"`
}

type Album struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Title2    string `json:"title2"`
	AlbumName string `json:"album_name"`
	AlbumPath string `json:"album_path"`
}

type Menu struct {
	Id    int
	Name  string
	Title []string
	Icon  string
	Type  int
}
