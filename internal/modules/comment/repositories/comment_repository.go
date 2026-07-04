package repositories

import (
	CommentModel "github.com/PenguinQier/melody-ledger/internal/modules/comment/models"
	"github.com/PenguinQier/melody-ledger/pkg/database"

	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func New() *CommentRepository {
	return &CommentRepository{
		DB: database.Connection(),
	}
}

func (commentRepository *CommentRepository) Create(comment CommentModel.Comment) CommentModel.Comment {
	commentRepository.DB.Create(&comment)
	return comment
}

func (commentRepository *CommentRepository) GetArticleComments(articleID uint) []CommentModel.Comment {
	var comments []CommentModel.Comment
	commentRepository.DB.Where("article_id = ?", articleID).
		Joins("User").
		Order("created_at DESC").
		Find(&comments)
	return comments
}

func (commentRepository *CommentRepository) GetCommentWithUser(id uint) CommentModel.Comment {
	var comment CommentModel.Comment
	commentRepository.DB.Joins("User").First(&comment, id)
	return comment
}

func (commentRepository *CommentRepository) FindByID(id uint) CommentModel.Comment {
	var comment CommentModel.Comment
	commentRepository.DB.First(&comment, id)
	return comment
}

func (commentRepository *CommentRepository) Delete(comment CommentModel.Comment) error {
	result := commentRepository.DB.Delete(&comment)
	return result.Error
}
