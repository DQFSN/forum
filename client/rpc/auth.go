package rpc

import (
	"github.com/DQFSN/blog/config"
	gpb "github.com/DQFSN/blog/proto/grpc"
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	address string
)

func init() {
	gRPCConfig := config.Get().GRPC
	host := gRPCConfig.Host
	port := gRPCConfig.Port
	address = host + ":" + port
}

func LogIn(email, pwd string) string {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := gpb.NewAuthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.LogIn(ctx, &gpb.LogInRequest{Email: email, Password: pwd})
	if err != nil {
		log.Fatalf("could not login: %v", err)
	}
	log.Printf("login: %s", r.Status)

	return r.Status
}

func SignUp(email, pwd, pwd2, authCode string) string {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := gpb.NewAuthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SignUp(ctx, &gpb.SignUpRequest{Email: email, Password: pwd, PasswordCheck: pwd2, AuthCode: authCode})
	if err != nil {
		log.Fatalf("could not signup: %v", err)
	}
	log.Printf("signup: %s", r.Status)

	return r.Status
}

func ModifyUser(emailPre, emailNow, pwdPre, pwdNow string) string {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := gpb.NewAuthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ModifyUser(ctx, &gpb.ModifyUserRequest{EmailPre: emailPre, EmailNow: emailNow, PasswordPre: pwdPre, PasswordNow: pwdNow})

	if err != nil {
		log.Fatalf("could not update: %v", err)
	}
	log.Printf("update: %s", r.Status)

	return r.Status
}
