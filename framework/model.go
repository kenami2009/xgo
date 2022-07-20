package framework

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
)
import "gorm.io/gorm"
import "gorm.io/driver/mysql"

var Db *gorm.DB

func getMysqlDsn() string {
	conf, err := ini.Load("./framework/config.ini")
	if err != nil {
		log.Fatal("配置文件获取错误")
	}
	section := conf.Section("mysql")
	host := section.Key("host").String()
	port := section.Key("port").String()
	user := section.Key("user").String()
	password := section.Key("password").String()
	dbname := section.Key("dbname").String()
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
}

//初始化数据库连接
func InitDb() {
	dsn := getMysqlDsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	Db = db
}
