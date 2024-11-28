package main

import (
	"net/http/httputil"
	"sync"
)

// Server представляет бэкенд сервер
type Server struct {
	URL          string
	ActiveConns  int
	mu           sync.Mutex
	ReverseProxy *httputil.ReverseProxy
}

// IncrementConnections увеличивает количество активных соединений на сервере
func (s *Server) IncrementConnections() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ActiveConns++
}

// DecrementConnections уменьшает количество активных соединений на сервере
func (s *Server) DecrementConnections() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.ActiveConns > 0 {
		s.ActiveConns--
	}
}
