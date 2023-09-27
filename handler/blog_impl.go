package handler

import (
	"blogpost/model"
	"blogpost/repo"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	blog    model.BlogArticle
	Comment model.ArticleComments
)

type Blog struct {
	db *gorm.DB
}
type Comments struct {
	db *gorm.DB
}

func BlogImpl(db *gorm.DB) BlogService {
	return Blog{
		db: db,
	}
}

func CommentImpl(db *gorm.DB) CommentService {
	return Comments{
		db: db,
	}
}

// Blog article creation impl
func (B Blog) CreateBlog(r *gin.Context) {
	fmt.Println("here comes")
	var blogs []model.BlogArticle
	if err := r.ShouldBindJSON(&blogs); err != nil {
		fmt.Println("here error")
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("enters")
	err := repo.BlogCreate(B.db, blogs)

	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}
	r.JSON(http.StatusOK, gin.H{"success": "Blogpost created successfully"})
}

// Get all the Blog article
func (B Blog) GetBlog(r *gin.Context) {

	blog, err := repo.GetArticles(B.db)

	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	r.IndentedJSON(http.StatusOK, gin.H{"Result": blog})

}

func (B Blog) GetBlogContent(r *gin.Context) {

	articleID, err := strconv.Atoi(r.Param("articleID"))
	if err != nil {
		// Handle the error if the string cannot be converted to an integer
		fmt.Println("Error:", err)
		return
	}

	articleContent, err := repo.GetArticleContent(B.db, articleID)

	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	r.IndentedJSON(http.StatusOK, gin.H{"Result": articleContent})
}

func (C Comments) CommentArticle(r *gin.Context) {

	if err := r.ShouldBindJSON(&Comment); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := repo.CommentArticle(C.db, Comment, Comment.ArticleID)

	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r.JSON(http.StatusOK, gin.H{"Success": "Comment added article successfully"})
}

func (C Comments) GetArticleComments(r *gin.Context) {
	articleID, err := strconv.Atoi(r.Param("articleID"))
	if err != nil {
		// Handle the error if the string cannot be converted to an integer
		fmt.Println("Error:", err)
		return
	}

	articleContent, err := repo.GetArticleComment(C.db, articleID)

	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if len(articleContent) > 0 {
		r.IndentedJSON(http.StatusOK, gin.H{"Comments in comment are": articleContent})
	} else {
		r.IndentedJSON(http.StatusOK, gin.H{"Comments in comment are": "No records Found"})
	}
}

func (C Comments) CommentOnComment(r *gin.Context) {

	articleID, err := strconv.Atoi(r.Param("articleID"))
	if err != nil {
		// Handle the error if the string cannot be converted to an integer
		fmt.Println("Error:", err)
		return
	}

	commentID, err := strconv.Atoi(r.Param("commentID"))
	if err != nil {
		// Handle the error if the string cannot be converted to an integer
		fmt.Println("Error:", err)
		return
	}

	comment, err := repo.GetCommentOnComment(C.db, articleID, commentID)

	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(comment) > 0 {
		r.JSON(http.StatusOK, gin.H{"Comments in comment are": comment})
	} else {
		r.JSON(http.StatusOK, gin.H{"Comments in comment are": "No records Found"})
	}
}
