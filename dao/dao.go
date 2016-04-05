package dao

import (
	"log"
    "restful/config"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "fmt"
)

var (
	RedisClients *redis.Pool
	DB           *gorm.DB
)

func Init() {
    initConfig()
	initRedis()
	InitDB()
}

//初始化配置文件
func initConfig() {
    config.Init()
}

//初始化数据库
func InitDB() {
	var err error
    
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", config.Config.DB.User, config.Config.DB.Password, config.Config.DB.Name))
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}

//初始化redis
func initRedis() *redis.Pool {
	RedisClients = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", config.Config.Redis.Host, config.Config.Redis.Port))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
	return RedisClients
}
