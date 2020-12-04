package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type AccessLoggerData struct {
	Level        string `json:"level"`
	Message      string `json:"message"`
	Uri          string `json:"uri"`
	RequestId    string `json:"request_id"`
	Type         string `json:"type"` // request response
	Request      string `json:"request"`
	Response     string `json:"response"`
	ResponseCode string `json:"response_code"`
	RequestTime  string `json:"request_time"`
	ResponseTime string `json:"response_time"`
	Cost         string `json:"cost"`
}

type AccessLogger struct {
	logger *logrus.Logger
	fl     *rotatelogs.RotateLogs // 日志切割
}

func NewAccessLogger(path string) (*AccessLogger, error) {
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
	return &AccessLogger{
		logger: logger,
		fl:     fl,
	}, nil
}

func (log *AccessLogger) Record(logInfo interface{}) {
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

func (log *AccessLogger) AsyncRecord(logInfo interface{}) {
	go log.Record(logInfo)
}

func (log *AccessLogger) Close() {
	if log != nil && log.fl != nil {
		log.fl.Close()
	}
}
