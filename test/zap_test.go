package test

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
)

func TestZap(t *testing.T) {

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		func() zapcore.WriteSyncer {

			x, _ := os.Create("./xzz.log")

			return zapcore.AddSync(x)
		}(),
		zap.NewAtomicLevel(),
	), zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.RegisterHooks(core, func(entry zapcore.Entry) error {
			return nil
		})
	}), zap.AddStacktrace(zap.ErrorLevel))

	logger.Panic("sdfaadfa")

}
