package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	func server() {
		router := gin.Default()
		router.LoadHTMLFiles("index.html")
		router.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		})
		router.Run(":8080")
	}
*/

func server2() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/index", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"img": "https://images.igdb.com/igdb/image/upload/t_cover_big_2x/co5rav.jpg",
		})
	})
	router.Run(":8080")
}
