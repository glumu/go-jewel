package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type ErrorLoggerData struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Func    string `json:"func"`
	Info    string `json:"info"`
}

type ErrorLogger struct {
	logger *logrus.Logger
	fl     *rotatelogs.RotateLogs // 日志切割
}

func NewErrorLogger(path string) (*ErrorLogger, error) {
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
	return &ErrorLogger{
		logger: logger,
		fl:     fl,
	}, nil
}

func (log *ErrorLogger) Record(logInfo interface{}) {
	if log != nil && log.logger != nil {
		d := Struct2Map(logInfo)
		msg := d["message"].(string)
		delete(d, "message")

		entry := log.logger.WithFields(d)
		entry.Error(msg)
	}
}

func (log *ErrorLogger) AsyncRecord(logInfo interface{}) {
	go log.Record(logInfo)
}

func (log *ErrorLogger) Close() {
	if log != nil && log.fl != nil {
		log.fl.Close()
	}
}
