package flag

import (
	"blog_server/global"
	"blog_server/models"
)

func MakeMigration() {
	var err error
	//global.DB.SetupJoinTable(&models.UserModel{}, "collectsModels", &models.UserCollectsModel{})
	//global.DB.SetupJoinTable(&models.MenuModel{}, "Image", &models.MenuImageModel{})

	global.Log.Infof("开始迁移数据库")
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		//&models.ImageModel{},
		//&models.TagModel{},
		//&models.MessageModel{},
		//&models.AdvertModel{},
		&models.UserModel{},
		//&models.CommentModel{},
		//&models.ArticleModel{},
		//&models.MenuModel{},
		//&models.MenuImageModel{},
		//&models.FeedbackModel{},
		//&models.LoginDataModel{},
	)
	if err != nil {
		global.Log.Error("数据库迁移失败: %v", err)
		return
	}
	global.Log.Info("数据库迁移成功")

}
