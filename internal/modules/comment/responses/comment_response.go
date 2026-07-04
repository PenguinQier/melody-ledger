package responses

import (
	"fmt"
	CommentModel "github.com/PenguinQier/melody-ledger/internal/modules/comment/models"
	UserResponse "github.com/PenguinQier/melody-ledger/internal/modules/user/responses"
)

type Comment struct {
	ID        uint
	Content   string
	CreatedAt string
	User      UserResponse.User
}

type Comments struct {
	Data []Comment
}

func ToComment(comment CommentModel.Comment) Comment {
	return Comment{
		ID:        comment.ID,
		Content:   comment.Content,
		CreatedAt: fmt.Sprintf("%d/%02d/%02d", comment.CreatedAt.Year(), comment.CreatedAt.Month(), comment.CreatedAt.Day()),
		User:      UserResponse.ToUser(comment.User),
	}
}

func ToComments(comments []CommentModel.Comment) Comments {
	var response Comments
	for _, comment := range comments {
		response.Data = append(response.Data, ToComment(comment))
	}
	return response
}
