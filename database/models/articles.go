package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	db      *gorm.DB `gorm:"-"  json:"-"`
	UUID    string   `gorm:"column:uuid" json:"uuid"`
	Title   string   `gorm:"column:title" json:"title"`
	Content string   `gorm:"column:content" json:"content"`
	Author  string   `gorm:"column:author" json:"author"`
}

func NewArticle(db *gorm.DB) *Article {
	return &Article{
		db: db,
	}
}

func (a *Article) TableName() string {
	return "articles"
}

// InsertArticle inserts an entry to articles table
func (a *Article) InsertArticle(data Article) (err error) {
	if err = a.db.Save(&data).Error; err != nil {
		return err
	}

	return nil
}

// GetArticle retrieve record from articles table based on the uuid provided
func (a *Article) GetArticle(id string) (article Article, err error) {
	if err = a.db.Table(a.TableName()).Select("*").Where("uuid = ?", id).First(&article).Error; err != nil {
		return
	}

	return article, nil
}

// GetArticles retrieves all articles from articles table
func (a *Article) GetArticles() (articles []Article, err error) {
	if err = a.db.Table(a.TableName()).Select("*").Find(&articles).Error; err != nil {
		return
	}

	return articles, nil
}
