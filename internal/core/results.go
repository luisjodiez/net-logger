package core

import "sync"

type SafeResults struct {
	mu sync.Mutex
	results []Result
}

func (s *SafeResults) Add(r Result) {
	s.mu.Lock()
	s.results = append(s.results, r)
	s.mu.Unlock()
}

func (s *SafeResults) All() []Result {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]Result(nil), s.results...)
}
