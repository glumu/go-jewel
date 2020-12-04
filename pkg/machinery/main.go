package gl_machinery

import (
	"fmt"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
)

type RedisConfig struct {
	MaxIdle int

	// Maximum number of connections allocated by the pool at a given time.
	// When zero, there is no limit on the number of connections in the pool.
	MaxActive int

	// Close connections after remaining idle for this duration in seconds. If the value
	// is zero, then idle connections are not closed. Applications should set
	// the timeout to a value less than the server's timeout.
	IdleTimeout int

	// If Wait is true and the pool is at the MaxActive limit, then Get() waits
	// for a connection to be returned to the pool before returning.
	Wait bool

	// ReadTimeout specifies the timeout in seconds for reading a single command reply.
	ReadTimeout int

	// WriteTimeout specifies the timeout in seconds for writing a single command.
	WriteTimeout int

	// ConnectTimeout specifies the timeout in seconds for connecting to the Redis server when
	// no DialNetDial option is specified.
	ConnectTimeout int

	// DelayedTasksPollPeriod specifies the period in milliseconds when polling redis for delayed tasks
	DelayedTasksPollPeriod int
}

type Config struct {
	Password    string //redis passowrd
	Endpoint    string // redis endpoint: http://120.0.0.1:6379
	Queue       string
	ConsumerTag string
	RedisConfig RedisConfig
}

type Logger struct {
	Level string
	Type  string
	Msg   string
}

type Machinery struct {
	Config   Config
	ErrorLog func(logger Logger)
	InfoLog  func(logger Logger)
	Server   *machinery.Server
	Debug    bool
}

func NewMachinery(conf Config, errorLog func(logger Logger), infoLog func(logger Logger)) (*Machinery, error) {
	var cnf = &config.Config{
		Broker:        fmt.Sprintf("redis://%s@%s", conf.Password, conf.Endpoint),
		DefaultQueue:  conf.Queue,
		ResultBackend: fmt.Sprintf("redis://%s@%s", conf.Password, conf.Endpoint),
		Redis: &config.RedisConfig{
			MaxIdle:                conf.RedisConfig.MaxIdle,
			MaxActive:              conf.RedisConfig.MaxActive,
			IdleTimeout:            conf.RedisConfig.IdleTimeout,
			Wait:                   conf.RedisConfig.Wait,
			ReadTimeout:            conf.RedisConfig.ReadTimeout,
			WriteTimeout:           conf.RedisConfig.WriteTimeout,
			ConnectTimeout:         conf.RedisConfig.ConnectTimeout,
			DelayedTasksPollPeriod: conf.RedisConfig.DelayedTasksPollPeriod,
		},
	}

	server, err := machinery.NewServer(cnf)
	if err != nil {
		return nil, err
	}
	return &Machinery{
		Config:   conf,
		ErrorLog: errorLog,
		InfoLog:  infoLog,
		Server:   server,
	}, nil
}
