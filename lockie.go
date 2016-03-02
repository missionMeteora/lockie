package lockie

import (
	"runtime"
	"sync/atomic"
)

// NewLockie returns a pointer to a new instance of Lockie
func NewLockie() *Lockie {
	return &Lockie{}
}

// Lockie is the primary interface for locking/unlocking
type Lockie struct {
	// Lock state
	// 0 represents an unlocked state
	// 1 represents a locked state
	lock int64
}

// Lock acquires a write-lock
func (l *Lockie) Lock() {
	// Loop until we are able to swap value of l.lock from 0 to 1
	for !atomic.CompareAndSwapInt64(&l.lock, 0, 1) {
		// Allow other go routines to utilize some CPU time
		runtime.Gosched()
	}
}

// Unlock releases a lock
func (l *Lockie) Unlock() {
	// Swaps the value of l.lock to 0
	atomic.StoreInt64(&l.lock, 0)
}
