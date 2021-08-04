package config

import (
	"github.com/Unknwon/goconfig"
	"log"
	"strconv"
)

// BUILD 开发环境
const BUILD = "prod"

//生产环境
//const BUILD = "prod"

// DbCfg Mysql数据库配置
type DbCfg struct {
	Username          string
	Password          string
	Database          string
	Host              string
	Port              string
	MaxIdleConnection int
	MaxOpenConnection int
	ShowType          string
}

// AppServer 应用程序配置
type AppServer struct {
	Port        string //App运行端口
	Lock        string //是否开启多终端登录 0开启 1不开启
	DemoEnabled bool   //是否是演示模式
}

// LoggerCfg 日志配置结构体
type LoggerCfg struct {
	LogPath string
	LogName string
}

// RedisCfg redis配置结构体
type RedisCfg struct {
	RedisHost string //地址
	Port      int64  //端口
	RedisPwd  string //密码
	RedisDB   int64  //数据库
	Timeout   int64  //超时时间
}

// MongoDb mongo配置结构体
type MongoDb struct {
	Url      string
	Port     string
	DB       string
	User     string
	Password string
}

var Cfg *goconfig.ConfigFile

func init() {
	var err error
	Cfg, err = goconfig.LoadConfigFile("./config/config-" + BUILD + ".ini")
	if err != nil {
		log.Fatal(err)
	}
	return
}

//读取mysql配置
func GetMysqlCfg() (mysql DbCfg) {
	mysql.Username, _ = Cfg.GetValue("mysql", "username")
	mysql.Password, _ = Cfg.GetValue("mysql", "password")
	mysql.Database, _ = Cfg.GetValue("mysql", "database")
	mysql.Host, _ = Cfg.GetValue("mysql", "host")
	mysql.Port, _ = Cfg.GetValue("mysql", "port")
	mysql.ShowType, _ = Cfg.GetValue("mysql", "sqlType")
	value, _ := Cfg.GetValue("mysql", "MaxIdleConnection")
	mysql.MaxIdleConnection, _ = strconv.Atoi(value)
	v, _ := Cfg.GetValue("mysql", "MaxOpenConnection")
	mysql.MaxOpenConnection, _ = strconv.Atoi(v)
	return mysql
}

//读取server配置
func GetServerCfg() (server AppServer) {
	server.Port, _ = Cfg.GetValue("app", "server")
	server.Lock, _ = Cfg.GetValue("app", "lock")
	demoEnabled, _ := Cfg.GetValue("app", "demoEnabled")
	server.DemoEnabled = demoEnabled == "0"
	return server
}

//获取Logger配置
func GetLoggerCfg() (logger LoggerCfg) {
	logger.LogPath, _ = Cfg.GetValue("logger", "logPath")
	logger.LogName, _ = Cfg.GetValue("logger", "logName")
	return logger
}

//获取mongo配置
func GetMongoCfg() (mongo MongoDb) {
	mongo.Url, _ = Cfg.GetValue("mongodb", "url")
	mongo.Port, _ = Cfg.GetValue("mongodb", "port")
	mongo.DB, _ = Cfg.GetValue("mongodb", "db")
	mongo.User, _ = Cfg.GetValue("mongodb", "user")
	mongo.Password, _ = Cfg.GetValue("mongodb", "password")
	return mongo
}

//获取redis配置
func GetRedisCfg() (r RedisCfg) {
	r.RedisHost, _ = Cfg.GetValue("redis", "host")
	getValue, _ := Cfg.GetValue("redis", "port")
	r.Port, _ = strconv.ParseInt(getValue, 10, 32)
	r.RedisPwd, _ = Cfg.GetValue("redis", "password")
	db, _ := Cfg.GetValue("redis", "db")
	r.RedisDB, _ = strconv.ParseInt(db, 10, 32)
	value, _ := Cfg.GetValue("redis", "timeout")
	r.Timeout, _ = strconv.ParseInt(value, 10, 64)
	return r
}
