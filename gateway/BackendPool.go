package main

import (
	"crypto/sha256"
	"encoding/json"
	"net/http/httputil"
	"net/url"
	"os"
	"sync"
)

// BackendPool содержит список серверов
type BackendPool struct {
	Servers []*Server
	mu      sync.Mutex
}

// AddServer добавляет сервер в пул
func (bp *BackendPool) AddServer(url_arg string) {
	bp.mu.Lock()
	defer bp.mu.Unlock()

	proxy := httputil.NewSingleHostReverseProxy(&url.URL{Scheme: "http", Host: url_arg})
	bp.Servers = append(bp.Servers, &Server{
		URL:          url_arg,
		ActiveConns:  0,
		ReverseProxy: proxy,
	})
}

func (bp *BackendPool) LoadServers(ConfigFileName string) error {
	plan, err := os.ReadFile(ConfigFileName)
	if err != nil {
		return err
	}
	var data []string
	err = json.Unmarshal(plan, &data)

	if err != nil {
		return err
	}

	for _, v := range data {
		bp.AddServer(v)
	}
	return nil
}

// GetServer возвращает сервер на основе IP Hash + Least Connections
func (bp *BackendPool) GetServer(clientIP string) *Server {
	bp.mu.Lock()
	servers := bp.Servers // Копируем ссылки на серверы
	bp.mu.Unlock()        // Снимаем блокировку сразу

	if len(servers) == 0 {
		return nil
	}

	// Хэшируем IP клиента для выбора сервера
	hash := sha256.Sum256([]byte(clientIP))
	serverIndex := int(hash[0]) % len(servers)

	// Выбираем сервер с минимальным количеством активных подключений
	bestServer := servers[serverIndex]
	for _, server := range servers {
		server.mu.Lock() // Блокируем только доступ к ActiveConns
		if server.ActiveConns < bestServer.ActiveConns {
			bestServer = server
		}
		server.mu.Unlock()
	}

	return bestServer
}
