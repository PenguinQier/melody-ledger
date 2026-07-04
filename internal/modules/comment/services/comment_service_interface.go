package services

import (
	"github.com/PenguinQier/melody-ledger/internal/modules/comment/requests/comments"
	CommentResponse "github.com/PenguinQier/melody-ledger/internal/modules/comment/responses"
	UserResponse "github.com/PenguinQier/melody-ledger/internal/modules/user/responses"
)

type CommentServiceInterface interface {
	Create(request comments.StoreRequest, user UserResponse.User) CommentResponse.Comment
	GetArticleComments(articleID uint) CommentResponse.Comments
	Delete(id uint, userID uint) error
	CanUserDelete(commentID uint, userID uint) bool
}
