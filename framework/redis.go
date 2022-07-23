package framework

import (
	"github.com/go-ini/ini"
	"log"
)
import "github.com/go-redis/redis"

var Redis *redis.Client

func InitRedis() {
	conf, err := ini.Load("./framework/config.ini")
	if err != nil {
		log.Fatal("配置文件获取错误")
	}
	section := conf.Section("redis")
	host := section.Key("host").String()
	port := section.Key("port").String()
	addr := host + ":" + port

	rdb := redis.NewClient(&redis.Options{
		Addr: addr})
	_, err = rdb.Ping().Result()
	if err != nil {
		Redis = nil
		return
	}
	Redis = rdb
}
