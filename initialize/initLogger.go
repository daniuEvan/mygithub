/*
 * @date: 2021/12/15
 * @desc: ...
 */

package initialize

import (
	"github.com/daniuEvan/mygithub/global"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

// initLogger 初始化logger
func initLogger() {
	logInfo := global.ServerConfig.LogInfo
	var (
		level        = logInfo.LogLevel
		logPath      = logInfo.LogPath
		logInConsole = logInfo.LogInConsole
		maxSize      = logInfo.MaxSize
		maxBackups   = logInfo.MaxBackups
		maxAge       = logInfo.MaxAge
		compress     = logInfo.Compress
	)

	// 设置日志级别
	//atomicLevel := zap.NewAtomicLevel()
	logLevel := getLogLevel(level)
	//atomicLevel.SetLevel(logLevel)

	// 状态配置
	writeSyncer := getLogWriter(logPath, maxSize, maxBackups, maxAge, compress, logInConsole)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)
	global.Logger = zap.New(core, zap.AddCaller())
}

// initSpiderLogger 初始化spider logger
func initSpiderLogger() {
	logInfo := global.ServerConfig.LogInfo
	var (
		level        = "info"
		logPath      = strings.SplitAfterN(logInfo.LogPath, "/", 2)[0] + "/spider.log"
		logInConsole = logInfo.LogInConsole
		maxSize      = logInfo.MaxSize
		maxBackups   = logInfo.MaxBackups
		maxAge       = logInfo.MaxAge
		compress     = logInfo.Compress
	)

	// 设置日志级别
	//atomicLevel := zap.NewAtomicLevel()
	logLevel := getLogLevel(level)
	//atomicLevel.SetLevel(logLevel)

	// 状态配置
	writeSyncer := getLogWriter(logPath, maxSize, maxBackups, maxAge, compress, logInConsole)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)
	global.SpiderLogger = zap.New(core, zap.AddCaller())
}

// getLogLevel 获取日志级别
func getLogLevel(logLevel string) zapcore.Level {
	logLevel = strings.ToLower(logLevel)
	var LogLevelInfo = map[string]zapcore.Level{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel,
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dpanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}
	level, ok := LogLevelInfo[logLevel]
	if !ok {
		level = zapcore.InfoLevel
	}
	return level
}

// getEncoder
func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // FullCallerEncoder 全路径编码器  ShortCallerEncoder 简单路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getLogWriter
func getLogWriter(logPath string, maxSize, maxBackups, maxAge int, compress, logInConsole bool) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logPath,    // 日志的位置
		MaxSize:    maxSize,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: maxBackups, // 保留旧文件的最大个数
		MaxAge:     maxAge,     // 保留旧文件的最大天数
		Compress:   compress,   // 是否压缩/归档旧文件
	}
	defer lumberJackLogger.Close()
	syncer := zapcore.AddSync(lumberJackLogger)
	if logInConsole {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return syncer
}
