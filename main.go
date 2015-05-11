package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	TIMETOVALIDATECLIENT = time.Second * 5  // as soon as the client connects, it should send a GET to the server to prove it's a real client
	VOTELIFELENGTH       = time.Minute * 60 // after that, it proves this again every few minutes and if we don't hear from it in X mins, we assume it's gone
)

var (
	// CODE SMELL global variables
	BASEURL = "getit.csc.kth.se:8080"
	local   bool
	_       fmt.Stringer
)

func registerFlags() {
	flag.BoolVar(&local, "local", false, "set to true if you want to run the server locally")
	flag.Parse()
	if local {
		BASEURL = "localhost:8080"
	}
}

func main() {
	registerFlags()
	// This is the map containing all of our sessions
	sessions := make(map[string]*session)
	http.HandleFunc("/", htmlHandler(local, sessions))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
