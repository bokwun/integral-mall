package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"integral-mall/user/command/api/config"
	"integral-mall/user/controller"
	"integral-mall/user/logic"
	"integral-mall/user/model"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/yakaa/log4g"
)

// 运行go run接受参数
var configFile = flag.String("f", "config/config.json", "use config")

func main() {
	// 加载配置文件
	flag.Parse()
	//new一个config结构体
	conf := new(config.Config)
	// 读取文件内容
	bs, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	// 把读取到的json配置绑定conf结构体
	if err := json.Unmarshal(bs, conf); err != nil {
		log.Fatal(err)
	}
	// 判断当前模式，将log输入到log4g里
	// if conf.Mode == gin.ReleaseMode {
	// 	log4g.Init(log4g.Config{Path: "logs"})
	// 	gin.DefaultWriter = log4g.ErrorLog
	// 	gin.DefaultErrorWriter = log4g.ErrorLog
	// }

	// 默认将信息输入到logs
	log4g.Init(log4g.Config{Path: "logs"})
	gin.DefaultWriter = log4g.InfoLog
	gin.DefaultErrorWriter = log4g.ErrorLog

	//连接mysql数据库
	engine, err := xorm.NewEngine("mysql", conf.Mysql.DataSource)
	if err != nil {
		log.Fatal(err)
	}
	//连接redis数据库
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.DataSource,
		Password: conf.Redis.Auth,
	})
	userModel := model.NewUserModel(engine, client, conf.Mysql.Table.User)
	userLogic := logic.NewUserLogic(userModel, client)
	userController := controller.NewUserController(userLogic)

	r := gin.Default()
	userRouterGroup := r.Group("/user")
	{
		userRouterGroup.POST("/register", userController.Register)
		userRouterGroup.POST("/login", userController.Login)
	}

	log4g.Error(r.Run(conf.Port))
}
