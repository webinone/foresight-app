package libs

import (
	"github.com/Sirupsen/logrus"
	"github.com/rifflock/lfshook"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"time"
	"fmt"
	"runtime"
	"strings"
)

func GetLogger(params ...string) *logrus.Logger {

	customFormatter := new(logrus.JSONFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logrus.SetFormatter(customFormatter)

	logger := logrus.New()
	logger.Formatter = customFormatter
	logger.Level = logrus.DebugLevel

	var logId string
	if len(params) > 0 {
		logId = params[0]
	} else {
		// default log name
		logId = "app"
	}

	fmt.Println("!!!!!!!!!!!!!!!!!!!!! GetLogger : ", logId)

	logPath := "logs/" + logId + ".log"

	writer := rotatelogs.New(
		//logPath + ".%Y%m%d%H%M", // rotation pattern
		logPath + ".%Y%m%d", // rotation pattern
		rotatelogs.WithLinkName(logPath),
		rotatelogs.WithMaxAge(2 * 24 * time.Hour), // 2일동안 보관
		rotatelogs.WithRotationTime(24 * time.Hour), // rotate 1일
	)

	logger.Hooks.Add(lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel : writer,
		logrus.InfoLevel  : writer,
		logrus.ErrorLevel : writer,
		logrus.FatalLevel : writer,
	}))

	return logger
}

func FileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

