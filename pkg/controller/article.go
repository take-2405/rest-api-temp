package controller

import (
	"backend-record/pkg/model/dao"
	"backend-record/pkg/model/dto"
	"backend-record/pkg/view"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// /read/articles
func ReadArticlesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request dto.ArticleListRequest
		genre := c.Query("genre")
		month := c.Query("month")
		year := c.Query("year")
		request.Year, _ = strconv.Atoi(year)
		request.Month, _ = strconv.Atoi(month)
		request.Genre = genre
		if request.Genre == "" {
			c.JSON(http.StatusBadRequest, "Request is error")
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"Request is error",
			)
		}
		client := dao.MakeReadArticlesClient(request)
		articles, err := client.Request()
		if err != nil {
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"Failed to get the list of articles",
			)
			return
		}
		c.JSON(http.StatusOK, view.ReturnArticleListResponse(&articles))
	}
}

//read/tag/articles
func ReadTagArticlesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request dto.TagRequest
		request.Tag = c.Query("tag")
		if request.Tag == "" {
			log.Println("request tag is nil\n ")
			c.JSON(http.StatusBadRequest, "Request is error")
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"Request is error",
			)
		}
		client := dao.MakeReadTagArticlesClient(request)
		articles, err := client.Request()
		if err != nil {
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"Failed to get a list of articles",
			)
			return
		}
		c.JSON(http.StatusOK, view.ReturnArticleListResponse(&articles))
	}
}

// /read/article
func ReadArticleHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request dto.ArticleDetailRequest
		// リクエストヘッダからUserを取得
		userID := c.GetHeader("UserID")
		if userID == "" {
			log.Println("[ERROR] userID is empty")
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"UserID is empty",
			)
			return
		}
		request.ArticleID = c.Query("articleID")
		if request.ArticleID == "" {
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"Request is error",
			)
			return
		}
		client := dao.MakeReadArticleClient(request)
		article, err := client.Request()
		if err != nil {
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"Failed to get articles from DB",
			)
			return
		}
		article.ArticleUserInfo.Nice, err = client.CheckNiceStatus(userID, request.ArticleID)
		if err != nil {
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"Failed to confirm the status of nices",
			)
			return
		}
		article.ArticleUserInfo.List, err = client.CheckListStatus(userID)
		if err != nil {
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"Failed to confirm the status of lists",
			)
			return
		}
		article.Comments, err = client.GetComments()
		if err != nil {
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"Failed to get comment",
			)
			return
		}
		// 生成した認証トークンを返却
		c.JSON(http.StatusOK, view.ReturnArticleDetailResponse(article))
	}
}
