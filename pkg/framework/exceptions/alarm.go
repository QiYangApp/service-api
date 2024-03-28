package exceptions

import (
	"encoding/json"
	"framework/log"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type ErrorString struct {
	s string
}

type ErrorInfo struct {
	Time     string `json:"time"`
	Alarm    string `json:"alarm"`
	Message  string `json:"message"`
	Filename string `json:"filename"`
	Line     int    `json:"line"`
	FuncName string `json:"funcName"`
}

func (e *ErrorString) Error() string {
	return e.s
}

func New(text string) error {
	alarm("INFO", text)
	return &ErrorString{text}
}

// Email 发邮件
func Email(text string) error {
	alarm("EMAIL", text)
	return &ErrorString{text}
}

// Sms 发短信
func Sms(text string) error {
	alarm("SMS", text)
	return &ErrorString{text}
}

// WeChat 发微信
func WeChat(text string) error {
	alarm("WX", text)
	return &ErrorString{text}
}

// 告警方法
func alarm(level string, str string) {
	// 当前时间
	currentTime := time.Now().String()

	// 定义 文件名、行号、方法名
	fileName, line, functionName := "?", 0, "?"

	pc, fileName, line, ok := runtime.Caller(2)
	if ok {
		functionName = runtime.FuncForPC(pc).Name()
		functionName = filepath.Ext(functionName)
		functionName = strings.TrimPrefix(functionName, ".")
	}

	var msg = ErrorInfo{
		Time:     currentTime,
		Alarm:    level,
		Message:  str,
		Filename: fileName,
		Line:     line,
		FuncName: functionName,
	}

	jsons, errs := json.Marshal(msg)

	if errs != nil {
		log.Client.Sugar().Warnf("json marshal error: %v", errs)
	}

	errorJsonInfo := string(jsons)

	log.Client.Sugar().Warn(errorJsonInfo)

	if level == "EMAIL" {
		// 执行发邮件

	} else if level == "SMS" {
		// 执行发短信

	} else if level == "WX" {
		// 执行发微信

	} else if level == "INFO" {
		// 执行记日志
	}
}
