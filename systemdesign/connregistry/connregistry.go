// Copyright 2025 Asher Buk
// SPDX-License-Identifier: MIT
// https://github.com/AshBuk/dsa-golang

// Connection Registry Pattern
// A server holds many active client connections (chat, WebSocket gateway,
// IoT, game lobby) and must reach any of them by ID, or all at once.
//
// Write-once on Connect, delete-once on Disconnect, read-many on Send.
// For settled keys sync.Map.Load is an atomic read with no lock at all,
// which beats RWMutex.RLock under high contention.
//
// NOTE: Educational minimum. Real gateways add per-conn write goroutines
// with bounded queues, heartbeats, graceful shutdown. Use nhooyr.io/websocket
// or gorilla/websocket. Don't roll your own.

package main

import (
	"fmt"
	"net"
	"sync"
)

type Registry struct {
	clients sync.Map // map[string]net.Conn
}

func New() *Registry {
	return &Registry{}
}

func (r *Registry) Connect(id string, conn net.Conn) {
	r.clients.Store(id, conn)
}

func (r *Registry) Disconnect(id string) {
	r.clients.Delete(id)
}

func (r *Registry) SendTo(id string, msg []byte) error {
	v, ok := r.clients.Load(id)
	if !ok {
		return fmt.Errorf("client %s not found", id)
	}
	_, err := v.(net.Conn).Write(msg)
	return err
}

func (r *Registry) Broadcast(msg []byte) {
	r.clients.Range(func(_, value any) bool {
		value.(net.Conn).Write(msg)
		return true
	})
}

// drain plays the role of a client in the demo: reads one message and prints it.
func drain(name string, conn net.Conn) {
	buf := make([]byte, 64)
	n, _ := conn.Read(buf)
	fmt.Printf("%s got: %s\n", name, buf[:n])
}

func main() {
	reg := New()

	// net.Pipe gives an in-memory (client, server) pair.
	aClient, aServer := net.Pipe()
	bClient, bServer := net.Pipe()

	reg.Connect("alice", aServer)
	reg.Connect("bob", bServer)

	var wg sync.WaitGroup

	wg.Go(func() { drain("alice", aClient) })
	reg.SendTo("alice", []byte("hi alice"))
	wg.Wait()

	fmt.Println(reg.SendTo("charlie", []byte("?"))) // client charlie not found

	wg.Go(func() { drain("alice", aClient) })
	wg.Go(func() { drain("bob", bClient) })
	reg.Broadcast([]byte("ping all"))
	wg.Wait()

	reg.Disconnect("alice")
	fmt.Println(reg.SendTo("alice", []byte("late"))) // client alice not found
}

/*
How Connection Registry works:

State: sync.Map keyed by client ID, value is net.Conn.
Connect = Store, Disconnect = Delete, SendTo = Load + Write,
Broadcast = Range + Write.

Key points:
- sync.Map fits write-once / delete-once / read-many over disjoint keys.
  Settled-key Load is an atomic read with no lock at all.
- map + RWMutex still wins for read-modify-write or when you need Len() /
  consistent snapshots. sync.Map has no Len() and Range is best-effort.
- Range during concurrent Store/Delete is safe but not transactional.
- Minimal SendTo blocks the broadcaster on a slow client. Real gateways
  give each conn a write goroutine + bounded chan; Broadcast does
  select { case ch <- msg: default: drop } so one slow client can't stall
  the rest.
- Real-world: nhooyr.io/websocket, gorilla/websocket, centrifugal/centrifuge.
*/
