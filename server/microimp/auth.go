package microimp

import (
	db "github.com/DQFSN/forum/server/db"
	"github.com/DQFSN/forum/server/model"
	"github.com/DQFSN/forum/server/util"
	"strings"
	"time"

	mpb "github.com/DQFSN/forum/proto/micro"
	"context"
	"fmt"
)

type AuthHandler struct{}

func (auth AuthHandler) LogIn(ctx context.Context, in *mpb.LogInRequest, out *mpb.LogInReply) error {

	fmt.Printf("新请求--->%v\n", in)
	mysqlDB := db.DB()
	user := model.User{}
	err := mysqlDB.Find(&user, "email = ? and deleted_at IS NULL ", in.Email).Error

	if err != nil {
		out.Status = fmt.Sprintf("LogIn : %s %s", in.Email, err)
		return  err
	}


	if util.ComparePasswords([]byte(in.Email),[]byte(in.Password),user.Password) {
		out.Status =  "ok: " + in.Email + " " + in.Password
		return  nil
	}

	out.Status = "wrong : " + in.Email + " " + in.Password
	return  nil

}

func (auth AuthHandler) SignUp(ctx context.Context, in *mpb.SignUpRequest, out *mpb.SignUpReply) error {
	if strings.Contains(in.Email, "@") && len(in.Password) > 0 && in.Password == in.PasswordCheck {
		mysqlDB := db.DB()

		hashPassword := util.HashAndSalt([]byte(in.Email),[]byte(in.Password))
		user := model.User{Email: in.Email, Password: hashPassword}
		err := mysqlDB.Create(&user).Error
		if err != nil {
			out.Status = fmt.Sprintf("insert: %s %s", in.Email, err)
			return  err
		}
		out.Status = "ok: " + in.Email + " " + in.Password
		return nil
	}
	out.Status = "wrong : " + in.Email + " " + in.Password
	return nil
}

func (auth AuthHandler) ModifyUser(ctx context.Context, in *mpb.ModifyUserRequest, out *mpb.ModifyUserReply) error {

	if strings.Contains(in.EmailNow, "@") && len(in.EmailPre) > 0 &&in.EmailPre != in.EmailNow && in.PasswordPre != in.PasswordNow {
		mysqlDB := db.DB()
		user := model.User{Email: in.EmailPre, Password: in.PasswordPre}

		result := mysqlDB.Where("email = ? and password = ? and deleted_at IS NULL ",user.Email, user.Password).Find(&user)

		if result.RowsAffected == 0 {
			out.Status = fmt.Sprintf("modify userinfo failed, email or passwod wrong ")
			return nil
		}

		//更新
		user.Email = in.EmailNow
		user.Password = util.HashAndSalt([]byte(in.EmailNow), []byte(in.PasswordNow))

		err := mysqlDB.Save(&user).Error

		if err != nil {
			out.Status = fmt.Sprintf("update: %s %s", in.EmailPre, err)
			return err
		}
		out.Status = fmt.Sprintf("update user info ok ")
		return nil
	}
	out.Status = fmt.Sprintf("modify user worng , eg: email format wrong, new passwod(email) and pre password(email) is same")
	return nil
}


func (auth AuthHandler) DelUser(ctx context.Context, in *mpb.DelUserRequest, out *mpb.DelUserReply) error {
	 if strings.Contains(in.Email,"@") && len(in.Password) > 0 {
	 	mysqlDB := db.DB()
	 	user := model.User{}
	 	result := mysqlDB.Where(&user, "email = ? and passwoed = ?",in.Email, in.Password).Find(&user)

		 if result.RowsAffected == 0 {
			 out.Status = fmt.Sprintf("del user failed , email or passwod wrong")
			 return nil
		 }else {
		 	now := time.Now()
		 	user.DeletedAt = &now

		 	err := mysqlDB.Save(&user).Error

			 if err != nil  {
				 out.Status = fmt.Sprintf("del user failed %v ", err)
				 return nil
			 }

		 	out.Status = fmt.Sprintf("del user succeed ")
		 	return nil
		 }

	 }else {
	 	out.Status = fmt.Sprintf("del user failed , email format wrong or passwod is empty")
	 	return nil
	 }
}