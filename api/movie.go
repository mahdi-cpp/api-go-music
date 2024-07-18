package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-music/repository"
	"strconv"
)

func AddMovieRoutes(rg *gin.RouterGroup) {

	movie := rg.Group("/movie")

	movie.GET("/feed", func(context *gin.Context) {

		movies := repository.GetMovieMovies()
		series := repository.GetMovieSeries()
		animations := repository.GetMovieAnimations()
		youtubes := repository.GetMovieYoutube()

		context.JSON(210, gin.H{
			"movies":     movies,
			"series":     series,
			"animations": animations,
			"youtubes":   youtubes,
		})
	})

	movie.GET("/movie", func(context *gin.Context) {
		movieId, _ := strconv.ParseUint(context.Query("movie_id"), 0, 64)
		movie := repository.GetMovieById(movieId)
		casts := repository.GetMovieCastsById(movieId)

		context.JSON(210, gin.H{"movie": movie, "casts": casts})
	})

	movie.GET("/movie_casts", func(context *gin.Context) {
		movieId, _ := strconv.ParseUint(context.Query("movie_id"), 0, 64)
		casts := repository.GetMovieCastsById(movieId)
		//fmt.Println(casts[0].Name)
		context.JSON(210, casts)
	})

	movie.GET("/movies", func(context *gin.Context) {
		movies := repository.GetAllMovies()
		fmt.Println(movies)
		context.JSON(210, gin.H{"movies": movies})
	})

	movie.GET("/photos", func(context *gin.Context) {
		photos := repository.GetAllPhotos()
		//fmt.Println(gin.H{"photos": photos})
		context.JSON(210, gin.H{"photos": photos})
	})

}
