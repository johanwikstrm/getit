package main

import (
	"code.google.com/p/go.net/websocket"
	urlshortener "code.google.com/p/google-api-go-client/urlshortener/v1"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func htmlHandler(local bool, sessions map[string]*session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/" {
			//start a session
			// TODO: make sure it's unique
			sessId := strconv.Itoa(rand.Int())
			votes := make(map[string]vote)
			// TODO: generate a short url
			longurl := "http://" + BASEURL + "/" + sessId + "/student"
			shorturl := "[short url]"
			if !local {
				if urlsSvc, err := urlshortener.New(http.DefaultClient); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				} else if url, err := urlsSvc.Url.Insert(&urlshortener.Url{LongUrl: longurl}).Do(); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				} else {
					shorturl = url.Id
				}
			}
			sessions[sessId] = &session{shorturl, votes, &sync.Mutex{}, sync.NewCond(&sync.Mutex{})}
			// redirect to '/url'
			http.Redirect(w, r, "/"+sessId+"/url", http.StatusTemporaryRedirect)
		} else if splitted := strings.Split(r.RequestURI, "/"); len(splitted) != 3 {
			// Got unknown request, lets see if the file exists in gui
			fmt.Println("404: ", splitted)
			serve404(w, r)
		} else if sess, ok := sessions[splitted[1]]; !ok {
			fmt.Println("wrong session id", splitted)
			serveWrongSessionIdPage(w, r)
		} else if sessionId, page := splitted[1], strings.Split(splitted[2], "?")[0]; page == "url" { // this is where the action is
			serveUrlPage(w, r, sessionId, sess)
		} else if page == "student" {
			serveStudentPage(w, r, sessionId, sess)
		} else if page == "teacher" {
			serveTeacherPage(w, r, sessionId, sess)
		} else if page == "datafeed" {
			websocket.Handler(DataFeed(sess)).ServeHTTP(w, r)
		} else {
			fmt.Println("Serving file gui/" + page)
			// TODO: custom 404
			http.ServeFile(w, r, "gui/"+page)
		}
	}
}
