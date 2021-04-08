package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	root := gin.Default()

	// root.GET("/ping",func(ctx *gin.Context)}{

	// })

	root.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ping",
		})
	})

	root.Run()

}
