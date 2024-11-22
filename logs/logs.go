package logs

import (
	"context"
	"github.com/Lu271/server-zys/internal/core"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"strings"
	"time"
)

var logging *zap.SugaredLogger

func InitLogger() {
	logWriter, _ := GetWriter(core.GlobalConfig.Logger.LogFile)

	fileCore := zapcore.NewCore(GetEncoder(zapcore.CapitalLevelEncoder), zapcore.AddSync(logWriter), LogLevel(core.GlobalConfig.Logger.LogLevel))
	consoleCore := zapcore.NewCore(GetEncoder(zapcore.CapitalColorLevelEncoder), zapcore.AddSync(os.Stdout), LogLevel(core.GlobalConfig.Logger.LogLevel))

	tee := zapcore.NewTee(fileCore, consoleCore)
	log := zap.New(tee, zap.AddCaller())
	logging = log.Sugar()
	return
}

func Info(ctx context.Context, log string) {
	WithContext(ctx).Infof(log)
}

func Warn(ctx context.Context, log string) {
	WithContext(ctx).Warnf(log)
}

func Debug(ctx context.Context, log string) {
	WithContext(ctx).Debugf(log)
}

func Error(ctx context.Context, log string) {
	WithContext(ctx).Errorf(log)
}

func WithContext(ctx context.Context) *zap.SugaredLogger {
	if ctx == nil {
		return logging
	}
	duration := (time.Now().UnixNano() - cast.ToInt64(ctx.Value("startTime"))) / int64(time.Millisecond)
	return logging.With("duration", duration).With("traceId", ctx.Value("traceId"))
}

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
