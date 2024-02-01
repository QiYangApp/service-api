package logger

import "go.uber.org/zap"

type Instance struct {
	logger *CoreBuilder
}

func (i *Instance) Logger() *zap.Logger {
	return i.logger.Logger()
}

func (i Instance) Debug(ctx string, fields ...zap.Field) {
	i.logger.Logger().Debug(ctx, fields...)
}

func (i Instance) Debugf(ctx string, fields ...interface{}) {
	i.logger.Logger().Sugar().Debugf(ctx, fields...)
}

func (i Instance) Debugln(fields ...interface{}) {
	i.logger.Logger().Sugar().Debugln(fields...)
}

func (i Instance) Info(ctx string, fields ...zap.Field) {
	i.logger.Logger().Info(ctx, fields...)
}

func (i Instance) Infof(ctx string, fields ...interface{}) {
	i.logger.Logger().Sugar().Infof(ctx, fields...)
}

func (i Instance) Infoln(fields ...interface{}) {
	i.logger.Logger().Sugar().Infoln(fields...)
}

func (i Instance) Warn(ctx string, fields ...zap.Field) {
	i.logger.Logger().Warn(ctx, fields...)
}

func (i Instance) Warnf(ctx string, fields ...interface{}) {
	i.logger.Logger().Sugar().Warnf(ctx, fields...)
}

func (i Instance) Warnln(fields ...interface{}) {
	i.logger.Logger().Sugar().Warnln(fields...)
}

func (i Instance) Error(ctx string, fields ...zap.Field) {
	i.logger.Logger().Error(ctx, fields...)
}

func (i Instance) Errorf(ctx string, fields ...interface{}) {
	i.logger.Logger().Sugar().Errorf(ctx, fields...)
}

func (i Instance) Errorln(fields ...interface{}) {
	i.logger.Logger().Sugar().Errorln(fields...)
}

func (i *Instance) DPanic(ctx string, args ...zap.Field) {
	i.logger.Logger().DPanic(ctx, args...)
}

func (i *Instance) Panic(ctx string, args ...zap.Field) {
	i.logger.Logger().Panic(ctx, args...)
}

func (i Instance) Fatal(ctx string, fields ...zap.Field) {
	i.logger.Logger().Fatal(ctx, fields...)

}

func (i Instance) Fatalf(ctx string, fields ...interface{}) {
	i.logger.Logger().Sugar().Fatalf(ctx, fields...)
}

func (i Instance) Fatalln(ctx string, fields ...interface{}) {
	i.logger.Logger().Sugar().Fatalf(ctx, fields...)
}
