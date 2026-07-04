package repositories

import (
	CommentModel "github.com/PenguinQier/melody-ledger/internal/modules/comment/models"
)

type CommentRepositoryInterface interface {
	Create(comment CommentModel.Comment) CommentModel.Comment
	GetArticleComments(articleID uint) []CommentModel.Comment
	GetCommentWithUser(id uint) CommentModel.Comment
	FindByID(id uint) CommentModel.Comment
	Delete(comment CommentModel.Comment) error
}
