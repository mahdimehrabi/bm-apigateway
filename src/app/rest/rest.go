package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdimehrabi/bm-apigateway/src/app/rest/routes/article"
)

func RunServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	a := article.NewRoute()
	a.Handle(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
