package microimp

import (
	"context"
	mpb "github.com/DQFSN/blog/proto/micro"
	"testing"
)

func Test_SignUp(t *testing.T) {
	var auth AuthHandler

	requests := []mpb.SignUpRequest{
		{AuthCode: "", Email: "7.com", Password: "d", PasswordCheck: "dd"},
		{AuthCode: "d", Email: "7@q.com", Password: "d", PasswordCheck: ""},
		{AuthCode: "d", Email: "7@q.com", Password: "d", PasswordCheck: "d"},
		{AuthCode: "sx", Email: "7@q.com", Password: "", PasswordCheck: "d"},
		{AuthCode: "1", Email: "6@q.com", Password: "dd", PasswordCheck: "d"},
	}

	reply := mpb.SignUpReply{}

	for _, in := range requests {
		err := auth.SignUp(context.TODO(), &in, &reply)
		if reply.Status == "" {
			t.Errorf("测试失败%v", err)
		} else {
			t.Logf("测试成功%v\t", reply.Status)
		}
	}
}


func Test_LogIn(t *testing.T) {
	var auth AuthHandler

	requests := []mpb.LogInRequest{
		{ Email: "7.com", Password: "d"},
		{ Email: "7@q.com", Password: "d"},
		{ Email: "", Password: "d"},
		{ Email: "7@q.com.com", Password: ""},
		{ Email: "", Password: ""},
	}

	reply := mpb.LogInReply{}

	for _, in := range requests {
		err := auth.LogIn(context.TODO(), &in, &reply)
		if reply.Status == "" {
			t.Errorf("测试失败%v", err)
		} else {
			t.Logf("测试成功%v\t", reply.Status)
		}
	}
}

func Test_ModifyUser(t *testing.T) {
	var auth AuthHandler

	requests := []mpb.ModifyUserRequest{
		{EmailNow: "7@q.com",EmailPre: "7@q.com",PasswordNow: "d",PasswordPre: "d"},
		{EmailNow: "7@q.com",EmailPre: "6@q.com",PasswordNow: "d",PasswordPre: "dd"},
		{EmailNow: "7@q.com",EmailPre: "6@q.com",PasswordNow: "d",PasswordPre: "s"},
		{EmailNow: "7@q.com",EmailPre: "7@q.com",PasswordNow: "dd",PasswordPre: "d"},
		{EmailNow: "7@q.com",EmailPre: "",PasswordNow: "dd",PasswordPre: "d"},
		{EmailNow: "5@q.com",EmailPre: "7@q.com",PasswordNow: "dd",PasswordPre: "d"},
		{EmailNow: "7q.com",EmailPre: "7@q.com",PasswordNow: "dd",PasswordPre: "d"},
		{EmailNow: "7@q.com",EmailPre: "7q.com",PasswordNow: "dd",PasswordPre: "d"},
		{EmailNow: "",EmailPre: "7q.com",PasswordNow: "dd",PasswordPre: "d"},
		{EmailNow: "@"},
	}

	reply := mpb.ModifyUserReply{}

	for _, in := range requests {
		err := auth.ModifyUser(context.TODO(), &in, &reply)
		if reply.Status == "" {
			t.Errorf("测试失败%v", err)
		} else {
			t.Logf("测试成功%v\t", reply.Status)
		}
	}
}


func Test_DelUser(t *testing.T) {
	var auth AuthHandler

	requests := []mpb.DelUserRequest{
		{ Email: "7.com", Password: "d"},
		{ Email: "7@q.com", Password: "d"},
		{ Email: "", Password: "d"},
		{ Email: "7@q.com.com", Password: ""},
		{ Email: "", Password: ""},
	}

	reply := mpb.DelUserReply{}

	for _, in := range requests {
		err := auth.DelUser(context.TODO(), &in, &reply)
		if reply.Status == "" {
			t.Errorf("测试失败%v", err)
		} else {
			t.Logf("测试成功%v\t", reply.Status)
		}
	}
}