package logic

import (
	"fmt"
	"github.com/epsChen/bluebell/dao/mysql"
	"github.com/epsChen/bluebell/models"
	"go.uber.org/zap"
)

func SiginUp(p *models.ParamSignUp) (err error) {
	//将被绑定的结构体对应的数据 传入数据库中
	//自动建表 防止表不存在
	err = mysql.CheckTable(&models.User{})
	if err != nil {
		fmt.Printf("table init failed, err:%v\n", err)
		return
	}

	//1.检查用户名是否存在
	//if err = mysql.CheckUser(p.Username); err != nil {
	//	zap.L().Error("user exist", zap.Error(err))
	//	return
	//}

	//尝试使用gorm的钩子函数进行解决参数校验的问题

	//传入数据库
	if err = mysql.InsertUser(p); err != nil {
		zap.L().Error("mysql.InsertUser failed", zap.Error(err))
		return
	}

	return
}
