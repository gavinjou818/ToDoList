package dao

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"grpc_todolist/app/user/internal/repository/db/model"
	userPb "grpc_todolist/idl/pb/user"
	"grpc_todolist/pkg/util/logger"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func (dao *UserDao) GetUserInfo(req *userPb.UserRequest) (r *model.User, err error) {
	//Model()：用来指定查询的表名，Model() 的入参是我们定义的模型结构体，按照约定用结构体名的蛇形命名复数形式或者其实现的TableName方法返回值作为sql的表名
	err = dao.Model(&model.User{}).Where("user_name=?", req.UserName).First(&r).Error
	return
}

func (dao *UserDao) CreateUser(req *userPb.UserRequest) (err error) {
	var user model.User
	var count int64
	dao.Model(&model.User{}).Where("user_name = ?", req.UserName).Count(&count)
	if count != 0 {
		return errors.New("UserName Exist")
	}

	user = model.User{
		UserName: req.UserName,
		NickName: req.NickName,
	}
	_ = user.SetPassword(req.Password)
	if err = dao.Model(&model.User{}).Create(&user).Error; err != nil {
		logger.LogrusObj.Error("Insert User Error:" + err.Error())
		return
	}
	return
}
