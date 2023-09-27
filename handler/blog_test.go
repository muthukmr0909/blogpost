package handler_test

import (
	"blogpost/handler"
	"blogpost/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = "3306"
	user     = "root"
	password = ""
	dbName   = "blog_test"
)

func setupTestDB() (*gorm.DB, error) {
	dsn := user + "@(" + host + ":" + port + ")/" + dbName + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err.Error(),
			"service": "blogpost",
		}).Warn("failed to connect to database")
		return db, err

	}
	db.AutoMigrate(&model.BlogArticle{})
	db.AutoMigrate(&model.ArticleComments{})
	return db, nil
}

func TestGetBlog(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	r := gin.Default()
	Bloghandler := handler.BlogImpl(db)

	r.GET("/getArticles", Bloghandler.GetBlog)
	req, err := http.NewRequest("GET", "/getArticles", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	response := model.BlogArticle{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	if reflect.DeepEqual([]model.BlogArticle{}, response) {
		t.Errorf("Error on response")
	}
}

func JSONStringify(data interface{}) string {
	byteData, _ := json.Marshal(data)
	stringData := string(byteData)
	return stringData
}
func TestCreateBlog(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	r := gin.Default()
	handler := handler.BlogImpl(db)
	assert.NoError(t, err)
	r.POST("/create_article", handler.CreateBlog)
	PostArticle := []model.BlogArticle{
		{
			ArticleID:      13,
			ArticleTitle:   "Title.",
			ArticleContent: "This is the new blog",
			Nickname:       "mine",
		},
	}
	strData := JSONStringify(PostArticle)
	req, err := http.NewRequest("POST", "/create_article", strings.NewReader(strData))
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	expected := map[string]interface{}{
		"success": "Blogpost created successfully",
	}
	err = json.Unmarshal(w.Body.Bytes(), &expected)
	assert.NoError(t, err)
	if reflect.DeepEqual(response, expected) {
		t.Errorf("Error on response")
	}

}

func TestGetBlogContent(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	r := gin.Default()
	handler := handler.BlogImpl(db)
	assert.NoError(t, err)
	r.GET("/getArticleContent/:articleID", handler.GetBlogContent)
	req, err := http.NewRequest("GET", "/getArticleContent/4", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	response := model.BlogArticle{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	if reflect.DeepEqual(model.BlogArticle{}, response) {
		t.Errorf("Error on response")
	}
}

func TestCommentArticle(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	d, _ := db.DB()
	defer d.Close()
	r := gin.Default()
	handler := handler.CommentImpl(db)
	assert.NoError(t, err)
	r.POST("/commentArticle", handler.CommentArticle)
	addComment := model.ArticleComments{
		ArticleID:       13,
		CommentID:       13,
		ParentCommentID: 1,
		Content:         "This is a sample article content.",
		Nickname:        "super",
	}
	strData := JSONStringify(addComment)
	req, err := http.NewRequest("POST", "/commentArticle", strings.NewReader(strData))
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	expected := map[string]interface{}{
		"Success": "Comment added article successfully",
	}
	assert.Equal(t, http.StatusOK, w.Code)
	err = json.Unmarshal(w.Body.Bytes(), &expected)
	assert.NoError(t, err)
	if reflect.DeepEqual(response, expected) {
		t.Errorf("Error on response")
	}

}

func TestCommentOnComment(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	d, _ := db.DB()
	defer d.Close()
	r := gin.Default()
	handler := handler.CommentImpl(db)
	assert.NoError(t, err)
	r.GET("/commentoncomment/:articleID/:commentID", handler.CommentOnComment)
	req, err := http.NewRequest("GET", "/commentoncomment/1/1", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	expected := map[string]interface{}{
		"Comments in comment are": response,
	}
	if reflect.DeepEqual(response, expected) {
		t.Errorf("Error on response")
	}
}

func TestGetArticleComments(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	d, _ := db.DB()
	defer d.Close()
	r := gin.Default()
	handler := handler.CommentImpl(db)
	assert.NoError(t, err)
	r.GET("/getArticleComments/:articleID", handler.GetArticleComments)
	req, err := http.NewRequest("GET", "/getArticleComments/1", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	expected := map[string]interface{}{
		"Comments in comment are": response,
	}
	if reflect.DeepEqual(response, expected) {
		t.Errorf("Error on response")
	}
}
