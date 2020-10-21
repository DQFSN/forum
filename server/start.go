package main

import (
	"github.com/DQFSN/forum/server/srv"
	"github.com/DQFSN/forum/server/web"
	"log"
)

func main() {

	if err := web.BugFix(); err != nil {
		log.Fatal(err)
	}

	go srv.BlogSrvRun()
	go srv.UserSrvRun()
	web.Run()


}