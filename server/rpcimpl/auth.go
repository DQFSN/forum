package rpcimpl

import (
	"fmt"
	gpb "github.com/DQFSN/blog/proto/grpc"
	db "github.com/DQFSN/blog/server/db"
	"github.com/DQFSN/blog/server/model"
	"github.com/DQFSN/blog/server/util"
	"golang.org/x/net/context"
	"strings"
)

type Auth struct{}

//登录
func (s *Auth) LogIn(ctx context.Context, in *gpb.LogInRequest) (*gpb.LogInReply, error) {

	fmt.Printf("新请求--->%v\n",in)
	mysqlDB := db.DB()
	user := model.User{}
	err := mysqlDB.Find(&user,"email = ?", in.Email).Error

	if err != nil {
		return &gpb.LogInReply{Status: fmt.Sprintf( "LogIn : %s %s", in.Email, err)}, err
	}

	if util.ComparePasswords([]byte(in.Email),[]byte(in.Password),user.Password)  {
		return &gpb.LogInReply{Status: "ok: " + in.Email + " " + in.Password}, nil
	}
	return &gpb.LogInReply{Status: "wrong : " + in.Email + " " + in.Password}, nil
}

//注册
func (s *Auth) SignUp(ctx context.Context, in *gpb.SignUpRequest) (*gpb.SignUpReply, error) {

	if strings.Contains(in.Email, "@")  && len(in.Password) > 0 && in.Password == in.PasswordCheck {
		mysqlDB := db.DB()

		hashPassword := util.HashAndSalt([]byte(in.Email), []byte(in.Password))
		user := model.User{Email: in.Email, Password: hashPassword}
		err := mysqlDB.Create(&user).Error
		if err != nil {
			return &gpb.SignUpReply{Status: fmt.Sprintf("insert: %s %s", in.Email, err)}, err
		}
		return &gpb.SignUpReply{Status: "ok: " + in.Email + " " + in.Password}, nil
	}
	return &gpb.SignUpReply{Status: "wrong : " + in.Email + " " + in.Password}, nil
}

//修改邮箱密码
func (s *Auth) ModifyUser(ctx context.Context, in *gpb.ModifyUserRequest) (*gpb.ModifyUserReply, error) {

	if strings.Contains(in.EmailNow, "@")  && in.EmailPre != in.EmailNow && in.PasswordPre != in.PasswordNow {
		mysqlDB := db.DB()
		user := model.User{Email: in.EmailPre, Password: in.PasswordPre}
		mysqlDB.First(&user)

		//更新
		user.Email = in.EmailNow
		user.Password = util.HashAndSalt([]byte(in.EmailNow), []byte(in.PasswordNow))
		err := mysqlDB.Save(&user).Error
		if err != nil {
			return &gpb.ModifyUserReply{Status: fmt.Sprintf("update: %s %s", in.EmailPre, err)}, err
		}
		return &gpb.ModifyUserReply{Status: "ok: " + in.EmailNow + " " + in.PasswordNow}, nil
	}
	return &gpb.ModifyUserReply{Status: "wrong : " + in.EmailNow + " " + in.PasswordNow}, nil
}