package controllers

import (
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/comment/requests/comments"
	CommentService "github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/comment/services"
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/user/helpers"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	commentService CommentService.CommentServiceInterface
}

func New() *Controller {
	return &Controller{
		commentService: CommentService.New(),
	}
}

func (controller *Controller) Store(c *gin.Context) {
	var request comments.StoreRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求数据无效",
		})
		return
	}

	log.Printf("Parsed request: %+v\n", request)

	user := helpers.Auth(c)
	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
		})
		return
	}

	comment := controller.commentService.Create(request, user)
	log.Printf("Created comment: %+v\n", comment)

	c.JSON(http.StatusOK, gin.H{
		"id":         comment.ID,
		"content":    comment.Content,
		"created_at": comment.CreatedAt,
		"user": gin.H{
			"id":    comment.User.ID,
			"name":  comment.User.Name,
			"image": user.Image,
		},
	})
}

func (controller *Controller) HandleDelete(c *gin.Context) {
	commentID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "评论ID无效"})
		return
	}

	user := helpers.Auth(c)
	if !controller.commentService.CanUserDelete(uint(commentID), user.ID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "您没有权限删除此评论"})
		return
	}

	err = controller.commentService.Delete(uint(commentID), user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "评论删除成功"})
}
