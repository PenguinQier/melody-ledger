package services

import (
	"errors"
	"fmt"
	ArticleModel "github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/article/models"
	ArticleRepository "github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/article/repositories"
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/article/requests/articles"
	ArticleResponse "github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/article/responses"
	UserResponse "github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/user/responses"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ArticleService struct {
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: ArticleRepository.New(),
	}
}

func (articleService *ArticleService) GetFeaturedArticles() ArticleResponse.Articles {
	// 获取精选文章列表
	articles := articleService.articleRepository.List(4)
	return ArticleResponse.ToArticles(articles)
}

func (articleService *ArticleService) GetStoriesArticles() ArticleResponse.Articles {
	// 获取故事文章列表
	articles := articleService.articleRepository.List(6)
	return ArticleResponse.ToArticles(articles)
}

func (articleService *ArticleService) Find(id int) (ArticleResponse.Article, error) {
	var response ArticleResponse.Article
	// 根据ID查找文章
	article := articleService.articleRepository.Find(id)

	if article.ID == 0 {
		return response, errors.New("未找到文章")
	}

	return ArticleResponse.ToArticle(article), nil
}

func (articleService *ArticleService) handleFileUploads(c *gin.Context, files []*multipart.FileHeader) ([]string, error) {
	var filePaths []string
	uploadDir := "./public/uploads/"

	// 确保上传目录存在
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return nil, err
	}

	for _, file := range files {
		// 生成唯一文件名
		ext := filepath.Ext(file.Filename)
		newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
		filePath := filepath.Join(uploadDir, newFileName)

		// 保存文件
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			return nil, err
		}

		// 存储相对路径
		filePaths = append(filePaths, "/public/uploads/"+newFileName)
	}

	return filePaths, nil
}

func (articleService *ArticleService) StoreAsUser(c *gin.Context, request articles.StoreRequest, user UserResponse.User) (ArticleResponse.Article, error) {
	var article ArticleModel.Article
	var response ArticleResponse.Article

	// 处理文件上传
	if len(request.Files) > 0 {
		filePaths, err := articleService.handleFileUploads(c, request.Files)
		if err != nil {
			return response, err
		}
		article.Files = strings.Join(filePaths, ",")

		// 如果上传的第一个文件是图片,设置为封面
		if strings.HasPrefix(request.Files[0].Header.Get("Content-Type"), "image/") {
			article.Image = filePaths[0]
		}
	}

	// 设置文章基本信息
	article.Title = request.Title
	article.Content = request.Content
	article.UserID = user.ID

	// 创建新文章
	newArticle := articleService.articleRepository.Create(article)
	if newArticle.ID == 0 {
		return response, errors.New("创建文章失败")
	}

	return ArticleResponse.ToArticle(newArticle), nil
}

func (articleService *ArticleService) Update(id int, request articles.UpdateRequest) (ArticleResponse.Article, error) {
	var response ArticleResponse.Article

	// 查找要更新的文章
	article := articleService.articleRepository.Find(id)
	if article.ID == 0 {
		return response, errors.New("未找到文章")
	}

	// 更新文章内容
	article.Title = request.Title
	article.Content = request.Content

	updatedArticle := articleService.articleRepository.Update(article)
	return ArticleResponse.ToArticle(updatedArticle), nil
}

func (articleService *ArticleService) CanUserEdit(articleID int, userID uint) bool {
	// 检查用户是否有权限编辑文章
	article := articleService.articleRepository.Find(articleID)
	return article.UserID == userID
}

func (articleService *ArticleService) Delete(id int) error {
	// 查找要删除的文章
	article := articleService.articleRepository.Find(id)
	if article.ID == 0 {
		return errors.New("未找到文章")
	}

	return articleService.articleRepository.Delete(article)
}

func (articleService *ArticleService) Search(keyword string) ArticleResponse.Articles {
	if strings.TrimSpace(keyword) == "" {
		return ArticleResponse.Articles{}
	}

	articles := articleService.articleRepository.Search(keyword)
	return ArticleResponse.ToArticles(articles)
}

// GetUserArticles 获取用户的文章列表
func (articleService *ArticleService) GetUserArticles(userID uint) ArticleResponse.Articles {
	articles := articleService.articleRepository.GetUserArticles(userID)
	return ArticleResponse.ToArticles(articles)
}

// 添加更新封面图方法
func (articleService *ArticleService) UpdateCover(id int, file *multipart.FileHeader) error {
	// 获取文章
	article := articleService.articleRepository.Find(id)
	if article.ID == 0 {
		return errors.New("文章不存在")
	}

	// 处理文件上传
	uploadDir := "./public/uploads/"
	ext := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("cover_%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(uploadDir, newFileName)

	// 保存文件
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	// 更新文章封面
	article.Image = "/public/uploads/" + newFileName
	articleService.articleRepository.Update(article)

	return nil
}
