package handler

import "github.com/gin-gonic/gin"

type BlogService interface {
	CreateBlog(r *gin.Context)
	GetBlog(r *gin.Context)
	GetBlogContent(r *gin.Context)
}

type CommentService interface {
	CommentArticle(r *gin.Context)
	GetArticleComments(r *gin.Context)
	CommentOnComment(r *gin.Context)
}
