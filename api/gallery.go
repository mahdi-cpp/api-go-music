package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-music/repository"
)

func AddGalleryRoutes(rg *gin.RouterGroup) {

	music := rg.Group("/gallery")

	music.GET("/feed", func(context *gin.Context) {

		days := repository.GetGalleryDays()
		peoples := repository.GetGalleryPeoples()
		trips := repository.GetGalleryTrips()
		playlists := repository.GetGalleryPlaylists()

		context.JSON(210, gin.H{
			"days":      days,
			"peoples":   peoples,
			"trips":     trips,
			"playlists": playlists,
		})
	})
}
