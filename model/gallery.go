package model

type GalleryAlbum struct {
	AlbumName  string
	photoCount uint16
	Photos     []GalleryPhoto
}

type GalleryPhoto struct {
	Photo     string
	Thumbnail string
	Width     uint16
	Height    uint16
}
