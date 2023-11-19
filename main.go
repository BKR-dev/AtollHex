package main

import (
	"fmt"

	"github.com/BKR-dev/AtollHex/server"
	"github.com/BKR-dev/AtollHex/util"
)

func main() {

	util.PopluateDB()
	fmt.Println("dummy data in db")
	srv := server.NewServer()
	fmt.Println("started server")
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Printf("server died bros, rip: %s\n", err)
	}
}
