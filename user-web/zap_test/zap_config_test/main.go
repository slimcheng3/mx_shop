package main

import (
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error){
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"./mxshop.log", "stdout"}
	return cfg.Build()
}

func main()  {
	// logger, _ := zap.NewProduction()
	logger, err := NewLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch url", "url", 22)
	sugar.Infof("Failed to fetch url: %s", 123)
}