package rpcimpl

import (
	"blog/server/model"
	"fmt"
	"golang.org/x/net/context"
	"log"
	pb "blog/api"
	db "blog/server/db"
)

type BlogServer struct {

}

func (bs *BlogServer) PublishBlog(ctx context.Context, in *pb.PublishRequest) (*pb.PublishReply, error) {

	if len(in.Title) > 0 && len(in.Author) > 0 {
		mysqlDB := db.DB()
		err := mysqlDB.Create(&model.Blog{Title: in.Title, Content: in.Content, Author: in.Author}).Error
		if err != nil {
			return &pb.PublishReply{
				Status: fmt.Sprintf("publish blog : %s", err),
			},err
		}

		return &pb.PublishReply{
			Status: fmt.Sprintf("publish ok : %s", in.Title),
		},nil
	}

	log.Fatal("publish failed, title or author can not be empty")
	return &pb.PublishReply{
		Status: fmt.Sprintln("publish failed, title or author can not be empty"),
	},nil
}

func (bs *BlogServer) GetBlogs(ctx context.Context, in *pb.BlogsRequest) (out *pb.BlogsReply, err error) {
	mysqlDB := db.DB()

	var blogs []*pb.Blog
	fmt.Printf("author-->%s\n", in.Author)
	if len(in.Author) > 0 {
		mysqlDB.Where(pb.Blog{Author: in.Author}).Find(&blogs)
	}else{
		mysqlDB.Where(pb.Blog{}).Find(&blogs)
	}

	return &pb.BlogsReply{Blogs: blogs},nil

}
