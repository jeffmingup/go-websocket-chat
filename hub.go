// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "log"

type RoomId string

type broadcast struct {
	message []byte
	roomId  RoomId
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {

	//房间的客户端列表
	room map[RoomId]map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan broadcast

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan broadcast),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		room:       make(map[RoomId]map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			if nil == h.room[client.roomId] {
				h.room[client.roomId] = map[*Client]bool{}
			}
			h.room[client.roomId][client] = true
			log.Println(h.room)
		case client := <-h.unregister:
			if _, ok := h.room[client.roomId][client]; ok {
				delete(h.room[client.roomId], client)
				close(client.send)
			}
		case broadcast := <-h.broadcast:
			for client := range h.room[broadcast.roomId] {
				select {
				case client.send <- broadcast.message:
				default:
					close(client.send)
					delete(h.room[broadcast.roomId], client)
				}
			}
		}
	}
}
