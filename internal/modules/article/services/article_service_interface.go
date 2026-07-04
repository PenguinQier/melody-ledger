package services

import (
	"github.com/PenguinQier/melody-ledger/internal/modules/article/requests/articles"
	ArticleResponse "github.com/PenguinQier/melody-ledger/internal/modules/article/responses"
	UserResponse "github.com/PenguinQier/melody-ledger/internal/modules/user/responses"

	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() ArticleResponse.Articles
	GetStoriesArticles() ArticleResponse.Articles
	Find(id int) (ArticleResponse.Article, error)
	StoreAsUser(c *gin.Context, request articles.StoreRequest, user UserResponse.User) (ArticleResponse.Article, error)
	Update(id int, request articles.UpdateRequest) (ArticleResponse.Article, error)
	CanUserEdit(articleID int, userID uint) bool
	Delete(id int) error
	Search(keyword string) ArticleResponse.Articles
	GetUserArticles(userID uint) ArticleResponse.Articles
	UpdateCover(id int, file *multipart.FileHeader) error
}
