package microimp

import (
	db "blog/server/db"
	"blog/server/model"
	"strings"

	mpb "blog/api/micro"
	"context"
	"fmt"
)

type Auth struct{}

func (auth Auth) LogIn(ctx context.Context, in *mpb.LogInRequest, out *mpb.LogInReply) error {

	fmt.Printf("新请求--->%v\n", in)
	mysqlDB := db.DB()
	user := model.User{}
	err := mysqlDB.Find(&user, "email = ?", in.Email).Error

	if err != nil {
		out =&mpb.LogInReply{Status: fmt.Sprintf("LogIn : %s %s", in.Email, err)}
		return  err
	}

	if user.Password == in.Password {
		out = &mpb.LogInReply{Status: "ok: " + in.Email + " " + in.Password}
		return  nil
	}

	out = &mpb.LogInReply{Status: "wrong : " + in.Email + " " + in.Password}
	return  nil

}

func (auth Auth) SignUp(ctx context.Context, in *mpb.SignUpRequest, out *mpb.SignUpReply) error {
	if strings.Contains(in.Email, "@") && len(in.Password) > 0 && in.Password == in.PasswordCheck {
		mysqlDB := db.DB()
		user := model.User{Email: in.Email, Password: in.Password}
		err := mysqlDB.Create(&user).Error
		if err != nil {
			out = &mpb.SignUpReply{Status: fmt.Sprintf("insert: %s %s", in.Email, err)}
			return  err
		}
		out = &mpb.SignUpReply{Status: "ok: " + in.Email + " " + in.Password}
		return nil
	}
	out = &mpb.SignUpReply{Status: "wrong : " + in.Email + " " + in.Password}
	return nil
}

func (auth Auth) ModifyUser(ctx context.Context, in *mpb.ModifyUserRequest, out *mpb.ModifyUserReply) error {

	if strings.Contains(in.EmailNow, "@") && in.EmailPre != in.EmailNow && in.PasswordPre != in.PasswordNow {
		mysqlDB := db.DB()
		user := model.User{Email: in.EmailPre, Password: in.PasswordPre}
		mysqlDB.First(&user)

		//更新
		user.Email = in.EmailNow
		user.Password = in.PasswordNow
		err := mysqlDB.Save(&user).Error
		if err != nil {
			out = &mpb.ModifyUserReply{Status: fmt.Sprintf("update: %s %s", in.EmailPre, err)}
			return err
		}
		out = &mpb.ModifyUserReply{Status: "ok: " + in.EmailNow + " " + in.PasswordNow}
		return nil
	}
	out = &mpb.ModifyUserReply{Status: "wrong : " + in.EmailNow + " " + in.PasswordNow}
	return nil
}
