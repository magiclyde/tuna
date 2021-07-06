/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/7/6 上午11:20
 * @note:
 */

package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var DefaultLogger = NewLogger(zap.InfoLevel, "")

func NewLogger(level zapcore.Level, filename string) (logger *zap.Logger) {
	// 编码器配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 日志输出配置
	var ws []zapcore.WriteSyncer
	ws = append(ws, zapcore.AddSync(os.Stdout)) // 打印到控制台

	if len(filename) > 0 {
		// 文件归档配置
		hook := lumberjack.Logger{
			Filename:   filename, // 日志文件路径
			MaxSize:    128,      // 每个日志文件保存的最大尺寸 单位：M
			MaxBackups: 30,       // 日志文件最多保存多少个备份
			MaxAge:     7,        // 文件最多保存多少天
			Compress:   true,     // 是否压缩
		}

		ws = append(ws, zapcore.AddSync(&hook)) // 打印到文件
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	// new core
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(ws...),
		atomicLevel,
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()

	// 开启文件及行号
	development := zap.Development()

	// 构造日志
	logger = zap.New(core, caller, development)

	return
}
