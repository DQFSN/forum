package main

import (
	"github.com/DQFSN/forum/server/web"
	"log"
)

func main() {

	if err := web.BugFix(); err != nil {
		log.Fatal(err)
	}

}