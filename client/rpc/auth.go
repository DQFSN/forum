package rpc

import (
	grpc2 "blog/api/grpc"
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "127.0.0.1:50051"
)

func LogIn(email, pwd string) string {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v",err)
	}
	defer conn.Close()
	c := grpc2.NewAuthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.LogIn(ctx, &grpc2.LogInRequest{Email: email, Password: pwd})
	if err!=nil {
		log.Fatalf("could not login: %v",err)
	}
	log.Printf("login: %s",r.Status)

	return r.Status
}

func SignUp(email, pwd, pwd2, authCode string) string {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v",err)
	}
	defer conn.Close()
	c := grpc2.NewAuthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SignUp(ctx, &grpc2.SignUpRequest{Email: email, Password: pwd, PasswordCheck: pwd2, AuthCode: authCode})
	if err!=nil {
		log.Fatalf("could not signup: %v",err)
	}
	log.Printf("signup: %s",r.Status)

	return r.Status
}

func ModifyUser(emailPre, emailNow, pwdPre, pwdNow string) string {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v",err)
	}

	defer conn.Close()
	c := grpc2.NewAuthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ModifyUser(ctx, &grpc2.ModifyUserRequest{EmailPre: emailPre, EmailNow: emailNow, PasswordPre: pwdPre, PasswordNow: pwdNow})

	if err != nil {
		log.Fatalf("could not update: %v",err)
	}
	log.Printf("update: %s",r.Status)

	return r.Status
}
