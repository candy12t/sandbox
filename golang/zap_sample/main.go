package main

import (
	"log"
	"path"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	l, err := newLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer l.Sync()

	l.Infow("hoge", "direction", "satori_to_kintone")
	l.Errorw("error", "direction", "satori_to_kintone")
	l.Warnf("warn", "direction", "satori_to_kintone")
}

func newLogger() (*zap.SugaredLogger, error) {
	outputPath := path.Join("/", "tmp", "hoge.log")

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		DisableStacktrace: true,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{outputPath, "stdout"},
		ErrorOutputPaths: []string{outputPath, "stderr"},
	}

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	return logger.Sugar(), nil
}
