package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var rdb *redis.Client

func initRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	return
}

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

//zap自带的初始化
func initLogger1() {
	logger, _ = zap.NewProduction()
}

//zap定制初始化
func initLogger() {
	writerSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)

	logger = zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConf := zap.NewProductionEncoderConfig()
	encoderConf.EncodeTime = zapcore.ISO8601TimeEncoder
	return zapcore.NewConsoleEncoder(encoderConf)
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer {
	file := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    1,     //最大200M
		MaxBackups: 5,     //备份数量
		MaxAge:     30,    //最大备份30天
		Compress:   false, //是否压缩
	}
	return zapcore.AddSync(file)
}

//初始化viper
func initViper() {
	//设置配置文件类型、文件名、文件目录
	viper.SetDefault("fileDir", "./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	//viper.SetConfigFile("config.yaml")
	viper.AddConfigPath("/etc/appName/")
	viper.AddConfigPath("$HOME/.appName")
	viper.AddConfigPath(".")
	//读配置
	err := viper.ReadInConfig()
	if err != nil {
		logger.Error(err.Error())
		return
	}
}

func main() {
	//redis相关操作
	if err := initRedis(); err != nil {
		return
	}
	defer rdb.Close()
	pipe := rdb.Pipeline()
	pipe.Incr("")
	pipe.Set("", "", 0)
	pipe.Get("").Val()
	_, _ = pipe.Exec()
	err := rdb.Watch(func(tx *redis.Tx) error {
		_, err := rdb.Pipelined(func(pipeliner redis.Pipeliner) error {
			pipeliner.Set("cn1", 1, 0)
			pipeliner.Incr("cn1")
			return nil
		})
		return err
	}, "key")
	if err != nil {
		return
	}

	//zap相关操作
	initLogger1()
	initLogger()
	for i := 0; i < 10000; i++ {
		logger.Info("test!",
			zap.String("method", http.MethodGet),
			zap.Int("status", http.StatusOK),
			zap.Duration("cost", time.Minute),
		)
		sugarLogger.Info("sugar!",
			zap.String("method", http.MethodGet),
			zap.Int("status", http.StatusOK),
			zap.Duration("cost", time.Minute),
		)
	}

	//实时监控配置变化
	viper.WatchConfig()
	//配置文件发生变更时会调用的回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config changed: ", e.Name)
	})

	//优雅关机：http.Server 内置的 Shutdown() 方法就支持优雅地关机
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		time.Sleep(time.Second * 5)
		c.String(http.StatusOK, "welcome to gin server")
	})
	// 第一种写法
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("listen: ", zap.String("error", err.Error()))
		}
	}()
	// 第二种写法
	//go func() {
	//	// 开启一个goroutine启动服务
	//	if err := http.ListenAndServe(":8080",r); err != nil && err != http.ErrServerClosed {
	//		logger.Fatal("listen: ", zap.String("error", err.Error()))
	//	}
	//}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1)
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // 阻塞在此，当接收到上述两种信号时才会往下执行
	fmt.Println("shutdown server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("server shutdown: ", zap.String("error", err.Error()))
	}
	fmt.Println("server exit done")
}
