package main

import "local/server"

func main() {
	srv := server.NewServer()
	srv.ListenAndServe()
}
