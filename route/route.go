package route

import (
	db "blogpost/db_connection"
	"blogpost/handler"

	"github.com/gin-gonic/gin"
)

var (
	Blogrec    = handler.BlogImpl(db.Db)
	Commentrec = handler.CommentImpl(db.Db)
)

func InitializeRouter() {
	router := gin.Default()
	router.POST("/create_article", Blogrec.CreateBlog)
	router.GET("/getArticles", Blogrec.GetBlog)
	router.GET("/getArticleContent/:articleID", Blogrec.GetBlogContent)
	router.POST("/commentArticle/", Commentrec.CommentArticle)
	router.GET("/commentoncomment/:articleID/:commentID", Commentrec.CommentOnComment)
	router.GET("/getArticleComments/:articleID", Commentrec.GetArticleComments)
	router.Run("localhost:8080")
}
