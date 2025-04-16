package locker

import (
	"sync"
)

var mtx = &sync.Mutex{}

func Lock() func() {
	mtx.Lock()
	return func() {
		mtx.Unlock()
	}
}
