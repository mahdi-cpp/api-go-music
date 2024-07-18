package repository

import (
	"github.com/mahdi-cpp/api-go-music/model"
	"log"
	"os"
)

func GetGalleryDays() []model.GalleryAlbum {

	var albums []model.GalleryAlbum
	var root = "/var/instagram/"
	var path = "music/ali3/"

	entries, err := os.ReadDir(root + path)

	if err != nil {
		log.Fatal(err)
	}

	var galleryPhotos []model.GalleryPhoto

	for _, e := range entries {

		galleryPhotos = append(galleryPhotos, model.GalleryPhoto{Photo: path + e.Name()})
		if len(galleryPhotos) == 4 {
			albums = append(albums, model.GalleryAlbum{AlbumName: "name", Photos: galleryPhotos})
			galleryPhotos = nil
		}
	}

	return albums
}

func GetGalleryPeoples() []model.Album {

	var albums []model.Album

	var root = "/var/instagram/"
	var path = "music/people/"

	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		albums = append(albums, model.Album{Id: 1, AlbumName: e.Name(), AlbumPath: path + e.Name()})
	}

	return albums
}

func GetGalleryTrips() []model.Album {

	var albums []model.Album

	var root = "/var/instagram/"
	var path = "/music/tinyhome/"

	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		albums = append(albums, model.Album{Id: 1, AlbumName: e.Name(), AlbumPath: path + e.Name()})
	}

	return albums
}

func GetGalleryPlaylists() []model.Album {

	var albums []model.Album

	var root = "/var/instagram/"
	var path = "music/birthaday/"

	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		albums = append(albums, model.Album{Id: 1, AlbumName: e.Name(), AlbumPath: path + e.Name()})
	}

	return albums
}
