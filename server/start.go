package main

import (
	"github.com/DQFSN/forum/server/srv"
	"github.com/DQFSN/forum/server/web"
)

func main() {
	go srv.BlogSrvRun()
	go srv.UserSrvRun()
	web.Run()
}