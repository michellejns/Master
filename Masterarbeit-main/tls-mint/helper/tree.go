package helper

import (
	"sync"
)

// Tree speichert temporäre Schlüssel in einer Baumstruktur
type Tree struct {
	sync.Mutex
	Keys map[string]*TemporaryKey
}

// NewTree initialisiert einen leeren Schlüsselbaum
func NewTree() *Tree {
	return &Tree{
		Keys: make(map[string]*TemporaryKey),
	}
}

// AddKey fügt einen neuen Schlüssel zum Baum hinzu
func (t *Tree) AddKey(sessionID string, key *TemporaryKey) {
	t.Lock()
	defer t.Unlock()
	t.Keys[sessionID] = key
}

// GetKey gibt einen Schlüssel anhand der Session-ID zurück
func (t *Tree) GetKey(sessionID string) (*TemporaryKey, bool) {
	t.Lock()
	defer t.Unlock()
	key, exists := t.Keys[sessionID]
	return key, exists
}

// RemoveKey löscht gezielt einen Schlüssel (Puncturing)
func (t *Tree) RemoveKey(sessionID string) {
	t.Lock()
	defer t.Unlock()
	delete(t.Keys, sessionID)
}
