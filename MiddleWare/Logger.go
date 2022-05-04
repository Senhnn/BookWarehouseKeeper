package MiddleWare

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"time"
)

var logger *log.Logger

// Logger 中间件日志
var Logger *log.Entry

// GLog 全局日志
var GLog *log.Logger

func InitGrpcLogger() {
	// Grpc日志
	logger = log.New()
	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	logText, err := os.OpenFile(viper.GetStringMapString("server")["grpclogpath"],
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(any("log init failed"))
	}
	logger.Out = logText
	logger.SetLevel(log.DebugLevel)
	Logger = &log.Entry{
		Logger:  logger,
		Data:    nil,
		Time:    time.Time{},
		Level:   log.DebugLevel,
		Caller:  nil,
		Message: "grpc-logger",
		Buffer:  nil,
		Context: nil,
	}
}

func InitGLog() {
	// 全局日志
	GLog = log.New()
	GLog.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	glogText, err := os.OpenFile(viper.GetStringMapString("server")["globallogpath"],
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(any("log init failed"))
	}
	GLog.Out = glogText
	GLog.Out = os.Stdout
	GLog.SetLevel(log.DebugLevel)
}

func LoggerInit() {
	InitGrpcLogger()
	InitGLog()
}
