// Copyright 2025 Asher Buk
// SPDX-License-Identifier: MIT
// https://github.com/AshBuk/dsa-golang

// Thread-Safe In-Memory Cache with TTL
// TTL (Time To Live) — each entry expires after a given duration.
// On Get, if expired — return false but don't delete (passive check).
// This allows RLock for concurrent reads (no map mutation in Get).
//
// Trade-off: expired entries leak memory until overwritten or cleaned.
// Typical fix: background goroutine with time.Ticker + full Lock to sweep
// expired entries periodically.
//
// NOTE: This is the educational minimum. For production use
// patrickmn/go-cache (simple TTL + sweeper) or dgraph-io/ristretto
// (high-QPS, cost-based eviction). Don't roll your own.

package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu    sync.RWMutex
	items map[string]entry
}

type entry struct {
	value     any
	expiresAt time.Time
}

func New() *Cache {
	return &Cache{items: make(map[string]entry)}
}

func (c *Cache) Set(key string, value any, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = entry{value, time.Now().Add(ttl)}
}

func (c *Cache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	e, ok := c.items[key]
	if !ok || time.Now().After(e.expiresAt) {
		return nil, false
	}
	return e.value, true
}

func main() {
	c := New()
	c.Set("k", "hello", 500*time.Millisecond) // TTL=500ms

	fmt.Println(c.Get("k"))    // hello true  — fresh hit
	fmt.Println(c.Get("miss")) // <nil> false — key not set

	time.Sleep(600 * time.Millisecond) // TTL elapses

	fmt.Println(c.Get("k")) // <nil> false — expired

	c.Set("k", "world", 500*time.Millisecond) // renew with fresh TTL
	fmt.Println(c.Get("k"))                   // world true — renewed
}

/*
How TTL Cache works:

Config: TTL=500ms per entry

INVARIANTS:
- Entry is valid iff time.Now() is before expiresAt
- Get is read-only on the map (RLock) — never mutates
- Set is the only path that mutates the map (full Lock)

TIMELINE (matches main() demo):

t=0ms   Set("k", "hello", 500ms)
        → items["k"] = {value:"hello", expiresAt: 500ms}

t=0ms   Get("k")    → found, now<expiresAt → return "hello", true
        Get("miss") → not in map           → return nil, false

t=600ms sleep 600ms (> TTL)

        Get("k")    → found, now>expiresAt → return nil, false
                      stale entry stays in the map (passive check)

t=600ms Set("k", "world", 500ms)
        → items["k"] overwritten, expiresAt = 1100ms
        Get("k")    → found, now<expiresAt → return "world", true

Key points:
- RWMutex: many concurrent Readers, single Writer — ideal for read-heavy caches
- Get takes RLock → N concurrent Gets scale linearly (no contention)
- Passive expiration: Get never mutates the map, so RLock is enough.
  If Get deleted on miss, we would need full Lock and lose read parallelism.
- Trade-off: expired entries leak memory until re-Set or actively swept
- To bound memory: background goroutine + time.Ticker iterates items,
  taking a full Lock to delete expired ones
- Real-world implementations: patrickmn/go-cache (simple TTL + sweeper),
  dgraph-io/ristretto (high-QPS, cost-based eviction)
*/
