package main

import (
	"context"
	"fmt"
	"time"

	mpb "github.com/DQFSN/blog/proto/micro"
	"github.com/micro/go-micro/v2"
)

func main() {
	// create and initialise a new service
	service := micro.NewService()
	service.Init()

	// create the proto client for Auth
	client := mpb.NewAuthService("blog", service.Client())

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