package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"integral-mall/integral/command/rpc/config"
	"integral-mall/integral/logic"
	"integral-mall/integral/model"
	"integral-mall/integral/protos"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/yakaa/grpcx"
	"github.com/yakaa/log4g"
	"google.golang.org/grpc"
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
	integralModel := model.NewIntegralModel(engine, client, conf.Mysql.Table.Integral)
	// userServerLogic, err := logic.NewIntegralLogic(
	// 	conf.RabbitMq.DataSource+conf.RabbitMq.VirtualHost,
	// 	conf.RabbitMq.QueueName,
	// 	integralModel,
	// )
	// "amqp://admin:123456@localhost:5672/",
	fmt.Println("conf.RabbitMq.DataSource:", conf.RabbitMq.DataSource)
	fmt.Println("conf.RabbitMq.QueueName:", conf.RabbitMq.QueueName)
	userServerLogic, err := logic.NewIntegralLogic(
		conf.RabbitMq.DataSource+conf.RabbitMq.VirtualHost,
		conf.RabbitMq.QueueName,
		integralModel,
	)
	if err != nil {
		log.Fatal(err)
	}
	userServerLogic.ConsumeMessage()
	userServerLogic.PushMessage("INSERT INTO `integral` (user_id, integral) VALUES(22,90)")
	defer userServerLogic.CloseRabbitMqConn()
	rpcServer, err := grpcx.MustNewGrpcxServer(conf.RpcServerConfig, func(server *grpc.Server) {
		protos.RegisterUserRpcServer(server, userServerLogic)
	})
	if err != nil {
		log.Fatal(err)
	}
	log4g.Error(rpcServer.Run())
}
