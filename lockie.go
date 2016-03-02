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
	lock int64
}

// Lock acquires a write-lock
func (l *Lockie) Lock() {
	// Acquire a write-lock
	for !atomic.CompareAndSwapInt64(&l.lock, 0, 1) {
		runtime.Gosched()
	}
}

// Unlock releases a write-lock
func (l *Lockie) Unlock() {
	atomic.StoreInt64(&l.lock, 0)
}

// NewRWLockie returns a pointer to a new instance of RWLockie
func NewRWLockie() *RWLockie {
	return &RWLockie{}
}

// RWLockie is no different than Lockie
// This is used to match a standard RW mutex interface
type RWLockie struct {
	Lockie
}

// RLock acquires a read-lock
func (l *RWLockie) RLock() {
	l.Lock()
}

// RUnlock releases a read-lock
func (l *RWLockie) RUnlock() {
	l.Unlock()
}
