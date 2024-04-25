package core

import (
	"blog_server/global"
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"time"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	logPrefix := global.Config.Logger.Prefix
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//	 自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)

		//	 自定义输出格式
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s %s \n", logPrefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s \n", logPrefix, timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger() *logrus.Logger {
	mLog := logrus.New()                                // 实例化
	mLog.SetOutput(os.Stdout)                           // 设置输出类型
	mLog.SetReportCaller(global.Config.Logger.ShowLine) // 开启返回函数名和行号
	mLog.SetFormatter(&LogFormatter{})                  // 设置自定义的formatter
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level) // 设置最低日志级别
	InitDefaultLogger()
	// todo: 设置输出到按日期命名的文件
	//logFilePath := getLogFilePath()
	//file, _ := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	//mLog.Out = file

	return mLog
}

func InitDefaultLogger() {
	// 全局log
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(global.Config.Logger.ShowLine)
	logrus.SetFormatter(&LogFormatter{})
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level) // 设置最低日志级别
}

func getLogFilePath() string {
	// 指定日志文件存放目录
	logDir := "logs"

	// 创建或确认目录存在
	if err := os.MkdirAll(logDir, 0o755); err != nil {
		panic(fmt.Sprintf("Failed to create log directory: %v", err))
	}

	// 生成以当前日期命名的日志文件名
	t := time.Now()
	fileName := fmt.Sprintf("%s_%s.log", t.Format("2006-01-02"), "app")

	// 返回完整文件路径
	return filepath.Join(logDir, fileName)
}
