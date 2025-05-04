package helper

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

// TemporaryKey speichert temporäre Sitzungsschlüssel
type TemporaryKey struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
}

// GenerateTemporaryKey erzeugt einen neuen temporären ECDH-Schlüssel
func GenerateTemporaryKey() (*TemporaryKey, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("fehler beim Generieren des Schlüssels: %v", err)
	}

	publicKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)

	return &TemporaryKey{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}

func IsKeyValid(sessionID string) bool {
	// Dummy
	return sessionID == "session1"
}
