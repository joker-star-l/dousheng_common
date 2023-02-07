package log

import "go.uber.org/zap"

var Log *zap.Logger
var Slog *zap.SugaredLogger

func init() {
	Log, _ = zap.NewDevelopment()
	Slog = Log.Sugar()
}

func Replace(env string, options ...zap.Option) {
	if env == "dev" {
		Log, _ = zap.NewDevelopment(options...)
	} else if env == "pro" {
		Log, _ = zap.NewProduction(options...)
	}
	Slog = Log.Sugar()
}
