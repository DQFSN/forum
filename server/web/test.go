package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/v2/registry/consul"
	"log"
)

func main() {
	consulReg := consul.NewRegistry( //新建一个consul注册的地址，也就是我们consul服务启动的机器ip+端口
		registry.Addrs("127.0.0.1:8500"),
	)

	//ginRouter := gin.Default()
	//ginRouter.Handle("GET", "/user", func(context *gin.Context) {
	//	context.String(200, "user api")
	//})
	//ginRouter.Handle("GET", "/news", func(context *gin.Context) {
	//	context.String(200, "news api")
	//})
	//server := web.NewService( //go-micro很灵性的实现了注册和反注册，我们启动后直接ctrl+c退出这个server，它会自动帮我们实现反注册
	//	web.Name("prodservice"), //注册进consul服务中的service名字
	//	web.Address(":8001"), //注册进consul服务中的端口
	//	web.Handler(ginRouter), //web.Handler()返回一个Option，我们直接把ginRouter穿进去，就可以和gin完美的结合
	//	web.Registry(consulReg),//注册到哪个服务器伤的consul中
	//)
	//server.Run()

	userSrv, err := consulReg.GetService("user service")

	if err != nil {
		log.Fatal(err)
	}
	srv := selector.Random(userSrv)
	node, _ := srv()
	fmt.Println(node.Address, node.Id)

}
