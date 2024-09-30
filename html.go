package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"net/http"
)
import _ "net/http"

func server() {
	router := gin.Default()
	var data = readData("response_body.json")

	router.GET("/games", func(c *gin.Context) {
		c.JSON(http.StatusOK, data)
	})
	router.Run(":8080")
}
