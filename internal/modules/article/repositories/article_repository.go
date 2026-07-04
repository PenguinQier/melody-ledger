package repositories

import (
	ArticleModel "github.com/PenguinQier/melody-ledger/internal/modules/article/models"
	"github.com/PenguinQier/melody-ledger/pkg/database"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func New() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connection(),
	}
}

func (articleRepository *ArticleRepository) List(limit int) []ArticleModel.Article {
	var articles []ArticleModel.Article

	articleRepository.DB.Limit(limit).Joins("User").Order("rand()").Find(&articles)

	return articles
}

func (articleRepository *ArticleRepository) Find(id int) ArticleModel.Article {
	var article ArticleModel.Article

	articleRepository.DB.Joins("User").First(&article, id)

	return article
}

func (articleRepository *ArticleRepository) Create(article ArticleModel.Article) ArticleModel.Article {
	var newArticle ArticleModel.Article

	articleRepository.DB.Create(&article).Scan(&newArticle)

	return newArticle
}

func (articleRepository *ArticleRepository) Update(article ArticleModel.Article) ArticleModel.Article {
	articleRepository.DB.Save(&article)
	return article
}

func (articleRepository *ArticleRepository) Delete(article ArticleModel.Article) error {
	result := articleRepository.DB.Delete(&article)
	return result.Error
}

func (articleRepository *ArticleRepository) Search(keyword string) []ArticleModel.Article {
	var articles []ArticleModel.Article

	articleRepository.DB.Joins("User").
		Where("LOWER(title) LIKE LOWER(?) OR LOWER(content) LIKE LOWER(?)",
			"%"+keyword+"%", "%"+keyword+"%").
		Order("created_at DESC").
		Limit(20).
		Find(&articles)

	return articles
}

// GetUserArticles 获取指定用户的文章列表
func (repo *ArticleRepository) GetUserArticles(userID uint) []ArticleModel.Article {
	var articles []ArticleModel.Article
	repo.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&articles)
	return articles
}
