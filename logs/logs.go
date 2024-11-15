package logs

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"server-zys/internal/core"
	"strings"
	"time"
)

var logging *zap.SugaredLogger

// GetWriter 支持日志定时切割
func GetWriter(filename string) (logf io.Writer, err error) {
	logf, err = rotatelogs.New(filename+".%Y%m%d%H",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	return
}

// GetEncoder 自定义日志输出格式
func GetEncoder(level zapcore.LevelEncoder) zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "time",
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    level,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
}

// LogLevel 自定义日志输出级别
func LogLevel(level string) (LogLevel zapcore.Level) {
	switch strings.ToUpper(level) {
	case "DEBUG":
		LogLevel = zapcore.DebugLevel
	case "INFO":
		LogLevel = zapcore.InfoLevel
	case "WARN":
		LogLevel = zapcore.WarnLevel
	case "ERROR":
		LogLevel = zapcore.ErrorLevel
	default:
		LogLevel = zapcore.InfoLevel // 默认返回 InfoLevel
	}
	return
}

func InitLogger() {
	logWriter, _ := GetWriter(core.GlobalConfig.Logger.LogFile)

	fileCore := zapcore.NewCore(GetEncoder(zapcore.CapitalLevelEncoder), zapcore.AddSync(logWriter), LogLevel(core.GlobalConfig.Logger.LogLevel))
	consoleCore := zapcore.NewCore(GetEncoder(zapcore.CapitalColorLevelEncoder), zapcore.AddSync(os.Stdout), LogLevel(core.GlobalConfig.Logger.LogLevel))

	tee := zapcore.NewTee(fileCore, consoleCore)
	log := zap.New(tee, zap.AddCaller())
	logging = log.Sugar()
	logging.Infof("log init success")
	return
}

func Info(log string) {
	logging.Infof(log)
}

func Warn(log string) {
	logging.Warnf(log)
}

func Debug(log string) {
	logging.Debugf(log)
}

func Error(log string) {
	logging.Errorf(log)
}
