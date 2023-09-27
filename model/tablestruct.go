package model

import (
	"time"
)

type BlogArticle struct {
	ArticleID      uint      `gorm:"primaryKey" json:"articleid"`
	ArticleTitle   string    `json:"articletitle"`
	ArticleContent string    `gorm:"type:longtext" json:"content"`
	Nickname       string    `json:"nickname"`
	CreationDate   time.Time `json:"date"`
}

func (BlogArticle) TableName() string {
	return "blog_articles"
}

type ArticleComments struct {
	ArticleID       int       `json:"articleid"`
	CommentID       int       `gorm:"primaryKey" json:"commentid"`
	ParentCommentID int       `json:"parentcommentid"`
	Content         string    `gorm:"type:longtext" json:"content"`
	Nickname        string    `json:"nickname"`
	CreationDate    time.Time `json:"date"`
}

func (ArticleComments) TableName() string {
	return "article_comments"
}
