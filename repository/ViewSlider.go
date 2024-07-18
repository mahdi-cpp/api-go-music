package repository

import (
	"github.com/mahdi-cpp/api-go-music/model"
	"log"
	"os"
)

func Stories() []model.Album {

	var albums []model.Album

	var root = "/var/instagram/"
	var path = "music/face/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		albums = append(albums, model.Album{Id: 1, AlbumName: e.Name(), AlbumPath: path + e.Name()})
	}
	return albums
}

func Albums() []model.Album {

	var albums []model.Album

	var root = "/var/instagram/"
	var path = "music/albums/"

	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		albums = append(albums, model.Album{Id: 1, AlbumName: e.Name(), AlbumPath: path + e.Name(), Title: "Mahdi"})
	}

	return albums
}

func RecentlyTracks() []model.Album {

	var albums []model.Album

	var root = "/var/instagram/"
	var path = "music/recently/"

	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		albums = append(albums, model.Album{Id: 1, AlbumName: e.Name(), AlbumPath: path + e.Name(), Title: "recently", Title2: "Artist"})
	}

	return albums
}

func Playlists() []model.Album {

	var albums []model.Album

	var root = "/var/instagram/"
	var path = "music/playlist/"

	//instagram username: mahdi_ubuntu password:  Mahdi@1234

	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		albums = append(albums, model.Album{Id: 1, AlbumName: e.Name(), AlbumPath: path + e.Name(), Title: "playlist"})
	}

	return albums
}

func Tracks() []model.Album {

	var albums []model.Album

	var root = "/var/instagram/"
	var path = "music/podcast/"

	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		albums = append(albums, model.Album{Id: 1, AlbumName: e.Name(), AlbumPath: path + e.Name(), Title: "Track Name", Title2: "Arist Name"})
	}

	return albums
}
