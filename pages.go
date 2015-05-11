package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"
)

var (
	_ fmt.Stringer
)

type UrlPage struct {
	StudentUrl string
	TeacherUrl string
}

type StudentPage struct {
	StudentId  string
	LastAnswer string
}

type TeacherPage struct {
	SuperShortUrl string
}

type WrongSessionIdPage struct {
	BaseUrl string
}

type NotFound404Page struct {
	BaseUrl string
}

func executePageTemplate(templateFilePath string, templateStructPtr interface{}, w http.ResponseWriter) {
	if tmpl, err := template.ParseFiles(templateFilePath); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := tmpl.Execute(w, templateStructPtr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func serveUrlPage(w http.ResponseWriter, r *http.Request, sessId string, sess *session) {
	urlpage := UrlPage{}
	urlpage.StudentUrl = sess.shorturl
	urlpage.TeacherUrl = BASEURL + "/" + sessId + "/teacher"
	executePageTemplate("gui/url.go.html", &urlpage, w)
}

// if no query parameters, this is a new student, else, this is an old student
func serveStudentPage(w http.ResponseWriter, r *http.Request, sessId string, sess *session) {
	timeUntilExpiry := VOTELIFELENGTH
	tmpl, err := template.ParseFiles("gui/student.go.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Find out student id
	//studentid := r.URL.Query().Get("studentid")
	// The host name should be unique for all clients
	studentid := strings.Split(r.RemoteAddr, ":")[0]
	fmt.Println("Host", studentid, "sent http request")
	var vt vote
	var exists bool
	if vt, exists = sess.votes[studentid]; !exists { // new student, let's be wary and give them shorter time to prove they're a real client
		timeUntilExpiry = TIMETOVALIDATECLIENT
		fmt.Println("Host", studentid, "is new, setting time to expiry to", timeUntilExpiry)
		vt = vote{}
	}
	// find out what the student voted for
	vt.answer = "yes" // yes by default
	if a := r.URL.Query().Get("answer"); a != "" {
		vt.answer = a
	}
	// refresh the vote
	vt.expires = time.Now().Add(timeUntilExpiry)
	// store the vote
	// TODO threadsafe
	sess.votes[studentid] = vt
	fmt.Println("Student:", studentid, "votes", sess.votes[studentid].answer)
	// purge expired votes
	sess.purge()
	// Update all websockets
	sess.newMsg.Broadcast()
	// Prepare the page
	studentpage := StudentPage{}
	studentpage.StudentId = studentid
	studentpage.LastAnswer = sess.votes[studentid].answer
	tmpl.Execute(w, &studentpage)
}

func serveTeacherPage(w http.ResponseWriter, r *http.Request, sessId string, sess *session) {
	teacherpage := TeacherPage{}
	teacherpage.SuperShortUrl = strings.TrimPrefix(sess.shorturl, "http://")
	executePageTemplate("gui/teacher.go.html", &teacherpage, w)
}

func serveWrongSessionIdPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("gui/wrongsessionid.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	wsidPage := WrongSessionIdPage{}
	wsidPage.BaseUrl = BASEURL
	if err := tmpl.Execute(w, &wsidPage); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func serve404(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("gui/404.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	page404 := NotFound404Page{}
	page404.BaseUrl = BASEURL
	tmpl.Execute(w, &page404)
}
