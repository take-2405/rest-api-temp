package view

import (
	"backend-record/pkg/model/dto"
	"github.com/gin-gonic/gin"
)


//Response
type ArticleListResponse struct{
	Articles *[]dto.Article `json:"articles"`
}

func ReturnArticleListResponse(articles *[]dto.Article)ArticleListResponse{
	//body:=ArticleListResponse{Articles: articles}
	return 	ArticleListResponse{Articles: articles}
}

type ArticleDetailResponse struct{
	ArticleDetail *dto.ArticleDetail `json:"articleDetail"`
}

func ReturnArticleDetailResponse(articleDetail *dto.ArticleDetail)ArticleDetailResponse{
	return ArticleDetailResponse{ArticleDetail: articleDetail}
}

type Error struct {
	Code		int		`json:"code"`
	Message		string	`json:"message"`
	Description	string	`json:"description"`
}

func ReturnErrorResponse(c *gin.Context, code int, msg, desc string) {
	body := Error{
		Code: code,
		Message: msg,
		Description: desc,
	}
	c.JSON(code, body)
}