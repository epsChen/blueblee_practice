package controller

import (
	"github.com/epsChen/bluebell/logic"
	"github.com/epsChen/bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// 完成用户注册功能

func SignUpController(c *gin.Context) {
	//获取参数
	//我们需要的参数是
	//用户名 密码 并对两次输入的密码进行校验
	//从用户传入的数据中获得数据 并绑定到一个结构体中
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		return
	}
	//调用逻辑层实现业务逻辑
	if err := logic.SiginUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		return
	}

	//返回响应
	c.JSON(http.StatusOK, "sign up success!")
}

// TODO 希望在用户登录完成后生成一个jwt令牌并返回
