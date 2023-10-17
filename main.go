package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"blog/config"
	"blog/dao/mysql"
	"blog/dao/redis"
	"blog/pkg/logger"
	"blog/pkg/snowflake"
	"blog/router"

	"go.uber.org/zap"
)

func main() {
	if err := config.Init(os.Args...); err != nil {
		fmt.Printf("config init err: %v", err)
		return
	}

	if err := logger.Init(config.Config.Log, config.Config.Mode); err != nil {
		fmt.Printf("logger init err: %v", err)
		return
	}

	if err := mysql.Init(config.Config.Mysql); err != nil {
		logger.Error("mysql init err", zap.Error(err))
		return
	}
	defer mysql.Close()

	if err := redis.Init(config.Config.Redis); err != nil {
		logger.Error("redis init err", zap.Error(err))
		return
	}
	defer redis.Close()

	if err := snowflake.Init(config.Config.StartTime, config.Config.MachineId); err != nil {
		logger.Error("snowflake init err", zap.Error(err))
		return
	}

	r := router.Init(config.Config.Mode)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.Config.Port),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("server start err", zap.Error(err))
		}
	}()

	// 服务关闭接受通道
	quit := make(chan os.Signal, 1)

	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit // 等待关闭信号

	logger.Info("shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 五秒超时context
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("server shutdown err", zap.Error(err))
	}

	logger.Info("Bye~")
}
