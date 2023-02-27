package initialize

import "go.uber.org/zap"

func InitLogger() {
	//注册了,zap全局可使用
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
