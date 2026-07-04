package seeder

import (
	"fmt"
	articleModel "github.com/PenguinQier/melody-ledger/internal/modules/article/models"
	userModel "github.com/PenguinQier/melody-ledger/internal/modules/user/models"
	"github.com/PenguinQier/melody-ledger/pkg/database"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	db := database.Connection()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("secret"), 12)
	if err != nil {
		log.Fatal("密码加密失败")
		return
	}
	user := userModel.User{Name: "Random name", Email: "random@email.com", Password: string(hashedPassword)}
	db.Create(&user) // 创建用户记录

	log.Printf("成功创建用户，邮箱地址: %s \n", user.Email)

	for i := 1; i <= 10; i++ {
		article := articleModel.Article{Title: fmt.Sprintf("Random title %d", i), Content: fmt.Sprintf("Random content %d", i), UserID: 1}
		db.Create(&article) // 创建文章记录

		log.Printf("成功创建文章，标题: %s \n", article.Title)
	}

	log.Println("数据填充完成...")
}
