package routes

import "github.com/gin-gonic/gin"

type Route interface {
	Handle(c *gin.Engine)
}
