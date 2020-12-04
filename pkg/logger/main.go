package logger

import (
	"errors"
	"reflect"
)

type LoggerIF interface {
	Record(logInfo interface{})
	Close()
}

type level int

const (
	// DEBUG level
	DEBUG level = iota
	// INFO level
	INFO
	// WARNING level
	WARNING
	// ERROR level
	ERROR
	// FATAL level
	FATAL
)

var Lvl = map[level]string{
	DEBUG:   "DEBUG",
	INFO:    "INFO",
	WARNING: "WARNING",
	ERROR:   "ERROR",
	FATAL:   "FATAL",
}

type ResourcePath struct {
	Path             string
	IdentityEndpoint string
}

type Paths struct {
	ErrorLogPath    string
	AcceccLogPath   string
	RuntimeLogPath  string
	ResourceLogPath ResourcePath
}

type Logger struct {
	AccLog      *AccessLogger
	RunLog      *RuntimeLogger
	ErrLog      *ErrorLogger
	ResourceLog *ResourceLogger
}

func (l *Logger) Record(data interface{}) {
	dType := reflect.TypeOf(data)
	switch dType.Name() {
	case "ResourceLogHeadData":
		if l.ResourceLog != nil {
			l.ResourceLog.RecordHead(data.(ResourceLogHeadData))
		}
	case "ResourceLogFootData":
		if l.ResourceLog != nil {
			l.ResourceLog.RecordFoot(data.(ResourceLogFootData))
		}
	case "LogResourceContentData":
		if l.ResourceLog != nil {
			l.ResourceLog.Record(data)
		}
	case "AccessLoggerData":
		if l.AccLog != nil {
			l.AccLog.Record(data)
		}
	case "ErrorLoggerData":
		if l.ErrLog != nil {
			l.ErrLog.Record(data)
		}
	case "RuntimeLoggerData":
		fallthrough
	default:
		if l.RunLog != nil {
			l.RunLog.Record(data)
		}
	}
}

func (l *Logger) Close() {
	if l.AccLog != nil {
		l.AccLog.Close()
	}
	if l.ErrLog != nil {
		l.ErrLog.Close()
	}
	if l.RunLog != nil {
		l.RunLog.Close()
	}
	if l.ResourceLog != nil {
		l.ResourceLog.Close()
	}
}

func NewLogger(paths Paths) (*Logger, []error) {
	var accLog *AccessLogger
	var runLog *RuntimeLogger
	var errLog *ErrorLogger
	var resLog *ResourceLogger
	var errs []error
	var err error

	if paths.AcceccLogPath != "" {
		accLog, err = NewAccessLogger(paths.AcceccLogPath)
		if err != nil {
			errs = append(errs, err)
		} else if accLog == nil {
			errs = append(errs, errors.New("access logger is nil."))
		}
	}
	if paths.RuntimeLogPath != "" {
		runLog, err = NewRuntimeLogger(paths.RuntimeLogPath)
		if err != nil {
			errs = append(errs, err)
		} else if runLog == nil {
			errs = append(errs, errors.New("runtime logger is nil."))
		}
	}
	if paths.ErrorLogPath != "" {
		errLog, err = NewErrorLogger(paths.ErrorLogPath)
		if err != nil {
			errs = append(errs, err)
		} else if errLog == nil {
			errs = append(errs, errors.New("error logger is nil."))
		}
	}
	if paths.ResourceLogPath.Path != "" {
		resLog, err = NewResourceLogger(paths.ResourceLogPath.Path, paths.ResourceLogPath.IdentityEndpoint)
		if err != nil {
			errs = append(errs, err)
		} else if resLog == nil {
			errs = append(errs, errors.New("resource logger is nil."))
		}
	}

	return &Logger{
		AccLog:      accLog,
		RunLog:      runLog,
		ErrLog:      errLog,
		ResourceLog: resLog,
	}, nil
}
