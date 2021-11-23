package utils

import "sync"

// A Lamport logical clock, which can be locked/unlocked.
type Lamport struct {
	clockValue int32
	mu         sync.Mutex
}

// Increment increments the lamport clock by 1.
func (l *Lamport) Increment() {
	defer l.mu.Unlock()
	l.mu.Lock()
	l.clockValue++
}

// MaxAndIncrement sets the lamport clock to the maximum value of itself and some other clock and increments it by 1
func (l *Lamport) MaxAndIncrement(other int32) {
	defer l.mu.Unlock()
	l.mu.Lock()
	if l.clockValue < other {
		l.clockValue = other
	}

	l.clockValue++
}

// Value returns the value of the Lamport clock.
func (l *Lamport) Value() int32 {
	return l.clockValue
}

// NewLamport creates a new Lamport Clock with clockValue = 0.
func NewLamport() *Lamport {
	return &Lamport{
		clockValue: 0,
	}
}

