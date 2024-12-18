package route

import (
	"fmt"
	"github.com/epsChen/bluebell/controller"
	"github.com/gin-gonic/gin"
)

func InitRoute() (r *gin.Engine, err error) {
	r = gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		fmt.Println("hello world!")
	})

	//完成注册用户的路由
	r.POST("/signup", controller.SignUpController)
	return
}
