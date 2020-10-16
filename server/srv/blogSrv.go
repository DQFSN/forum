package main

import (
	"github.com/DQFSN/forum/config"
	mpb "github.com/DQFSN/forum/proto/micro"
	"github.com/DQFSN/forum/server/microimp"
	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/v2/registry/consul"

	"log"
)

func main() {

	// 获取consul配置
	conf := config.Get().Consul
	address := conf.Host + ":" + conf.Port
	consulReg := consul.NewRegistry(
		registry.Addrs(address),
	)

	service := micro.NewService(
		micro.Name("blog service"),
		micro.Registry(consulReg),
	)

	service.Init()

	err := mpb.RegisterPublishHandler(service.Server(), microimp.BlogServer{})
	if err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}