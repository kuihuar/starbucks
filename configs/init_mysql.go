package configs

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	MysqlClientBiz1 *gorm.DB
	MysqlClientBiz2 *gorm.DB
	//MysqlClients map[string]*gorm.DB
)

func InitMysql()  {
	for name, conf := range Resource.Mysql{
		//MysqlClients[name],_ = initMysqlClient(conf)
		switch name {
		case "biz1":
			MysqlClientBiz1,_ = initMysqlClient(conf)
		case "biz2":
			MysqlClientBiz2,_ = initMysqlClient(conf)
		}
	}
}

func initMysqlClient(conf MysqlConf) (*gorm.DB, error) {
	c := &gorm.Config{}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?timeout=%s&readTimeout=%s&writeTimeout=%s&parseTime=True&loc=Asia%%2FShanghai",
		conf.User,conf.Password,conf.Addr,conf.DataBase,conf.ConnTimeOut,conf.RadeTimeOut,conf.WriteTimeOut)
	client, err := gorm.Open(mysql.Open(dsn), c)

	if err != nil {
		return nil, err
	}

	db , err := client.DB()

	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetConnMaxLifetime(conf.ConnMaxLifeTime)
	return client,nil

}