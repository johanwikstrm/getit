package main

import (
	"sync"
	"testing"
	"time"
)

func TestPurge(t *testing.T) {
	s := &session{votes: map[string]vote{}, voteMutex: &sync.Mutex{}}
	s.votes["olle"] = vote{"no", time.Now().Add(-time.Second)}
	s.votes["pelle"] = vote{"yes", time.Now().Add(time.Minute)}
	if len(s.votes) != 2 {
		t.Fatal("Wrong len of map")
	}
	s.purge()
	if len(s.votes) != 1 {
		t.Fatal("Wrong len of map")
	}
	// add one million students with expired dates
	for i := 0; i < 100000; i++ {
		s.votes["olle"] = vote{"no", time.Now().Add(-time.Second)}
	}
	// run four purges
	wg := &sync.WaitGroup{}
	wg.Add(4)
	for i := 0; i < 4; i++ {

		go func() {
			s.purge()
			wg.Done()
		}()

	}
	wg.Wait()
}
