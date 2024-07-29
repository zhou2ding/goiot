package logger

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"goiot/pkg/conf"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
)

var Logger *zap.SugaredLogger

func InitLogger(routineName string) {
	//日志级别
	lv := new(zapcore.Level)
	if err := lv.UnmarshalText([]byte(viper.GetString("log.level"))); err != nil {
		_ = lv.UnmarshalText([]byte("info"))
	}
	path := conf.Conf.GetString("log.path")
	maxSize := conf.Conf.GetInt("log.size")
	maxAge := conf.Conf.GetInt("log.expire")
	maxBackups := conf.Conf.GetInt("log.limit")

	writeSyncer := make([]zapcore.WriteSyncer, 0)
	if len(path) > 0 {
		writeSyncer = append(writeSyncer, zapcore.AddSync(&lumberjack.Logger{
			Filename:   strings.TrimRight(path, "/") + "/" + routineName + ".log",
			LocalTime:  true,
			MaxSize:    maxSize,
			MaxAge:     maxAge,
			MaxBackups: maxBackups,
		}))
	}
	if conf.Conf.GetBool("log.stdout") {
		writeSyncer = append(writeSyncer, zapcore.AddSync(os.Stdout))
	}
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.NewMultiWriteSyncer(writeSyncer...),
		zap.NewAtomicLevelAt(*lv),
	)
	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.Development()).Sugar()
}
