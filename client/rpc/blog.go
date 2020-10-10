package rpc

import (
	grpc2 "blog/api/grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)


func PublishBlog(title, content, author string) string {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("connect err ", err)
	}
	defer conn.Close()
	c := grpc2.NewPublishClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.PublishBlog(ctx, &grpc2.PublishRequest{Title: title, Content: content, Author: author})
	if err != nil {
		log.Fatalf("counld not publish: %v", err)
	}
	log.Printf("publish blog :%s", r.Status)

	return r.Status
}