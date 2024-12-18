package main

import (
	"fmt"
	"github.com/epsChen/bluebell/dao/mysql"
	"github.com/epsChen/bluebell/logger"
	"github.com/epsChen/bluebell/route"
	"github.com/epsChen/bluebell/setting"
	"github.com/epsChen/bluebell/utils"
)

func main() {
	//加载配置文件
	if err := setting.InitSettings(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}
	//初始化数据库（记得关闭数据库连接）
	//mysql
	if err := mysql.InitMySQL(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	//redis TODO后续配置 用于点赞功能

	//初始化日志 TODO 不理解的内容较多 主要是zap日志库中的函数
	if err := logger.InitLogger(setting.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	err := utils.InitSnowfalke(setting.Conf.StartTime, setting.Conf.MachineId)
	if err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	//注册路由
	r, err := route.InitRoute()
	if err != nil {
		fmt.Printf("init route falied, err:%v\n", err)
		return
	}
	r.Run(":8080")

}
