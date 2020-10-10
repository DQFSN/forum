package main

import (
	mpb "github.com/DQFSN/blog/api/micro"
	"github.com/DQFSN/blog/server/microimp"
	micro "github.com/micro/go-micro/v2"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name("blog"),
		)

	service.Init()

	err := mpb.RegisterAuthHandler(service.Server(), microimp.AuthHandler{})
	if err != nil {
		log.Fatal(err)
	}
	err = mpb.RegisterPublishHandler(service.Server(), microimp.BlogServer{})
	if err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}