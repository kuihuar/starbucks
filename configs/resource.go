package configs



import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	Resource ResourceConf
)

type MysqlConf struct {
	Service         string        `yaml:"service"`
	DataBase        string        `yaml:"database"`
	Addr            string        `yaml:"addr"`
	User            string        `yaml:"user"`
	Password        string        `yaml:"password"`
	Charset         string        `yaml:"charset"`
	MaxIdleConns    int           `yaml:"maxidleconns"`
	MaxOpenConns    int           `yaml:"maxopenconns"`
	ConnMaxLifeTime time.Duration `yaml:"connMaxLifeTime"`
	ConnTimeOut     time.Duration `yaml:"ConnTimeOut"`
	WriteTimeOut    time.Duration `yaml:"writeTimeOut"`
	RadeTimeOut     time.Duration `yaml:"readTimeOut"`
}
type RedisConf struct {
	Service         string        `yaml:"service"`
	Addr            string        `yaml:"addr"`
	Password        string        `yaml:"password"`
	MaxIdle         int           `yaml:"maxIdle"`
	MaxActive       int           `yaml:"maxActive"`
	IdleTimeout     time.Duration `yaml"idletimeout"`
	MaxConnLifetime time.Duration `yaml:"maxConnLifetime"`
	ConnTimeOut     time.Duration `yaml:"ConnTimeOut"`
	WriteTimeOut    time.Duration `yaml:"writeTimeOut"`
	ReadeTimeOut     time.Duration `yaml:"readTimeOut"`
}
type ResourceConf struct {
	Mysql map[string]MysqlConf
	Redis map[string]RedisConf
}

var DBS map[string]*gorm.DB


func initDB (){
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {

	}
	DBS[dsn] = db
}
func InitResource(){
	InitMysql()
	InitRedis()
}
