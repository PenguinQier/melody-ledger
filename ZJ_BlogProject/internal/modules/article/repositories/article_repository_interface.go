package repositories

import (
	ArticleModel "github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/article/models"
)

type ArticleRepositoryInterface interface {
	List(limit int) []ArticleModel.Article
	Find(id int) ArticleModel.Article
	Create(article ArticleModel.Article) ArticleModel.Article
	Update(article ArticleModel.Article) ArticleModel.Article
	Delete(article ArticleModel.Article) error
	Search(keyword string) []ArticleModel.Article
	GetUserArticles(userID uint) []ArticleModel.Article
}
