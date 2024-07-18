package repository

import (
	"github.com/mahdi-cpp/api-go-music/model"
	"log"
	"os"
)

func GetMusicArtists() []model.Artist {

	var artists []model.Artist

	entries, err := os.ReadDir("/var/instagram/artists/")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		artists = append(artists, model.Artist{Id: 1, ProfilePath: e.Name()})
	}

	return artists
}

func GetMusicStories() []model.Album {
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

func GetMusicAlbums() []model.Album {

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

func GetMusicPlaylists() []model.Album {

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

func GetMusicRecentlyTracks() []model.Album {

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

func GetMusicTracks() []model.Album {

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

func GetMusicMenu() []model.Menu {
	var menus []model.Menu

	menus = append(menus, model.Menu{Id: 1, Icon: "icon1.png", Name: "A"})
	menus = append(menus, model.Menu{Id: 1, Icon: "icon1.png", Name: "A"})
	menus = append(menus, model.Menu{Id: 1, Icon: "icon1.png", Name: "A"})

	return menus
}
