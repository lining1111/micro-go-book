package main

import (
	"micro-go-book/ch13-seckill/pkg/bootstrap"
	conf "micro-go-book/ch13-seckill/pkg/config"
	"micro-go-book/ch13-seckill/pkg/mysql"
	"micro-go-book/ch13-seckill/sk-admin/setup"
)

func main() {
	mysql.InitMysql(conf.MysqlConfig.Host, conf.MysqlConfig.Port, conf.MysqlConfig.User, conf.MysqlConfig.Pwd, conf.MysqlConfig.Db) // conf.MysqlConfig.Db
	//setup.InitEtcd()
	setup.InitZk()
	setup.InitServer(bootstrap.HttpConfig.Host, bootstrap.HttpConfig.Port)

}
