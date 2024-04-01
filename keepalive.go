package ws

import (
	"sync"
	"time"
)

type keepAliveResponse struct {
	lastResponse time.Time
	mu           sync.Mutex
}

func (k *keepAliveResponse) update() {
	k.mu.Lock()
	k.lastResponse = time.Now()
	k.mu.Unlock()
}

func (k *keepAliveResponse) isAlive() bool {
	k.mu.Lock()
	defer k.mu.Unlock()
	return time.Since(k.lastResponse) < time.Second*5
}
