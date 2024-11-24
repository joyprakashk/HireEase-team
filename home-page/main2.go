package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Room represents a video chat room
type Room struct {
	ID      string
	Clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

// Message represents the WebSocket message structure
type Message struct {
	Type       string          `json:"type"`
	RoomID     string          `json:"roomId"`
	Content    string          `json:"content"`
	SignalData json.RawMessage `json:"signalData,omitempty"`
}

var (
	// Upgrader for WebSocket connections
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all origins in development
		},
	}

	// Global room management
	rooms      = make(map[string]*Room)
	roomsMutex sync.Mutex
)

// GetOrCreateRoom returns an existing room or creates a new one
func getOrCreateRoom(roomID string) *Room {
	roomsMutex.Lock()
	defer roomsMutex.Unlock()

	if room, exists := rooms[roomID]; exists {
		return room
	}

	room := &Room{
		ID:      roomID,
		Clients: make(map[*websocket.Conn]bool),
	}
	rooms[roomID] = room
	return room
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	defer conn.Close()

	log.Printf("New WebSocket connection established")

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Read error: %v", err)
			cleanupConnection(conn)
			break
		}

		switch msg.Type {
		case "join":
			handleJoinRoom(conn, msg)
		case "leave":
			handleLeaveRoom(conn, msg)
		case "signal":
			handleSignal(conn, msg)
		default:
			log.Printf("Unknown message type: %s", msg.Type)
		}
	}
}

func handleJoinRoom(conn *websocket.Conn, msg Message) {
	room := getOrCreateRoom(msg.RoomID)

	room.mu.Lock()
	room.Clients[conn] = true
	clientCount := len(room.Clients)
	room.mu.Unlock()

	log.Printf("Client joined room %s. Total clients: %d", msg.RoomID, clientCount)

	response := Message{
		Type:    "joined",
		RoomID:  msg.RoomID,
		Content: "Successfully joined room",
	}
	err := conn.WriteJSON(response)
	if err != nil {
		log.Printf("Error sending join confirmation: %v", err)
	}
}

func handleLeaveRoom(conn *websocket.Conn, msg Message) {
	cleanupConnection(conn)

	log.Printf("Client left room %s", msg.RoomID)
}

func handleSignal(conn *websocket.Conn, msg Message) {
	roomsMutex.Lock()
	room, exists := rooms[msg.RoomID]
	roomsMutex.Unlock()

	if !exists {
		log.Printf("Attempt to signal in non-existent room: %s", msg.RoomID)
		return
	}

	room.mu.Lock()
	defer room.mu.Unlock()

	for client := range room.Clients {
		if client != conn {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error broadcasting signal: %v", err)
				client.Close()
				delete(room.Clients, client)
			}
		}
	}
}

func cleanupConnection(conn *websocket.Conn) {
	roomsMutex.Lock()
	defer roomsMutex.Unlock()

	for _, room := range rooms {
		room.mu.Lock()
		if _, exists := room.Clients[conn]; exists {
			delete(room.Clients, conn)

			if len(room.Clients) == 0 {
				delete(rooms, room.ID)
			}
		}
		room.mu.Unlock()
	}
}

func enableCORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func main() {

	http.HandleFunc("/ws", handleWebSocket)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", enableCORS(fs))

	port := ":8080"
	log.Printf("Server starting on%s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
