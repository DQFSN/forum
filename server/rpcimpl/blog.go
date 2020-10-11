package rpcimpl

import (
	"fmt"
	gpb "github.com/DQFSN/blog/proto/grpc"
	db "github.com/DQFSN/blog/server/db"
	"github.com/DQFSN/blog/server/model"
	"golang.org/x/net/context"
	"log"
)

type BlogServer struct {

}

func (bs *BlogServer) PublishBlog(ctx context.Context, in *gpb.PublishRequest) (*gpb.PublishReply, error) {

	if len(in.Title) > 0 && len(in.Author) > 0 {
		mysqlDB := db.DB()
		err := mysqlDB.Create(&model.Blog{Title: in.Title, Content: in.Content, Author: in.Author}).Error
		if err != nil {
			return &gpb.PublishReply{
				Status: fmt.Sprintf("publish blog : %s", err),
			},err
		}

		return &gpb.PublishReply{
			Status: fmt.Sprintf("publish ok : %s", in.Title),
		},nil
	}

	log.Fatal("publish failed, title or author can not be empty")
	return &gpb.PublishReply{
		Status: fmt.Sprintln("publish failed, title or author can not be empty"),
	},nil
}

func (bs *BlogServer) GetBlogs(ctx context.Context, in *gpb.BlogsRequest) (out *gpb.BlogsReply, err error) {
	mysqlDB := db.DB()

	var blogs []*gpb.Blog
	fmt.Printf("author-->%s\n", in.Author)
	if len(in.Author) > 0 {
		mysqlDB.Where(gpb.Blog{Author: in.Author}).Find(&blogs)
	}else{
		mysqlDB.Where(gpb.Blog{}).Find(&blogs)
	}

	return &gpb.BlogsReply{Blogs: blogs},nil

}

func (bs *BlogServer) ModifyBlog(ctx context.Context, in *gpb.ModifyBlogRequest) (out *gpb.ModifyBlogReply, err error) {
	mysqlDB := db.DB()

	err = mysqlDB.Where(gpb.ModifyBlogRequest{Id: in.Id}).Error

	if err != nil {
		out.Status = fmt.Sprintf("update blog err ：%v", err)
		return out,nil
	}
	out.Status = "update blog Ok"
	return out,nil
}
