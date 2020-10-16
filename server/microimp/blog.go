package microimp

import (
	"fmt"
	mpb "github.com/DQFSN/forum/proto/micro"
	db "github.com/DQFSN/forum/server/db"
	"github.com/DQFSN/forum/server/model"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"time"
)

type BlogServer struct {
}

func (bs BlogServer) PublishBlog(ctx context.Context, in *mpb.PublishRequest, out *mpb.PublishReply) error {

	if len(in.Title) > 0 && len(in.Author) > 0 {
		mysqlDB := db.DB()
		err := mysqlDB.Create(&model.Blog{Title: in.Title, Content: in.Content, Author: in.Author}).Error
		if err != nil {
			out.Status = fmt.Sprintf("publish blog : %s", err)
			return err
		}

		out.Status = fmt.Sprintf("publish ok : %s", in.Title)
		return nil
	}

	out.Status = fmt.Sprintln("publish failed, title or author can not be empty")
	return nil
}

func (bs BlogServer) GetBlogs(ctx context.Context, in *mpb.BlogsRequest, out *mpb.BlogsReply) error {
	mysqlDB := db.DB()

	var blogs []*mpb.Blog
	if len(in.Author) > 0 {
		mysqlDB.Where(model.Blog{},"deleted_at IS NULL and author = ? ", in.Author).Find(&blogs)
	} else {
		mysqlDB.Where(model.Blog{}).Find(&blogs)
	}

	out.Blogs = blogs
	return nil

}

func (bs BlogServer) ModifyBlog(ctx context.Context, in *mpb.ModifyBlogRequest, out *mpb.ModifyBlogReply) (err error) {

	//过滤无效请求
	if in.Id == 0 || in.Author == "" {
		out.Status = fmt.Sprintf("update blog err : id and author can not be empty")
		return nil
	}

	mysqlDB := db.DB()

	blog := model.Blog{Model: gorm.Model{ID: uint(in.Id)}}
	blogsDB := mysqlDB.Find(&blog, "id = ? and deleted_at IS NULL", blog.ID)

	if blogsDB.RowsAffected == 0 {
		out.Status = fmt.Sprintf("update blog err : no such blog")
		return nil
	} else {
		//id 和 author 需匹配
		if blog.Author == in.Author {
			err := mysqlDB.Save(&mpb.Blog{Id: in.Id, Title: in.Title, Content: in.Content, Author: in.Author}).Error
			if err != nil {
				out.Status = fmt.Sprintf("update err : %v %s", in.Id, err)
				return err
			} else {
				out.Status = "update blog Ok"
				return nil
			}
		} else {
			out.Status = "update blog err, id and author not match"
			return nil
		}

	}
}



func (bs BlogServer) DelBlog(ctx context.Context, in *mpb.DelBlogRequest, out *mpb.DelBlogReply) (err error) {

	//过滤无效请求
	if in.Id == 0 || in.Author == "" {
		out.Status = fmt.Sprintf("del blog err : id and author can not be empty")
		return nil
	}

	mysqlDB := db.DB()

	blog := model.Blog{}
	blogsDB := mysqlDB.Find(&blog, "id = ? and author = ? ", in.Id, in.Author)

	if blogsDB.RowsAffected == 0 {
		out.Status = fmt.Sprintf("del blog err : no such blog")
		return nil
	} else {
		now := time.Now()
		blog.DeletedAt = &now
		err := mysqlDB.Save(blog).Error
		if err != nil {
			out.Status = fmt.Sprintf("del blog err : %v %s", in.Id, err)
			return err
		} else {
			out.Status = "del blog succeed "
			return nil
		}
	}
}