package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/api/status", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("server is running."))
	})
	r.GET("/connect", connect)
	r.Run(":5201") // listen and serve on 0.0.0.0:5201
}
