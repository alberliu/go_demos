package main

import (
	"testing"
	"go.uber.org/zap"
	"time"
	"go.uber.org/zap/zapcore"
		"os"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func TestZapSugar(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url","www.baidu.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", "www.baidu.com")
}

func TestZap(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func TestLogrotate(t *testing.T) {
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		os.Stdout,
		zap.DebugLevel,
	)
	logger := zap.New(core)
	logger.Debug("this is debug log")
	logger.Info("this is info log")
}

func TestOne(t *testing.T) {
	encoder_cfg := zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	Curr_level := zap.NewAtomicLevelAt(zap.DebugLevel)

	custom_cfg := zap.Config{
		Level:            Curr_level,
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    encoder_cfg,
		OutputPaths:      []string{"stderr","qihu-secret-business.log"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, _ := custom_cfg.Build()
	defer logger.Sync()

	logger.Debug("adv_event_type_handle", zap.String("a", "1"))
	logger.Info("adv_event_type_handle",
		// Structured context as strongly-typed Field values.
		zap.String("url", "www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func BenchmarkZapSugar(b *testing.B) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	for i := 0; i < b.N; i++ {
		sugar.Infow("failed to fetch URL",
			"url", "http://example.com",
			"attempt", 3,
			"backoff", time.Second,
		)
		sugar.Infof("failed to fetch URL: %s", "http://example.com")
	}
}

func BenchmarkZapLogger(b *testing.B) {
	logger := zap.NewExample()
	defer logger.Sync()
	for i := 0; i < b.N; i++ {
		logger.Info("failed to fetch URL",
			zap.String("url", "http://example.com"),
			zap.Int("attempt", 3),
			zap.Duration("backoff", time.Second),
		)
	}
}
