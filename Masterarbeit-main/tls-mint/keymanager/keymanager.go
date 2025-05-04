package keymanager

import (
	"log"
	"sync"
	"time"
)

var (
	usedSessions = make(map[string]time.Time)
	mu           sync.Mutex
	//  Gültigkeitsdauer für temporäre Session-IDs
	sessionLifetime = 10 * time.Minute
)

// IsSessionValid prüft, ob die Session-ID gültig und noch nicht verwendet wurde.
func IsSessionValid(sessionID string) bool {
	mu.Lock()
	defer mu.Unlock()

	expiry, exists := usedSessions[sessionID]
	if !exists {
		return true // Noch nicht verwendet
	}
	if time.Now().After(expiry) {
		delete(usedSessions, sessionID)
		return true // Abgelaufen, wieder nutzbar
	}
	return false // Replay detected
}

// RegisterSession speichert die Session-ID als verwendet.
func RegisterSession(sessionID string) {
	mu.Lock()
	defer mu.Unlock()

	usedSessions[sessionID] = time.Now().Add(sessionLifetime)
	log.Printf("Session registriert: %s (gültig bis %s)", sessionID, usedSessions[sessionID])
}
