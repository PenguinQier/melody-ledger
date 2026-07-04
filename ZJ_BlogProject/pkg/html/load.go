package html

import (
	"html/template"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func LoadHTML(router *gin.Engine) {
	// 注册自定义模板函数
	router.SetFuncMap(template.FuncMap{
		"split": strings.Split,
		"add": func(a, b int) int {
			return a + b
		},
		"isImage": func(path string) bool {
			ext := strings.ToLower(filepath.Ext(path))
			return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif"
		},
	})

	// router.LoadHTMLGlob("internal/**/*tmpl")
	router.LoadHTMLGlob("internal/**/**/**/*tmpl")
}
