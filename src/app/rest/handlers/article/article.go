package article

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahdimehrabi/bm-apigateway/src/internal/entity"
	"github.com/mahdimehrabi/bm-articles/src/apps/grpc/proto/article"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

const articleServer = "localhost:6001"

type Handler struct {
}

func (h Handler) Create(c *gin.Context) {
	artEnt := entity.Article{}

	if err := c.ShouldBindJSON(&artEnt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	cc, err := grpc.Dial(articleServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	defer cc.Close()

	client := article.NewArticleClient(cc)
	r := article.ArticleReq{
		Title:  artEnt.Title,
		Body:   artEnt.Body,
		Price:  artEnt.Price,
		Tags:   artEnt.Tags,
		UserID: artEnt.UserID,
	}
	_, err = client.CreateArticle(context.Background(), &r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
