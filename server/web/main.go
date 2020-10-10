package main

import (
	"fmt"
	pb "github.com/DQFSN/blog/proto/micro"
	"github.com/micro/go-micro/v2"

	//pb "github.com/DQFSN/blog/proto/grpc"
	db "github.com/DQFSN/blog/server/db"
	"github.com/DQFSN/blog/server/model"
	"github.com/gin-gonic/gin"
	//"google.golang.org/grpc"
	"log"
	"strconv"
)

const (
	address = "127.0.0.1:50051"
)

func main() {

	router := gin.Default()

	//grpc 提供服务
	//conn, err := grpc.Dial(address, grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("connect err %s",err)
	//}
	//
	//defer conn.Close()
	//authClient := pb.NewAuthClient(conn)
	//blogClient := pb.NewPublishClient(conn)

	//micro 提供服务
	service := micro.NewService()
	service.Init()
	authClient := pb.NewAuthService("blog", service.Client())
	blogClient := pb.NewPublishService("blog", service.Client())

	//请求blogs
	router.GET("/blogs", func(ctx *gin.Context) {
		author := ctx.Query("author")
		resp, err := blogClient.GetBlogs(ctx, &pb.BlogsRequest{Author: author})
		if err != nil {
			log.Fatalf("Getblogs err %s", err)
		}
		ctx.IndentedJSON(200, resp.Blogs)
	})

	//注册
	router.GET("/signup", func(ctx *gin.Context) {
		email := ctx.Query("email")
		pwd := ctx.Query("pwd")
		pwdCheck := ctx.Query("pwdcheck")
		authCode := ctx.DefaultQuery("code", "000")
		resp, err := authClient.SignUp(ctx, &pb.SignUpRequest{Email: email, Password: pwd, PasswordCheck: pwdCheck, AuthCode: authCode})
		if err != nil {
			log.Fatalf("signUp err %s", err)
		}
		ctx.IndentedJSON(200, resp.Status)
	})

	//需要认证的路由
	authoriza := router.Group("/user", gin.BasicAuth(getUsers()))
	{
		authoriza.GET("/modify", func(ctx *gin.Context) {
			emailPre := ctx.Query("emailpre")
			emailNow := ctx.Query("emailnow")
			pwdPre := ctx.Query("pwdpre")
			pwdNow := ctx.Query("pwdnow")
			resp, err := authClient.ModifyUser(ctx, &pb.ModifyUserRequest{EmailPre: emailPre, EmailNow: emailNow, PasswordPre: pwdPre, PasswordNow: pwdNow})
			if err != nil {
				log.Fatalf("modigy userinfo err %s", err)
			}
			ctx.JSON(200, fmt.Sprintf("%v  %v", resp.Status, ctx.MustGet(gin.AuthUserKey)))
		})

		authoriza.GET("/modifyblog", func(ctx *gin.Context) {
			idStr := ctx.Query("id")
			id, _ := strconv.ParseInt(idStr, 10, 32)

			title := ctx.Query("title")
			content := ctx.Query("content")
			author := ctx.Query("author")

			if author == ctx.MustGet(gin.AuthUserKey) {
				resp, err := blogClient.ModifyBlog(ctx, &pb.ModifyBlogRequest{Id: int32(id), Title: title, Content: content})

				if err != nil {
					log.Fatalf("modigy blog err %s", err)
				}
				ctx.JSON(200, fmt.Sprintf("%v  %v", resp.Status, ctx.MustGet(gin.AuthUserKey)))
			} else {
				ctx.JSON(200, fmt.Sprintf("你只能修改用户名为 '%v' 的blog  ", ctx.MustGet(gin.AuthUserKey)))
			}

		})
	}

	router.Run()

}

//返回所有用户的email和pwd
func getUsers() (account map[string]string) {
	var users []*model.User
	db.DB().Where(model.User{}).Find(&users)

	account = make(map[string]string)

	for _, user := range users {
		account[user.Email] = user.Password
	}

	return account
}
