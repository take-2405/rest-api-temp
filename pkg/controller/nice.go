package controller

import (
	"backend-record/pkg/model/dao"
	"backend-record/pkg/model/dto"
	"backend-record/pkg/view"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// /update/add/like
func UpdateAddLikeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request dto.NiceRequest
		var nice    *dto.Nice
		userID := c.GetHeader("UserID")
		if userID==""{
			log.Println("[ERROR] userID is empty")
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"UserID is empty",
			)
			return
		}
		err:=c.ShouldBindJSON(&request)
		if err != nil {
			log.Println("[ERROR] Faild Bind JSON")
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"Request is error",
			)
			return
		}
		client:=dao.MakeUpdateAddLikeClient(request)
		nice, err = client.Request(userID)
		if err != nil {
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"DB operation failed",
			)
			return
		}
		// // 生成した認証トークンを返却
		c.JSON(http.StatusOK, view.ReturnNiceResopnse(nice))
	}
}
