package migration

import (
	"fmt"
	articleModels "github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/article/models"
	userModels "github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/user/models"
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/pkg/database"
	"log"
)

func Migrate() {
	db := database.Connection()

	err := db.AutoMigrate(&userModels.User{}, &articleModels.Article{})

	if err != nil {
		log.Fatal("数据库迁移失败")
		return
	}

	fmt.Println("数据库迁移完成...")
}
