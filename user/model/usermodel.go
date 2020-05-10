package model

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
)

type (
	User struct {
		Id         int64
		Name       string    `xorm:"varchar(20) notnull 'name'"`
		Mobile     string    `xorm:"varchar(25) notnull unique 'mobile'"`
		Password   string    `xorm:"varchar(25) notnull 'mobile'"`
		CreateTime time.Time `xorm:"DateTime 'create_time'"`
	}

	UserModel struct {
		mysql      *xorm.Engine
		redisCache *redis.Client
		table      string
	}
)

func NewUserModel(mysql *xorm.Engine, redisCache *redis.Client, table string) *UserModel {
	return &UserModel{
		mysql:      mysql,
		redisCache: redisCache,
		table:      table,
	}
}
