package repository

import (
	"github.com/mahdi-cpp/api-go-music/config"
	"github.com/mahdi-cpp/api-go-music/model"
	"log"
	"os"
)

func GetMovieMovies() []model.Album {
	var albums []model.Album
	var root = "/var/instagram/"
	var path = "posters/thumbnail_2/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		albums = append(albums, model.Album{Id: 1, AlbumName: e.Name(), AlbumPath: path + e.Name()})
	}
	return albums
}

func GetMovieSeries() []model.Album {

	var albums []model.Album

	var root = "/var/instagram/"
	var path = "posters/series/"

	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		albums = append(albums, model.Album{Id: 1, AlbumName: e.Name(), AlbumPath: path + e.Name()})
	}

	return albums
}

func GetMovieAnimations() []model.Album {

	var albums []model.Album

	var root = "/var/instagram/"
	var path = "/posters/animation/"

	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		albums = append(albums, model.Album{Id: 1, AlbumName: e.Name(), AlbumPath: path + e.Name()})
	}

	return albums
}

func GetMovieYoutube() []model.Album {

	var albums []model.Album

	var root = "/var/instagram/"
	var path = "posters/youtube/"

	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		albums = append(albums, model.Album{Id: 1, AlbumName: e.Name(), AlbumPath: path + e.Name()})
	}

	return albums
}

func GetMovieById(ID uint64) model.Movie {
	var movie model.Movie
	config.DB.Where("id = ?", ID).Find(&movie)
	return movie
}

func GetMovieCastsById(ID uint64) []model.Cast {
	var casts []model.Cast
	config.DB.Find(&casts)
	return casts
}

func GetAllMovies() []model.Movie {
	var movies []model.Movie
	config.DB.
		Offset(0).
		Limit(500).
		Order("id ASC").
		Find(&movies)
	return movies
}

func GetAllPhotos() []model.Photo {

	var photos []model.Photo

	entries, err := os.ReadDir("/var/instagram/messi/")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		photos = append(photos, model.Photo{Id: 1, PhotoPath: e.Name()})
	}

	return photos
}
