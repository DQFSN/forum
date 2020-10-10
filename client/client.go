package main

import (
	"github.com/DQFSN/blog/client/rpc"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		usage()
		os.Exit(1)
	}

	method := os.Args[1]

	switch method {
	case "login":
		if len(os.Args) < 4 {
			usage()
			os.Exit(1)
		}
		fmt.Println(rpc.LogIn(os.Args[2], os.Args[3]))
	case "signup":
		if len(os.Args) < 6 {
			usage()
			os.Exit(1)
		}
		fmt.Println(rpc.SignUp(os.Args[2], os.Args[3], os.Args[4], os.Args[5]))
	case "publish":
		if len(os.Args) < 5 {
			usage()
			os.Exit(1)
		}
		fmt.Println(rpc.PublishBlog(os.Args[2], os.Args[3], os.Args[4]))
	case "modifyuser":
		if len(os.Args) < 6 {
			usage()
			os.Exit(1)
		}
		fmt.Println(rpc.ModifyUser(os.Args[2], os.Args[3], os.Args[4], os.Args[5]))
	default:
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Println("Welcome to blog")
	fmt.Println("Usage:")
	fmt.Println("signup email password passwordCheck authCode")
	fmt.Println("modifyuser emailPre emailNow passwordPre passwordNow")
	fmt.Println("login email password")
	fmt.Println("Enjoy")
}