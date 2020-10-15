package microimp

import (
	"context"
	mpb "github.com/DQFSN/blog/proto/micro"
	"testing"
)

func Test_PublishBlog(t *testing.T) {

	blogServer := BlogServer{}
	requests := []mpb.PublishRequest{
		{Title: "da", Content: "", Author: "7@q.com"},
		{Title: "da", Content: "", Author: "sds"},
		{Title: "", Content: "dsdsa", Author: ""},
		{Title: "das", Content: "", Author: "dasd"},
	}

	reply := mpb.PublishReply{}
	for _, in := range requests {
		err := blogServer.PublishBlog(context.TODO(), &in, &reply)

		if reply.Status == "" {
			t.Errorf("测试失败， %v",err)
		}else {
			t.Logf("测试成功， %v",reply.Status)
		}
	}
}

func Test_GetBlogs(t *testing.T) {

	blogServer := BlogServer{}
	requests := []mpb.BlogsRequest{
		{Author: "7@q.com"},
		{Author: ""},
	}

	reply := mpb.BlogsReply{}
	for _, in := range requests {
		err := blogServer.GetBlogs(context.TODO(), &in, &reply)

		if reply.Blogs == nil {
			t.Errorf("测试失败， %v",err)
		}else {
			t.Logf("测试成功， %v",reply.Blogs)
		}
	}
}
func Test_ModifyBlog(t *testing.T) {

	blogServer := BlogServer{}
	requests := []mpb.ModifyBlogRequest{
		{Title: "",Content: ""},
		{Title: "sdsda",Content: ""},
		{Title: "",Content: "dasdad"},
		{Title: "dadas",Content: "dada"},
	}

	reply := mpb.ModifyBlogReply{}
	for _, in := range requests {
		err := blogServer.ModifyBlog(context.TODO(), &in, &reply)

		if reply.Status == "" {
			t.Errorf("测试失败， %v",err)
		}else {
			t.Logf("测试成功， %v",reply.Status)
		}
	}
}
