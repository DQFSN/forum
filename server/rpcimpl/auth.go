package rpcimpl

import (
	pb "blog/api"
	db "blog/server/db"
	"blog/server/model"
	"fmt"
	"golang.org/x/net/context"
	"strings"
)

type Auth struct{}

//登录
func (s *Auth) LogIn(ctx context.Context, in *pb.LogInRequest) (*pb.LogInReply, error) {

	fmt.Printf("新请求--->%v\n",in)
	mysqlDB := db.DB()
	user := model.User{}
	err := mysqlDB.Find(&user,"email = ?", in.Email).Error

	if err != nil {
		return &pb.LogInReply{Status:fmt.Sprintf( "LogIn : %s %s", in.Email, err)}, err
	}

	if user.Password ==  in.Password  {
		return &pb.LogInReply{Status: "ok: " + in.Email + " " + in.Password}, nil
	}
	return &pb.LogInReply{Status: "wrong : " + in.Email + " " + in.Password}, nil
}

//注册
func (s *Auth) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpReply, error) {

	if strings.Contains(in.Email, "@")  && len(in.Password) > 0 && in.Password == in.PasswordCheck {
		mysqlDB := db.DB()
		user := model.User{Email: in.Email, Password: in.Password}
		err := mysqlDB.Create(&user).Error
		if err != nil {
			return &pb.SignUpReply{Status: fmt.Sprintf("insert: %s %s", in.Email, err)}, err
		}
		return &pb.SignUpReply{Status: "ok: " + in.Email + " " + in.Password}, nil
	}
	return &pb.SignUpReply{Status: "wrong : " + in.Email + " " + in.Password}, nil
}

//修改邮箱密码
func (s *Auth) ModifyUser(ctx context.Context, in *pb.ModifyUserRequest) (*pb.ModifyUserReply, error) {

	if strings.Contains(in.EmailNow, "@")  && in.EmailPre != in.EmailNow && in.PasswordPre != in.PasswordNow {
		mysqlDB := db.DB()
		user := model.User{Email: in.EmailPre, Password: in.PasswordPre}
		mysqlDB.First(&user)

		//更新
		user.Email = in.EmailNow
		user.Password = in.PasswordNow
		err := mysqlDB.Save(&user).Error
		if err != nil {
			return &pb.ModifyUserReply{Status: fmt.Sprintf("update: %s %s", in.EmailPre, err)}, err
		}
		return &pb.ModifyUserReply{Status: "ok: " + in.EmailNow + " " + in.PasswordNow}, nil
	}
	return &pb.ModifyUserReply{Status: "wrong : " + in.EmailNow + " " + in.PasswordNow}, nil
}