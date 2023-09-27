package repo

import (
	"blogpost/model"
	"fmt"

	"gorm.io/gorm"
)

var (
	blog    = model.BlogArticle{}
	comment = model.ArticleComments{}
)

// Blog create DB operation
func BlogCreate(dBase *gorm.DB, blogs []model.BlogArticle) error {

	for _, blog := range blogs {
		fmt.Println("this")

		if err := dBase.Debug().Create(&blog).Error; err != nil {
			fmt.Println("this11", err)

			return err
		}
	}
	return nil
}

// Retrive blog article
func GetArticles(dBase *gorm.DB) ([]model.BlogArticle, error) {

	var blogs []model.BlogArticle
	if err := dBase.Find(&blogs).Error; err != nil {
		return blogs, err
	}
	return blogs, nil

}

func GetArticleContent(dBase *gorm.DB, articleID int) (model.BlogArticle, error) {
	if err := dBase.Where("article_id = ?", articleID).Find(&blog).Error; err != nil {
		return blog, err
	}
	return blog, nil
}

func CommentArticle(dBase *gorm.DB, comment model.ArticleComments, articleID int) error {

	_, err := GetArticleContent(dBase, articleID)

	if err != nil {
		return err
	}

	if err := dBase.Debug().Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

func GetArticleComment(dBase *gorm.DB, articleID int) ([]model.ArticleComments, error) {
	var comment []model.ArticleComments
	if err := dBase.Debug().Where("article_id", articleID).Find(&comment).Error; err != nil {
		return comment, err
	}
	return comment, nil
}

func GetCommentOnComment(dBase *gorm.DB, articleID, commentID int) ([]model.ArticleComments, error) {

	var comment []model.ArticleComments

	if err := dBase.Debug().Where("article_id = ? and comment_id = ?", articleID, commentID).Find(&comment).Error; err != nil {
		return comment, err
	}
	return comment, nil
}
