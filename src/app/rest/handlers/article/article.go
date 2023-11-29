package article

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	articleGRPC "github.com/mahdimehrabi/bm-apigateway/src/external/article"
	"github.com/mahdimehrabi/bm-apigateway/src/internal/entity"
	"github.com/mahdimehrabi/bm-articles/src/apps/grpc/proto/article"
	"net/http"
)

type Handler struct {
}

func (h Handler) Create(c *gin.Context) {
	artEnt := entity.Article{}

	if err := c.ShouldBindJSON(&artEnt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	aGrpc := articleGRPC.ArticleGRPC{}
	err := aGrpc.Connect()
	defer aGrpc.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	cc := aGrpc.CC
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
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h Handler) GetAll(c *gin.Context) {
	aGRPC := articleGRPC.ArticleGRPC{}
	err := aGRPC.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer aGRPC.Close()
	client := article.NewArticleClient(aGRPC.CC)
	articles, err := client.GetArticles(context.Background(), &article.Empty{})
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	articleEnts := make([]entity.Article, 0)
	for _, art := range articles.Articles {
		articleEnts = append(articleEnts, entity.Article{
			ID:       art.ID,
			Title:    art.Title,
			Body:     art.Body,
			Price:    art.Price,
			Tags:     art.Tags,
			BuyCount: art.BuyCount,
		})
	}
	c.JSON(http.StatusOK, articleEnts)
}
