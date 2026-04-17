// Copyright 2025 Asher Buk
// SPDX-License-Identifier: MIT
// https://github.com/AshBuk/dsa-golang

// Thread-Safe In-Memory Cache with TTL and LRU eviction
//
// TTL (Time To Live) — each entry has an expiration time.
// On Get, if expired — delete and return false (lazy expiration).
//
// LRU (Least Recently Used) — when cache is full, evict the entry
// that hasn't been accessed the longest. Doubly-linked list gives O(1)
// move-to-front on access and O(1) eviction from tail.
//   Front of list = MRU (most recently used, newest access)
//   Back of list  = LRU (least recently used, oldest access)
//
// Data structures: map for O(1) lookup + doubly-linked list for access order.
// container/list is Go's stdlib doubly-linked list — use it, don't reimplement.
// sync.Mutex because Get mutates the list (moves element to front).
//
// Why combine LRU + TTL: pure LRU doesn't kill stale data that stays warm;
// pure TTL doesn't bound memory under hot load. Together they cover both.
// Used by Memcached and Redis (maxmemory-policy=allkeys-lru + EXPIRE).
//
// NOTE: This is the educational minimum. For production use
// hashicorp/golang-lru (thread-safe LRU + optional TTL, battle-tested).
// Don't roll your own.

package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type entry struct {
	key       string
	value     any
	expiresAt time.Time
}

type Cache struct {
	mu       sync.Mutex
	items    map[string]*list.Element
	order    *list.List
	capacity int
	ttl      time.Duration
}

func New(capacity int, ttl time.Duration) *Cache {
	return &Cache{
		items:    make(map[string]*list.Element, capacity),
		order:    list.New(),
		capacity: capacity,
		ttl:      ttl,
	}
}

func (c *Cache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if el, ok := c.items[key]; ok {
		c.order.MoveToFront(el)
		el.Value.(*entry).value = value
		el.Value.(*entry).expiresAt = time.Now().Add(c.ttl)
		return
	}
	if c.order.Len() >= c.capacity {
		back := c.order.Back()
		c.order.Remove(back)
		delete(c.items, back.Value.(*entry).key)
	}
	el := c.order.PushFront(&entry{key, value, time.Now().Add(c.ttl)})
	c.items[key] = el
}

func (c *Cache) Get(key string) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	el, ok := c.items[key]
	if !ok {
		return nil, false
	}
	if time.Now().After(el.Value.(*entry).expiresAt) {
		c.order.Remove(el)
		delete(c.items, el.Value.(*entry).key)
		return nil, false
	}
	c.order.MoveToFront(el)
	return el.Value.(*entry).value, true
}

func main() {
	c := New(2, 500*time.Millisecond) // capacity=2, TTL=500ms

	c.Set("a", 1)
	c.Set("b", 2)
	fmt.Println(c.Get("a")) // 1 true       — hit, "a" moves to front (now "b" is LRU)

	c.Set("c", 3)           // capacity full → evict "b" (least recently used)
	fmt.Println(c.Get("b")) // <nil> false  — evicted by LRU
	fmt.Println(c.Get("a")) // 1 true       — still cached

	time.Sleep(600 * time.Millisecond) // TTL elapses

	fmt.Println(c.Get("a")) // <nil> false  — expired by TTL
}

/*
How LRU + TTL Cache works:

Config: capacity=2, TTL=500ms

INVARIANTS:
- map gives O(1) lookup by key
- doubly-linked list keeps entries in access order (front = newest)
- items[key] points to a list.Element, enabling O(1):
    * find entry by key                  → items[key]
    * mark as most-recently-used         → order.MoveToFront(el)
    * evict least-recently-used          → order.Back() + Remove + delete(items)
- TTL is stored on each entry, checked lazily on Get

TIMELINE (matches main() demo):

t=0ms   Set("a", 1)        items: {a}     order: [a]            a.exp=500
t=0ms   Set("b", 2)        items: {a,b}   order: [b, a]         b.exp=500
                                                 ↑MRU↑LRU

t=0ms   Get("a") → 1, true → MoveToFront(a)
                           order: [a, b]
                                  ↑MRU↑LRU
                           (now "b" is LRU because "a" was just touched)

t=0ms   Set("c", 3)        len == capacity → evict back = "b"
                           delete items["b"], order.Remove(b)
                           PushFront(c)
                           items: {a,c}   order: [c, a]         c.exp=500
                                                 ↑MRU↑LRU

t=0ms   Get("b") → nil, false  ("b" was evicted by LRU policy)
t=0ms   Get("a") → 1, true     (still fresh, MoveToFront)
                               order: [a, c]

t=600ms sleep 600ms

        Get("a") → entry found, but now > expiresAt (500ms) → expired
                   order.Remove(el) + delete(items, "a")  ← lazy expiration
                   return nil, false
                           items: {c}     order: [c]
                           ("c" is also expired but lingers until accessed)

WHY TTL IS ORTHOGONAL TO LRU:
- LRU machinery (map + list) doesn't know about expiresAt
- TTL is a single field on entry + one if-check in Get
- Both removal paths reuse the SAME two lines:
        c.order.Remove(el)
        delete(c.items, key)
  No coupling, no shared state. Two concepts composed cleanly.

Key points:
- Two containers stay in sync: map[key] → list.Element → entry(key,value,exp)
- Storing key INSIDE entry is required: eviction walks the list, not the map,
  and needs to know which map key to delete
- sync.Mutex (not RWMutex): Get MUTATES the list via MoveToFront,
  so RLock would be unsafe — direct contrast to the passive-check ttlcache
- Lazy expiration: Get deletes on miss, no background sweeper needed for
  hot keys; cold expired entries linger until LRU evicts them
- TTL is global (set in New), not per-entry — simpler API, matches Memcached
- container/list is Go's stdlib doubly-linked list — use it, don't reimplement
- Edge case: capacity=0 panics on Set (Back() returns nil) — validate in prod
- Real-world: hashicorp/golang-lru, Memcached, Redis (allkeys-lru + EXPIRE)
*/
