package main

import (
	"context"
	"fmt"
	"github.com/DQFSN/blog/config"
	"github.com/micro/go-micro/v2/registry"
	"time"

	mpb "github.com/DQFSN/blog/proto/micro"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/v2/registry/consul"
)

func main() {
	//consul配置和注册
	config := config.Get().Consul
	consulReg := consul.NewRegistry(
			registry.Addrs(config.Host+":"+config.Port),
		)

	//此服务的服务注册发现中心设置为consul
	service := micro.NewService(
		micro.Registry(consulReg),
		)
	service.Init()

	// create the proto client for Auth
	client := mpb.NewAuthService("user service", service.Client())

	{
		rsp, err := client.SignUp(context.Background(), &mpb.SignUpRequest{
			Email: "6@q.com",
			Password: "123456",
			PasswordCheck: "123456",
			AuthCode: "s",
		})

		if err != nil {
			fmt.Println("Error calling SignUp: ", err)
		}
		// print the response

		if rsp != nil {
			fmt.Println("Response: ", rsp.Status)
		}
	}

	{
		//call an endpoint on the service
		rsp, err := client.LogIn(context.Background(), &mpb.LogInRequest{
			Email: "6@q.com",
			Password: "123456",
		})

		if err != nil {
			fmt.Println("Error calling LogIn: ", err)
		}

		if rsp != nil {
			fmt.Println("Response: ", rsp.Status)
		}

	}



	// let's delay the process for exiting for reasons you'll see below
	time.Sleep(time.Second * 2)
}