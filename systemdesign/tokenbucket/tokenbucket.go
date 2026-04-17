// Copyright 2025 Asher Buk
// SPDX-License-Identifier: MIT
// https://github.com/AshBuk/dsa-golang

// Token Bucket Rate Limiter
// There is a bucket with fixed capacity of tokens.
// Each request consumes one token. Tokens refill at a constant rate.
// If bucket is empty — request is rejected.
//
// Why token bucket: allows short bursts (up to capacity) while
// enforcing average rate. Simple to implement, used in production
// (nginx, AWS API Gateway, Go's stdlib golang.org/x/time/rate).
//
// Core idea: on each Allow() call, calculate how many tokens have
// accumulated since the last call using elapsed time, then try to
// consume one.
//
// NOTE: This is the educational minimum for understanding the algorithm.
// For production use golang.org/x/time/rate — it adds Wait/Reserve/AllowN,
// context cancellation, and is battle-tested. Don't roll your own.

package main

import (
	"fmt"
	"sync"
	"time"
)

type Limiter struct {
	mu         sync.Mutex
	tokens     float64
	capacity   float64
	rate       float64
	lastRefill time.Time
}

func New(rate, capacity float64) *Limiter {
	return &Limiter{
		tokens:     capacity,
		capacity:   capacity,
		rate:       rate,
		lastRefill: time.Now(),
	}
}

func (l *Limiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(l.lastRefill).Seconds()
	l.tokens += elapsed * l.rate
	if l.tokens > l.capacity {
		l.tokens = l.capacity
	}
	l.lastRefill = now

	if l.tokens < 1 {
		return false
	}
	l.tokens--
	return true
}

func main() {
	l := New(1, 2) // rate=1 token/sec, capacity=2 (burst of 2)

	// Burst drains the bucket
	fmt.Println(l.Allow()) // true  — 1/2 consumed
	fmt.Println(l.Allow()) // true  — 2/2 consumed, drained
	fmt.Println(l.Allow()) // false — empty, rejected

	time.Sleep(1 * time.Second) // +1 token refilled

	fmt.Println(l.Allow()) // true  — recovered
}

/*
How Token Bucket works:

Config: rate=1 token/sec, capacity=2

INVARIANTS:
- Bucket never exceeds capacity (cap on bursts)
- Tokens accrue continuously at `rate` per second (fractional allowed)
- Allow() consumes exactly 1 token or returns false

TIMELINE (matches main() demo):

t=0.00s  bucket starts full: tokens=2.00
    call 1: tokens=2.00 → consume → tokens=1.00   allow=true
    call 2: tokens=1.00 → consume → tokens=0.00   allow=true  (drained)
    call 3: tokens=0.00 → <1, reject              allow=false (empty)

t=1.00s  sleep 1 second
    elapsed = 1.0s × 1 token/s = +1.00 token refilled
    tokens = 0.00 + 1.00 = 1.00

    call 4: tokens=1.00 → consume → tokens=0.00   allow=true  (recovered)

REFILL MATH:
    elapsed = now - lastRefill
    tokens += elapsed * rate
    tokens = min(tokens, capacity)   ← cap prevents unbounded accrual
    lastRefill = now

Key points:
- Lazy refill: no background goroutine — compute tokens only on Allow()
- `float64` for tokens: lets small elapsed intervals accrue fractional tokens
- Capacity cap is what enforces the BURST limit
- Refill rate is what enforces the AVERAGE rate over time
- Tune: rate = steady-state QPS, capacity = max burst you tolerate
- Real-world: nginx limit_req, AWS API Gateway, golang.org/x/time/rate
*/
