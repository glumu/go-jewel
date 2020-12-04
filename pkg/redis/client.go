package redisclient

import (
	"time"

	redigo "github.com/garyburd/redigo/redis"
)

type RedisClient struct {
	pool *redigo.Pool
}

type Conn struct {
	conn redigo.Conn
}

type Config struct {
	MaxIdle     int  //最大连接数，即最多的tcp连接数，一般建议往大的配置
	MaxActive   int  // 最大空闲连接数，即会有这么多个连接提前等待着，但过了超时时间也会关闭。
	IdleTimeout int  // 空闲连接超时时间，但应该设置比redis服务器超时时间短。
	Wait        bool //如果超过最大连接，是报错，还是等待。
}

func NewRedisClient(endpoint string, passwd string, db int, conf *Config) (client *RedisClient, err error) {
	cnf := &Config{
		MaxIdle:     5,
		MaxActive:   500,
		IdleTimeout: 5,
		Wait:        true,
	}

	if conf != nil {
		cnf = conf
	}

	pool := &redigo.Pool{
		MaxIdle:     cnf.MaxIdle,
		MaxActive:   cnf.MaxActive,
		IdleTimeout: time.Duration(cnf.IdleTimeout) * time.Second,
		Wait:        cnf.Wait,
		Dial: func() (redigo.Conn, error) {
			con, err := redigo.Dial("tcp", endpoint,
				redigo.DialPassword(passwd),
				redigo.DialDatabase(db))
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
	return &RedisClient{
		pool: pool,
	}, nil
}

func (client *RedisClient) GetConn() *Conn {
	return &Conn{
		conn: client.pool.Get(),
	}
}

func (c *Conn) Get(key string) (string, error) {
	val, err := redigo.String(c.conn.Do("GET", key))
	if err != nil {
		if err.Error() == "redigo: nil returned" {
			return "", nil
		}
		return "", err
	}
	return val, nil
}

func (c *Conn) Del(key string) error {
	_, err := c.conn.Do("DEL", key)
	return err
}

func (c *Conn) Set(key string, val string, expire string) error {
	var err error
	if expire == "" {
		_, err = c.conn.Do("SET", key, val)
	} else {
		_, err = c.conn.Do("SET", key, val, "EX", expire)
	}
	return err
}

func (c *Conn) Append(key string, val string) error {
	res, err := c.Exists(key)
	if err != nil {
		return err
	}
	if res == true {
		_, err = c.conn.Do("APPEND", key, val)
	} else {
		err = c.Set(key, val, "")
	}
	return err
}

func (c *Conn) Expire(key string, val int) error {
	_, err := c.conn.Do("EXPIRE", key, val)
	return err
}

func (c *Conn) Exists(key string) (bool, error) {
	exists, err := redigo.Bool(c.conn.Do("EXISTS", key))
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (c *Conn) Close() {
	c.conn.Close()
}
