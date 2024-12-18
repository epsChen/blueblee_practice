package models

import (
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"time"
)

//user表中应该有什么
/*
ID（自增）
用户ID（根据一种算法自动生成）
用户名（必须唯一）
密码（加密后）
性别
邮箱
创建时间
更新时间
*/

type User struct {
	ID        uint `gorm:"primaryKey"`
	UserID    int64
	UserName  string `gorm:"unique"`
	Password  string
	Gender    int            //1为男 0为女
	Email     sql.NullString //该字段可以为空
	CreatedAt time.Time      //自动生成创建时间
	UpdatedAt time.Time      //自动生成更新时间
}

// 定义创建用户前的钩子函数
// 要求实现的功能有 用户名不能重复的校验
func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	var count int64
	db.Where("user_name =?", u.UserName).Count(&count)
	if count > 0 {
		err = errors.New("用户名已存在")
		return
	}
	return
}
