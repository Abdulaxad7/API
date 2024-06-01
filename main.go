package main

import (
	"ALBUMS/config"
	"ALBUMS/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	_ = godotenv.Load()
	_, err := config.OpenDB()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.GET("/albums", routers.GetAlbums)
	router.GET("/album/:id", routers.GetAlbumById)
	router.POST("/albums", routers.PostAlbums)
	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

}
