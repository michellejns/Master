package helper

import (
	"fmt"
	"sync"
)

// PuncturingTree erweitert die Tree-Struktur um Puncturing-Mechanismen
type PuncturingTree struct {
	sync.Mutex
	Keys map[string]*TemporaryKey
}

// NewPuncturingTree initialisiert den Baum
func NewPuncturingTree() *PuncturingTree {
	return &PuncturingTree{
		Keys: make(map[string]*TemporaryKey),
	}
}

// AddKey fügt einen neuen Schlüssel hinzu
func (t *PuncturingTree) AddKey(sessionID string, key *TemporaryKey) {
	t.Lock()
	defer t.Unlock()
	t.Keys[sessionID] = key
}

// RemoveKey entfernt einen Schlüssel (Puncturing)
func (t *PuncturingTree) RemoveKey(sessionID string) {
	t.Lock()
	defer t.Unlock()
	delete(t.Keys, sessionID)
	fmt.Println("Schlüssel für Session", sessionID, "wurde invalidiert (Puncturing).")
}

// IsKeyValid prüft, ob ein Schlüssel noch gültig ist
func (t *PuncturingTree) IsKeyValid(sessionID string) bool {
	t.Lock()
	defer t.Unlock()
	_, exists := t.Keys[sessionID]
	return exists
}
