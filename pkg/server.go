package pkg

import (
	"backend-record/pkg/controller"
	"github.com/gin-gonic/gin"
)

var (
	//Server gin flameworkのserver
	Server *gin.Engine
)

func init() {
	Server = gin.Default()
	//記事の一覧情報送信
	Server.GET("/read/articles", controller.ReadArticlesHandler())
	//記事の詳細情報送信
	Server.GET("/read/article", controller.ReadArticleHandler())
	//記事のタグ検索
	Server.GET("/read/tag/articles", controller.ReadTagArticlesHandler())
	//記事のいいね数更新
	Server.POST("/update/add/like", controller.UpdateAddLikeHandler())
}
