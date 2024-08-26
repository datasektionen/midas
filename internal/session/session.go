package session

import "sync"

type Manager struct {
	cookieName  string
	lock        sync.Mutex
	maxLifeTime int64
}

func NewManager(cookieName string, maxLifeTime int64) *Manager {
	return &Manager{cookieName: cookieName, maxLifeTime: maxLifeTime}
}
