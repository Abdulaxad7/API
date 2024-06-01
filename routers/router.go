package routers

import (
	"ALBUMS/album_data"
	"ALBUMS/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, config.SelectAll())
}
func PostAlbums(context *gin.Context) {
	var newAlbum album_data.Album
	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}
	albums := config.SelectAll()
	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, albums)

}
func GetAlbumById(context *gin.Context) {
	id := context.Param("id")
	for _, v := range config.SelectAll() {
		if v.Id == id {
			context.IndentedJSON(http.StatusOK, v)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
