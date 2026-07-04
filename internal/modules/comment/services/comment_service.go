package services

import (
	"errors"
	CommentModel "github.com/PenguinQier/melody-ledger/internal/modules/comment/models"
	CommentRepository "github.com/PenguinQier/melody-ledger/internal/modules/comment/repositories"
	"github.com/PenguinQier/melody-ledger/internal/modules/comment/requests/comments"
	CommentResponse "github.com/PenguinQier/melody-ledger/internal/modules/comment/responses"
	UserResponse "github.com/PenguinQier/melody-ledger/internal/modules/user/responses"
)

type CommentService struct {
	commentRepository CommentRepository.CommentRepositoryInterface
}

func New() *CommentService {
	return &CommentService{
		commentRepository: CommentRepository.New(),
	}
}

func (commentService *CommentService) Create(request comments.StoreRequest, user UserResponse.User) CommentResponse.Comment {
	// 创建新评论
	comment := CommentModel.Comment{
		Content:   request.Content,
		ArticleID: request.ArticleID,
		UserID:    user.ID,
	}

	// 保存评论到数据库
	newComment := commentService.commentRepository.Create(comment)

	// 获取包含用户信息的完整评论
	commentWithUser := commentService.commentRepository.GetCommentWithUser(newComment.ID)

	return CommentResponse.ToComment(commentWithUser)
}

func (commentService *CommentService) GetArticleComments(articleID uint) CommentResponse.Comments {
	// 获取文章的所有评论
	comments := commentService.commentRepository.GetArticleComments(articleID)
	return CommentResponse.ToComments(comments)
}

func (commentService *CommentService) Delete(id uint, userID uint) error {
	// 查找评论
	comment := commentService.commentRepository.FindByID(id)
	if comment.ID == 0 {
		return errors.New("未找到评论")
	}

	// 检查权限
	if comment.UserID != userID {
		return errors.New("您没有权限删除此评论")
	}

	// 删除评论
	return commentService.commentRepository.Delete(comment)
}

func (commentService *CommentService) CanUserDelete(commentID uint, userID uint) bool {
	// 检查用户是否有权限删除评论
	comment := commentService.commentRepository.FindByID(commentID)
	return comment.UserID == userID
}
