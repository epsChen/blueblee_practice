package logger

import (
	"github.com/epsChen/bluebell/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

//	func InitLogger() {
//		writeSyncer := getLogWriter()
//		encoder := getEncoder()
//		core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
//
//		logger := zap.New(core)
//		sugarLogger = logger.Sugar()
//	}
//
//	func getEncoder() zapcore.Encoder {
//		return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
//	}
//
//	func getLogWriter() zapcore.WriteSyncer {
//		file, _ := os.Create("./test.log")
//		return zapcore.AddSync(file)
//	}
var lg *zap.Logger

func InitLogger(cfg *setting.LogConfig) (err error) {

	encoder := getEncoder()
	writeSyncer := getLogWrtier(cfg.Filename, cfg.MaxSize, cfg.MaxSize, cfg.MaxBackups)

	//读取日志的level
	var l = new(zapcore.Level)
	if err = l.UnmarshalText([]byte(cfg.Level)); err != nil {
		return
	}
	var core zapcore.Core
	//如果是开发者模式 我们要将日志输出到控制台并且添加到日志文件
	if cfg.Level == "dev" {
		//console控制台
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		//该函数用于把多个core合在一起
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, l),
			//TODO lock是一个互斥锁 不理解
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		//只添加到日志文件不添加到控制台
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}

	//AddCaller 添加将调用函数信息记录到日志中的功能
	lg = zap.New(core, zap.AddCaller())
	//替换为全局变量
	zap.ReplaceGlobals(lg)
	zap.L().Info("init logger success")
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder         //修改时间编码器
	encoderConfig.TimeKey = "time"                                //TODO 不理解具体含义
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder       //在日志文件中使用大写字母记录日志级别
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder //TODO 不理解具体含义
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder       //TODO 不理解具体含义
	return zapcore.NewConsoleEncoder(encoderConfig)
}

//func getWriteSyncer() zapcore.WriteSyncer {
//	//file, _ := os.Create("./bluebell.log")
//	//writeSycner := zapcore.AddSync(file)
//	//return writeSycner
//	file, _ := os.Create("./bluebell.log")
//	//将日志同时输出到文件和终端
//	ws := io.MultiWriter(file, os.Stdout)
//	return zapcore.AddSync(ws)
//}

func getLogWrtier(filename string, maxSize, maxAge, maxBackups int) zapcore.WriteSyncer {
	lumberjack := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
	}

	return zapcore.AddSync(lumberjack)
}
