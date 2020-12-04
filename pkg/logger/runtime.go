package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type RuntimeLoggerData struct {
	Level   string `json:"level"`
	Message string `json:"message"`
	Type    string `json:"type"`
	Func    string `json:"func"`
	Result  string `json:"result"`
	Info    string `json:"info"`
}

type RuntimeLogger struct {
	logger *logrus.Logger
	fl     *rotatelogs.RotateLogs // 日志切割
}

func NewRuntimeLogger(path string) (*RuntimeLogger, error) {
	fl, err := rotatelogs.New(path)
	//fl, err := rotatelogs.New(
	//	fmt.Sprintf("%s.%%Y%%m%%d", path),
	//	rotatelogs.WithLinkName(path),
	//	rotatelogs.WithRotationTime(24*time.Hour),
	//	rotatelogs.WithMaxAge(7*24*time.Hour),
	//)
	if err != nil {
		return nil, err
	}

	logger := logrus.New()
	logger.SetOutput(fl)
	logger.SetFormatter(&logrus.JSONFormatter{})
	return &RuntimeLogger{
		logger: logger,
		fl:     fl,
	}, nil
}

func (log *RuntimeLogger) Record(logInfo interface{}) {
	if log != nil && log.logger != nil {
		d := Struct2Map(logInfo)
		msg := ""
		level := ""
		if _, ok := d["message"]; ok == true {
			msg = d["message"].(string)
			delete(d, "message")
		}
		if _, ok := d["level"]; ok == true {
			level = d["level"].(string)
			delete(d, "level")
		}
		entry := log.logger.WithFields(d)

		if level != "" {
			switch level {
			case "DEBUG":
				entry.Debug(msg)
			case "ERROR":
				entry.Error(msg)
			case "WARNING":
				entry.Warning(msg)
			default:
				entry.Info(msg)
			}
		} else {
			entry.Info(msg)
		}
	}
}

func (log *RuntimeLogger) AsyncRecord(logInfo interface{}) {
	go log.Record(logInfo)
}

func (log *RuntimeLogger) Close() {
	if log != nil && log.fl != nil {
		log.fl.Close()
	}
}
