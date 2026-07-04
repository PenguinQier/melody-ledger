package responses

import (
	ArticleModel "github.com/PenguinQier/melody-ledger/internal/modules/article/models"
	UserResponse "github.com/PenguinQier/melody-ledger/internal/modules/user/responses"
)

type Article struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Files     string `json:"files"`
	Image     string `json:"image"`
	CreatedAt string `json:"created_at"`
	UserID    uint   `json:"user_id"`
	User      UserResponse.User
}

type Articles struct {
	Data []Article
}

func ToArticle(article ArticleModel.Article) Article {
	// 如果文章没有设置封面，使用默认图片
	image := article.Image
	if image == "" {
		image = "/assets/img/demopic/default.jpg"
	}

	// 格式化时间为 "2006-01-02 15:04:05" 格式
	formattedTime := article.CreatedAt.Format("2006-01-02 15:04:05")

	return Article{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Files:     article.Files,
		Image:     image,
		CreatedAt: formattedTime,
		UserID:    article.UserID,
		User:      UserResponse.ToUser(article.User),
	}
}

func ToArticles(article []ArticleModel.Article) Articles {
	var response Articles

	for _, article := range article {
		response.Data = append(response.Data, ToArticle(article))
	}

	return response
}
