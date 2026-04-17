# System Design Patterns for Go — Roadmap

Common **"implement X"** patterns asked at backend/infra Go interviews.
Each is implemented here as a single educational file: minimal correct code, walkthrough with step-by-step timeline, and a pointer to the canonical production library — **do not roll your own in real systems**.

## Tier 1

- ✅ [**TTL Cache**](ttlcache/ttlcache.go) — `map + RWMutex`, lazy passive expiration
- ✅ [**LRU Cache (with TTL)**](lrucache/lrucache.go) — `map + container/list` + lazy expiration
- ✅ [**Token Bucket Rate Limiter**](tokenbucket/tokenbucket.go) — lazy refill, allows bursts
- ✅ [**Circuit Breaker**](circuitbreaker/circuitbreaker.go) — three-state machine (Closed / Open / HalfOpen)
- ⬜ **Bounded Concurrency** — Worker Pool / Semaphore / `errgroup` (one folder, related primitives)
- ⬜ **Bloom Filter** — bit array + N hash functions, probabilistic membership
- ⬜ **Consistent Hashing** — hash ring with virtual nodes, sharding

## Tier 2

- ⬜ **Connection Pool** — generic pool with health checks
- ⬜ **Pub/Sub (in-process)** — broker on channels
- ⬜ **Leaky Bucket** — strict pacing (paired teaching file with token bucket)
- ⬜ **Distributed Lock** — Redis `SETNX` / Redlock pattern
- ⬜ **Retry with Exponential Backoff + Jitter** — failure recovery

## Tier 3

- ⬜ **Job Scheduler** — bounded workers + retry + queue (follow-up to bounded concurrency)
- ⬜ **URL Shortener** — hashing, collision handling
- ⬜ **Middleware Chain** — composable interceptors (the "API gateway" core)
- ⬜ **In-Memory KV with Snapshot** — `Set/Get/Delete` + persist/restore (lighter than full WAL)
