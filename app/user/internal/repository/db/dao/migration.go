package dao

import (
	"grpc_todolist/app/user/internal/repository/db/model"
	"grpc_todolist/pkg/util/logger"
	"os"
)

func migration() {
	// 使用 AutoMigrate 可以方便地进行数据库表的初始化和更新，而无需手动执行 SQL 语句
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{})
	if err != nil {
		logger.LogrusObj.Infoln("register table fail")
		os.Exit(0)
	}
	logger.LogrusObj.Infoln("register table success")
}
