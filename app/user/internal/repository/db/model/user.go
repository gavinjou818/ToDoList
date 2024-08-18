package model

import (
	"golang.org/x/crypto/bcrypt"
	"grpc_todolist/consts"
)

type User struct {
	UserId         int64  `gorm:"primarykey"`
	UserName       string `gorm:"unique"`
	NickName       string
	PasswordDigest string
}

func (*User) TableName() string {
	return "user"
}

func (this *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), consts.PassWordCost)
	if err != nil {
		return err
	}
	this.PasswordDigest = string(bytes)
	return nil
}

func (this *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(this.PasswordDigest), []byte(password))
	return err == nil
}
