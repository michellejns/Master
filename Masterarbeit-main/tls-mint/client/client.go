package main

import (
	"log"
	"tls-example/helper"

	"github.com/bifurcation/mint"
)

func main() {
	// Adresse des Servers
	serverAddr := "127.0.0.1:4433"

	// Zertifikat und privater Schlüssel erzeugen für den Client
	_, cert, err := helper.MakeNewSelfSignedCert()
	if err != nil {
		log.Fatalf("Fehler beim Erzeugen des Zertifikats: %v", err)
	}

	// TLS-Konfiguration für den Client
	config := &mint.Config{
		InsecureSkipVerify: true,
		Certificates: []*mint.Certificate{
			cert, // Zertifikat des Servers
		},
		CipherSuites: []mint.CipherSuite{
			mint.TLS_AES_128_GCM_SHA256,
			mint.TLS_AES_256_GCM_SHA384,
		},
		Groups: []mint.NamedGroup{
			mint.X25519,
			mint.P256,
		},
		SignatureSchemes: []mint.SignatureScheme{
			mint.ECDSA_P256_SHA256,
		},
	}

	// Verbindung zum Server herstellen
	conn, err := mint.Dial("tcp", serverAddr, config)
	if err != nil {
		log.Fatalf("Fehler beim Aufbau der TLS-Verbindung: %v", err)
	}
	defer conn.Close()

	log.Println("TLS-Verbindung erfolgreich hergestellt")

	// Nachricht senden
	_, err = conn.Write([]byte("Hallo Server, hier ist der Client über mint!"))
	if err != nil {
		log.Fatalf("Fehler beim Senden der Nachricht: %v", err)
	}

	// Antwort lesen
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatalf("Fehler beim Lesen der Antwort: %v", err)
	}

	log.Printf("Antwort vom Server: %s", string(buffer[:n]))
}
