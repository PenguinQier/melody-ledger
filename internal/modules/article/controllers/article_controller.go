package controllers

import (
	"github.com/PenguinQier/melody-ledger/internal/modules/article/requests/articles"
	ArticleService "github.com/PenguinQier/melody-ledger/internal/modules/article/services"
	"github.com/PenguinQier/melody-ledger/internal/modules/comment/services"
	"github.com/PenguinQier/melody-ledger/internal/modules/user/helpers"
	"github.com/PenguinQier/melody-ledger/pkg/html"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	articleService ArticleService.ArticleServiceInterface
	commentService services.CommentServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
		commentService: services.New(),
	}
}

// Create 显示创建文章页面
func (controller *Controller) Create(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/article/html/create", gin.H{
		"title": "创建文章",
	})
}

// Store 保存新文章
func (controller *Controller) Store(c *gin.Context) {
	var request articles.StoreRequest

	// 解析多部分表单
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil { // 32MB max
		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	// 获取上传的文件
	form, _ := c.MultipartForm()
	request.Files = form.File["files[]"]
	request.Title = c.PostForm("title")
	request.Content = c.PostForm("content")

	authUser := helpers.Auth(c)
	_, err := controller.articleService.StoreAsUser(c, request, authUser)
	if err != nil {
		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	c.Redirect(http.StatusFound, "/")
}

// Show 显示文章详情
func (controller *Controller) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := controller.articleService.Find(id)
	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	comments := controller.commentService.GetArticleComments(uint(id))

	html.Render(c, http.StatusOK, "modules/article/html/show", gin.H{
		"title":    article.Title,
		"article":  article,
		"comments": comments,
	})
}

// Edit 显示编辑文章页面
func (controller *Controller) Edit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	authUser := helpers.Auth(c)
	if !controller.articleService.CanUserEdit(id, authUser.ID) {
		c.Redirect(http.StatusFound, "/articles/"+c.Param("id"))
		return
	}

	article, err := controller.articleService.Find(id)
	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	html.Render(c, http.StatusOK, "modules/article/html/edit", gin.H{
		"title":   "编辑文章",
		"article": article,
	})
}

// HandleUpdate 处理更新文章请求
func (controller *Controller) HandleUpdate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	authUser := helpers.Auth(c)
	if !controller.articleService.CanUserEdit(id, authUser.ID) {
		c.Redirect(http.StatusFound, "/articles/"+c.Param("id"))
		return
	}

	var request articles.UpdateRequest
	request.Title = c.PostForm("title")
	request.Content = c.PostForm("content")

	article, err := controller.articleService.Update(id, request)
	if err != nil {
		c.Redirect(http.StatusFound, "/articles/"+c.Param("id")+"/edit")
		return
	}

	c.Redirect(http.StatusFound, "/articles/"+strconv.Itoa(int(article.ID)))
}

// HandleDelete 处理删除文章请求
func (controller *Controller) HandleDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	authUser := helpers.Auth(c)
	if !controller.articleService.CanUserEdit(id, authUser.ID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "没有权限删除此文章"})
		return
	}

	err := controller.articleService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "文章删除成功"})
}

// Search 搜索文章
func (controller *Controller) Search(c *gin.Context) {
	keyword := c.Query("keyword")
	articles := controller.articleService.Search(keyword)

	html.Render(c, http.StatusOK, "modules/article/html/search", gin.H{
		"title":    "搜索结果",
		"articles": articles,
		"keyword":  keyword,
	})
}

// UpdateCover 更新文章封面
func (controller *Controller) UpdateCover(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	authUser := helpers.Auth(c)
	if !controller.articleService.CanUserEdit(id, authUser.ID) {
		c.Redirect(http.StatusFound, "/articles/"+strconv.Itoa(id))
		return
	}

	file, err := c.FormFile("cover")
	if err != nil {
		c.Redirect(http.StatusFound, "/articles/"+strconv.Itoa(id))
		return
	}

	err = controller.articleService.UpdateCover(id, file)
	if err != nil {
		c.Redirect(http.StatusFound, "/articles/"+strconv.Itoa(id))
		return
	}

	c.Redirect(http.StatusFound, "/articles/"+strconv.Itoa(id))
}

// Home 显示首页
func (controller *Controller) Home(c *gin.Context) {
	featured := controller.articleService.GetFeaturedArticles()
	stories := controller.articleService.GetStoriesArticles()

	html.Render(c, http.StatusOK, "modules/article/html/home", gin.H{
		"title":    "首页",
		"featured": featured,
		"stories":  stories,
	})
}

// Index 显示文章列表
func (controller *Controller) Index(c *gin.Context) {
	articles := controller.articleService.GetFeaturedArticles()

	html.Render(c, http.StatusOK, "modules/article/html/index", gin.H{
		"title":    "文章列表",
		"articles": articles,
	})
}

// APISearch 处理 API 搜索请求
func (controller *Controller) APISearch(c *gin.Context) {
	keyword := c.Query("keyword")
	articles := controller.articleService.Search(keyword)
	c.JSON(http.StatusOK, articles)
}
