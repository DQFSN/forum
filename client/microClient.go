package main

import (
	"context"
	"fmt"
	"time"

	"github.com/micro/go-micro/v2"
	mpb "github.com/DQFSN/blog/proto/micro"
)

func main() {
	// create and initialise a new service
	service := micro.NewService()
	service.Init()

	// create the proto client for Auth
	client := mpb.NewAuthService("blog", service.Client())

	// call an endpoint on the service
	rsp, err := client.LogIn(context.Background(), &mpb.LogInRequest{
		Email: "7@q.com",
		Password: "d",
	})
	if err != nil {
		fmt.Println("Error calling sigin: ", err)
		return
	}

	// print the response
	fmt.Println("Response: ", rsp.Status)

	// let's delay the process for exiting for reasons you'll see below
	time.Sleep(time.Second * 2)
}