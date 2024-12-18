package controller

import (
	"errors"
	"github.com/epsChen/bluebell/logic"
	"github.com/epsChen/bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 完成用户注册功能
func SignUpContoller(c *gin.Context) {
	//获取参数
	//我们需要的参数是
	//用户名 密码 并对两次输入的密码进行校验
	//从用户传入的数据中获得数据 并绑定到一个结构体中
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//调用逻辑层实现业务逻辑
	if err := logic.SiginUp(p); err != nil {
		//用户名重复时会返回错误 但不是钩子函数中的ErrUserExist 而是 Duplicate entry 'user1' for key 'users.uni_users_user_name
		//已解决 不使用钩子函数 请跳转mysql.user中的insert
		if errors.Is(err, models.ErrUserExist) {
			zap.L().Error("user exist", zap.Error(err))
			ResponseError(c, CodeUserExist)
			return
		}
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	//返回响应
	ResponseSuccess(c, CodeSuccess)
}

// TODO 希望在用户登录完成后生成一个jwt令牌并返回
