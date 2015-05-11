package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	_ fmt.Stringer
)

type session struct {
	shorturl  string
	votes     map[string]vote
	voteMutex sync.Locker
	newMsg    *sync.Cond
}

// Possible question/answer types
/*
	yes/no
	slider
	slider with "I'm bored" option
	2D scalar variables
	3,4,5 options questions
*/

// Highlight the option that the student answered last

type vote struct {
	answer  string
	expires time.Time
}

func (s *session) countvotes() map[string]int {
	counts := make(map[string]int)
	// This is cheating, remove later?
	counts["yes"] = 0
	counts["no"] = 0
	for _, v := range s.votes {
		counts[v.answer]++
	}
	return counts
}

// removes all votes whose expirytimes are past
func (s *session) purge() {
	now := time.Now()
	s.voteMutex.Lock() // if two goroutines get here at once, one does all the purging
	for id, vote := range s.votes {
		if vote.expires.Before(now) {
			fmt.Println("The vote of student ", id, " has expired")
			delete(s.votes, id)
		}
	}
	s.voteMutex.Unlock()
}
