package cron

import (
	"framework/log"
	"go.uber.org/zap"
	"strings"
	"time"
)

type Logger struct {
}

func (c Logger) Info(msg string, keysAndValues ...interface{}) {
	//keysAndValues = formatTimes(keysAndValues)
	//
	//log.Client().Sugar().Debug(
	//	zap.String("name", "cron"),
	//	zap.String("msg", msg),
	//	zap.Any("data", keysAndValues),
	//)

}

func (c Logger) Error(err error, msg string, keysAndValues ...interface{}) {
	keysAndValues = formatTimes(keysAndValues)

	log.Client().Sugar().Error(
		zap.Error(err),
		zap.String("msg", msg),
		zap.Any("data", keysAndValues),
	)
}

// formatString returns a logfmt-like format string for the number of
// key/values.
func formatString(numKeysAndValues int) string {
	var sb strings.Builder
	sb.WriteString("%s")
	if numKeysAndValues > 0 {
		sb.WriteString(", ")
	}
	for i := 0; i < numKeysAndValues/2; i++ {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("%v=%v")
	}
	return sb.String()
}

// formatTimes formats any time.Time values as RFC3339.
func formatTimes(keysAndValues []interface{}) []interface{} {
	var formattedArgs []interface{}
	for _, arg := range keysAndValues {
		if t, ok := arg.(time.Time); ok {
			arg = t.Format(time.RFC3339)
		}
		formattedArgs = append(formattedArgs, arg)
	}
	return formattedArgs
}
