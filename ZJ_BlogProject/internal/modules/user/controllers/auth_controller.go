package controllers

import (
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/user/helpers"
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/user/requests/auth"
	Userservices "github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/user/services"
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/pkg/converters"
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/pkg/errors"
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/pkg/html"
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/pkg/old"
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/pkg/sessions"
	"log"
	"net/http"
	"strconv"

	ArticleService "github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/article/services"

	"github.com/gin-gonic/gin"
)

// Controller 用户认证控制器
type Controller struct {
	userService    Userservices.UserServiceInterface
	articleService ArticleService.ArticleServiceInterface
}

// New 创建新的认证控制器实例
func New() *Controller {
	return &Controller{
		userService:    Userservices.New(),
		articleService: ArticleService.New(),
	}
}

func (controller *Controller) Register(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/register", gin.H{
		"title": "注册",
	})
}

func (controller *Controller) HandleRegister(c *gin.Context) {
	// 验证注册请求数据
	var request auth.RegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFormErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/register")
		return
	}

	if controller.userService.CheckUserExists(request.Email) {
		errors.Init()
		errors.Add("Email", "该邮箱已被注册")
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/register")
		return
	}

	// 创建用户
	user, err := controller.userService.Create(request)

	// 检查用户创建过程中是否有错误
	if err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	// 创建用户成功后，自动登录
	loginRequest := auth.LoginRequest{
		Email:    request.Email,
		Password: request.Password,
	}

	loggedInUser, err := controller.userService.HandleUserLogin(loginRequest)
	if err != nil {
		log.Printf("Auto login failed for new user: %v", err)
	} else {
		// 设置登录 session
		sessions.Set(c, "auth", strconv.Itoa(int(loggedInUser.ID)))
	}

	// 创建用户成功后，记录日志并重定向到首页
	log.Printf("用户 %s 创建成功并自动登录\n", user.Name)
	c.Redirect(http.StatusFound, "/")
}

func (controller *Controller) Login(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/login", gin.H{
		"title": "登录",
	})
}

func (controller *Controller) HandleLogin(c *gin.Context) {
	// 验证登录请求数据
	var request auth.LoginRequest
	// 根据 content-type 推断使用哪个绑定器
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFormErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/login")
		return
	}

	user, err := controller.userService.HandleUserLogin(request)
	if err != nil {
		errors.Init()
		errors.Add("email", err.Error())
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/login")
		return
	}

	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

	// 登录成功后，记录日志并重定向到首页
	log.Printf("用户 %s 登录成功\n", user.Name)
	c.Redirect(http.StatusFound, "/")
}

// HandleLogout 处理用户登出请求
func (controller *Controller) HandleLogout(c *gin.Context) {
	sessions.Remove(c, "auth")
	c.Redirect(http.StatusFound, "/")
}

func (controller *Controller) Profile(c *gin.Context) {
	// 获取当前用户
	user := helpers.Auth(c)

	// 获取用户发表的文章列表
	articles := controller.articleService.GetUserArticles(user.ID)

	html.Render(c, http.StatusOK, "modules/user/html/profile", gin.H{
		"title":    "个人中心",
		"AUTH":     user,
		"articles": articles,
	})
}

func (controller *Controller) HandleUpdateProfile(c *gin.Context) {
	// 验证个人资料更新请求
	var request auth.ProfileRequest
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFormErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/profile")
		return
	}

	// 获取当前用户并更新资料
	authUser := helpers.Auth(c)
	user, err := controller.userService.UpdateProfile(authUser.ID, request)

	if err != nil {
		errors.Init()
		errors.Add("email", err.Error())
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/profile")
		return
	}

	// 用户 %s 的个人资料更新成功
	log.Printf("用户 %s 的个人资料更新成功\n", user.Name)
	c.Redirect(http.StatusFound, "/profile")
}
