package main

import (
	"github.com/DailyActivity/rest-api-with-ginkgo/signatures"
	"log"
)

/*
Create a new MongoDB session, using a database
named "signatures". Create a new server using
that session, then begin listening for HTTP requests.
*/

func main() {
	db, err := signatures.InitDB()
	if err != nil {
		log.Println(err)
		return
	}
	server := signatures.NewServer(db)
	server.Run()
}
