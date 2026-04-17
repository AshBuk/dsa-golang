// Copyright 2025 Asher Buk
// SPDX-License-Identifier: MIT
// https://github.com/AshBuk/dsa-golang

// Circuit Breaker Pattern
// Protects from cascading failures when a downstream dependency starts failing.
// Without it: DB goes down -> goroutines hang on timeouts -> resources exhausted
// -> your service goes down -> callers go down. Cascade.
//
// Three states:
//   Closed:   requests pass, count consecutive failures.
//             If threshold hit -> Open.
//   Open:     requests rejected immediately (fast fail), wait timeout.
//             After timeout -> HalfOpen.
//   HalfOpen: a test request is allowed.
//             Success -> Closed. Failure -> back to Open.
//
// Used in production: Netflix Hystrix, resilience4j, sony/gobreaker.
//
// NOTE: This is the educational minimum for understanding the state machine.
// For production use sony/gobreaker — it adds sliding-window failure rate,
// per-error classification, and state-change hooks. Don't roll your own.

package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type State int

const (
	Closed State = iota
	Open
	HalfOpen
)

type CircuitBreaker struct {
	mu          sync.Mutex
	state       State
	failures    int
	maxFailures int
	timeout     time.Duration
	openedAt    time.Time
}

func New(maxFailures int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		maxFailures: maxFailures,
		timeout:     timeout,
	}
}

func (cb *CircuitBreaker) Execute(fn func() error) error {
	cb.mu.Lock()
	if cb.state == Open {
		if time.Since(cb.openedAt) > cb.timeout {
			cb.state = HalfOpen
		} else {
			cb.mu.Unlock()
			return errors.New("circuit open")
		}
	}
	cb.mu.Unlock()

	err := fn()

	cb.mu.Lock()
	defer cb.mu.Unlock()
	if err != nil {
		cb.failures++
		if cb.failures >= cb.maxFailures {
			cb.state = Open
			cb.openedAt = time.Now()
		}
		return err
	}
	cb.failures = 0
	cb.state = Closed
	return nil
}

func main() {
	cb := New(2, 500*time.Millisecond) // trip on 2 consecutive failures

	flaky := func() error { return errors.New("down") }
	healthy := func() error { return nil }

	// Two failures trip the breaker
	fmt.Println(cb.Execute(flaky))   // down  — Closed, failures=1
	fmt.Println(cb.Execute(flaky))   // down  — tripped → Open

	// Open: fast-fail, fn not called
	fmt.Println(cb.Execute(healthy)) // circuit open

	time.Sleep(600 * time.Millisecond) // timeout elapses

	// HalfOpen probe succeeds → Closed
	fmt.Println(cb.Execute(healthy)) // <nil> — recovered
}

/*
How Circuit Breaker works:

Config: maxFailures=2, timeout=500ms

STATE TRANSITIONS:
    Closed   ──(failures >= maxFailures)──▶  Open
    Open     ──(after timeout elapsed)───▶  HalfOpen
    HalfOpen ──(probe success)──────────▶  Closed
    HalfOpen ──(probe failure)──────────▶  Open

TIMELINE (matches main() demo):

STEP 1 — Closed, failures accumulate:
    call 1: flaky() fails → failures=1  state=Closed
    call 2: flaky() fails → failures=2  state=Open (tripped!)
                                         openedAt = now

    Note: counter tracks CONSECUTIVE failures.
    A single success in Closed resets it back to 0.

STEP 2 — Open, fast fail (protect the downstream):
    call 3: fn() NOT executed → returns "circuit open" immediately
            no goroutines held on timeouts, no resources burned

STEP 3 — After timeout, probe the service:
    sleep 600ms (> 500ms timeout)

    call 4: state was Open, elapsed 600ms > timeout
            → transition to HalfOpen
            → execute healthy() as the probe
            → success → state=Closed, failures=0
            fully recovered

FAILURE PATH (alternative to STEP 3):
    If the probe had failed instead:
            → failures crosses maxFailures → state back to Open
            → openedAt reset → another timeout cycle starts

Key points:
- Counts CONSECUTIVE failures, not total failures over time
- Success in Closed resets the counter to 0
- In Open, fn is NEVER executed — pure fast-fail, zero side effects
- Simplified: HalfOpen passes any incoming call. Production impls admit
  exactly one probe and fast-fail the rest (see sony/gobreaker)
- Typical production defaults: maxFailures=5, timeout=30s (tune per dependency)
- Real-world implementations: Netflix Hystrix, resilience4j, sony/gobreaker
*/
