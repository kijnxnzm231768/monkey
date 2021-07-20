package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"monkey-admin/config"
	"monkey-admin/pkg/common"
	redisTool "monkey-admin/pkg/redistool"
	"time"
)

// X 全局DB
var (
	//MongoSession *mgo.Session
	//MongoDB      *mgo.DialInfo
	SqlDB   *xorm.Engine
	RedisDB *redisTool.RedisClient
)

func init() {
	var err error
	//配置mysql数据库
	mysql := config.GetMysqlCfg()
	jdbc := mysql.Username + ":" + mysql.Password + "@tcp(" + mysql.Host + ":" + mysql.Port + ")/" + mysql.Database + "?charset=utf8&parseTime=True&loc=Local"
	SqlDB, _ = xorm.NewEngine(mysql.ShowType, jdbc)

	if err != nil {
		log.Fatalf("db error: %#v\n", err.Error())
	}

	err = SqlDB.Ping()
	if err != nil {
		log.Fatalf("db connect error: %#v\n", err.Error())
	}
	SqlDB.SetMaxIdleConns(10)
	SqlDB.SetMaxOpenConns(100)
	_ = SqlDB.Sync2(
	//new(model.User),
	)
	timer := time.NewTicker(time.Minute * 30)
	go func(x *xorm.Engine) {
		for _ = range timer.C {
			err = x.Ping()
			if err != nil {
				log.Fatalf("db connect error: %#v\n", err.Error())
			}
		}
	}(SqlDB)
	SqlDB.ShowSQL(true)
	//初始化redis开始
	redisCfg := config.GetRedisCfg()
	redisOpt := common.RedisConnOpt{
		true,
		redisCfg.RedisHost,
		int32(redisCfg.Port),
		redisCfg.RedisPwd,
		int32(redisCfg.RedisDB),
		240,
	}

	RedisDB = redisTool.NewRedis(redisOpt)
	//配置redis结束
	// 配置Mongo
	//mongoCfg := config.GetMongoCfg()
	//uri := "mongodb://" + mongoCfg.User + ":" + mongoCfg.Password + "@" + mongoCfg.Url + ":" + mongoCfg.Port + "/" + mongoCfg.DB + ""
	//MongoDB, err = mgo.ParseURL(uri)
	//MongoSession, err = mgo.Dial(uri)
	//if err != nil {
	//	fmt.Printf("Can't connect to mongo, go error %v\n", err)
	//	panic(err.Error())
	//}
	//MongoSession.SetSafe(&mgo.Safe{})

	//进行缓存数据存贮
	//saveCache()
}
