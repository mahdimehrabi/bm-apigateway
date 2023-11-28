package article

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdimehrabi/bm-apigateway/src/app/rest/handlers/article"
	"github.com/mahdimehrabi/bm-apigateway/src/app/rest/routes"
)

type route struct {
}

func NewRoute() routes.Route {
	return &route{}
}

func (r route) Handle(c *gin.Engine) {
	handler := article.Handler{}
	g := c.Group("/articles")
	{
		g.POST("/", handler.Create)
	}
}
