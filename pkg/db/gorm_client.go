package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewDB(user, password, location, port, name string, logMode bool) (*gorm.DB, error) {
	dbArg := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		user, password, location, port, name, true, "Local",
	)

	db, err := gorm.Open("mysql", dbArg)
	if err != nil {
		return nil, err
	}

	if logMode {
		if err = db.LogMode(true).Error; err != nil {
			return nil, err
		}
	}

	// 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，
	// 可以避免并发太高导致连接mysql出现too many connections的错误。
	// db.DB().SetMaxOpenConns(20000)
	// 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	db.DB().SetMaxIdleConns(0)
	return db, nil
}
