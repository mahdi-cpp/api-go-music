package api

import (
	"github.com/gin-gonic/gin"
)

func AddMusicRoutes(rg *gin.RouterGroup) {

	music := rg.Group("/music")

	music.GET("/feed", func(context *gin.Context) {

		//music.GET("/feed", func(context *gin.Context) {
		//
		//	stories := repository.GetViewSliderDTO(
		//		component.STORY,
		//		5,
		//		1,
		//		15, 50,
		//		[]string{"Stories", "داستان ها"},
		//	)
		//	albums := component.GetViewSliderDTO(
		//		component.ALBUM,
		//		6,
		//		3,
		//		15,
		//		0,
		//		[]string{"Albums", "آلبوم ها"},
		//	)
		//	recentlyTracks := component.GetViewSliderDTO(
		//		component.RECENTLY,
		//		8,
		//		4,
		//		15,
		//		0,
		//		[]string{"Recently Played", "اخیرا گوش داده اید"},
		//	)
		//	playlist := component.GetViewSliderDTO(
		//		component.PLAYLIST,
		//		8,
		//		4,
		//		15,
		//		0,
		//		[]string{"Play Lists", "لیست های پخش"},
		//	)
		//	tracks := component.GetViewSliderDTO(
		//		component.TRACK,
		//		8,
		//		6,
		//		15, 8,
		//		[]string{"Tracks", "تک آهنگ ها"},
		//	)

		//context.JSON(210, gin.H{
		//	"storyDTO":         stories,
		//	"albumDTO":         albums,
		//	"recentlyTrackDTO": recentlyTracks,
		//	"playlistDTO":      playlist,
		//	"trackDTO":         tracks,
		//})

	})
}
