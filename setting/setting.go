package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(Config)

type Config struct {
	//内存对齐可以节约内存 将相同类型的参数定义在一起
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineId int64  `mapstructure:"machine_id"`
	Port      int    `mapstructure:"port"`

	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*LogConfig   `mapstructure:"log"`
}

type MySQLConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"dbname"`
	Port     int    `mapstructure:"port"`
	//MaxOpenConns int    `mapstructure:max_open_conns`
	//MaxIdleConns int    `mapstructure:max_idle_conns`//TODO gorm初始化未完成
}

type RedisConfig struct {
	Port     int `mapstructure:"port"`
	DB       int `mapstructure:"db"`
	PoolSize int `mapstructure:"pool_size"`
}

type LogConfig struct {
	Filename   string `mapstructure:"filename"`
	Level      string `mapstructure:"level"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_bakups"` //备份
}

func InitSettings() (err error) {
	//读取配置文件config.yaml
	viper.SetConfigFile("./config/config.yaml")
	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			//配置文件未找到
			fmt.Println("cannot find the config file")
		} else {
			//配置文件找到了但是有其他错误
			fmt.Println("read in config failed")
		}
		return
	}

	//将配置文件中的信息初始化到对应的结构体中
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.unmarshal failed, err:%v\n", err)
		return
	}
	//监控config有没有变化 如果有重新读取config
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("配置文件被修改了-----")
		//重新读取配置文件
	})
	return
}
