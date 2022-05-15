package main

import (
	"github.com/toluwase1/playstore/rest"
	"log"
)

func main() {
	log.Println("Main log....")
	log.Fatal(rest.RunAPI("127.0.0.1:8000"))
}
