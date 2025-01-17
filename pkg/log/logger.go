package log

import (
	"admin-api/common/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var log *logrus.Logger
var logToFile *logrus.Logger

type Formatter struct{}

// 自定义日志输出格式
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	log := entry.Time.Format("2006-01-02 15:04:05.000")
	log += " [" + strings.ToUpper(entry.Level.String()) + "]"
	if entry.Caller != nil {
		log += " [" + filepath.Base(entry.Caller.File) + ":" + strconv.Itoa(entry.Caller.Line) + "] "
	}
	log += entry.Message + "\n"
	return []byte(log), nil
}

func logFile() *logrus.Logger {
	if logToFile == nil {
		logPath := filepath.Join(config.Config.Log.Path, config.Config.Log.Name)
		logToFile = logrus.New()
		logToFile.SetReportCaller(true)
		logToFile.SetLevel(logrus.DebugLevel)
		logWriter, _ := rotatelogs.New(
			logPath+"_%Y%m%d.log",
			rotatelogs.WithMaxAge(30*24*time.Hour),    // 设置日志最大保存时间为30天
			rotatelogs.WithRotationTime(24*time.Hour), // 设置日志切割时间间隔为1天
		)
		// 创建文件写入器
		writeMap := lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.FatalLevel: logWriter,
			logrus.DebugLevel: logWriter,
			logrus.WarnLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.PanicLevel: logWriter,
		}
		// 绑定文件写入器和自定义日志格式到lfshook
		logToFile.AddHook(lfshook.NewHook(writeMap, &Formatter{}))
	}
	return logToFile
}

func Log() *logrus.Logger {
	// 输出日志到文件
	if config.Config.Log.Model == "file" {
		return logFile()
	} else {
		// 输出日志到控制台
		log = logrus.New()
		log.Out = os.Stdout
		log.SetFormatter(&Formatter{})
		log.SetReportCaller(true)
		log.SetLevel(logrus.DebugLevel)
	}
	return log
}
