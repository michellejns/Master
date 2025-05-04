package main

import (
	"log"
	"tls-example/helper" // Importiere das Helper-Package, das das Zertifikat erstellt

	"github.com/bifurcation/mint"
)

func main() {
	// Port des Servers
	port := "127.0.0.1:4433"

	// selbstsigniertes Zertifikat
	_, cert, err := helper.MakeNewSelfSignedCert() // Der private Schlüssel wird nicht mehr benötigt
	if err != nil {
		log.Fatalf("Fehler beim Erzeugen des Zertifikats: %v", err)
	}

	// TLS-Konfiguration für den Server
	config := &mint.Config{
		InsecureSkipVerify: true,
		Certificates:       []*mint.Certificate{cert}, // Zertifikat hinzufügen
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

	// Listener erstellen
	listener, err := mint.Listen("tcp", port, config)
	if err != nil {
		log.Fatalf("Fehler beim Starten des Listeners: %v", err)
	}

	log.Printf("TLS-Server läuft auf %s", port)

	// Warten auf eingehende Verbindungen
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Fehler beim Verbindungsaufbau: %v", err)
			continue
		}

		// Umwandlung der Verbindung in einen mint.Conn-Typ
		mintConn, ok := conn.(*mint.Conn)
		if !ok {
			log.Println("Verbindung konnte nicht als *mint.Conn umgewandelt werden")
			conn.Close()
			continue
		}

		// Handshake wird hier von mintConn automatisch behandelt
		log.Printf("Handshake abgeschlossen mit %s", conn.RemoteAddr())

		// Jetzt können wir sicher Daten lesen und schreiben
		go handleConnection(mintConn)
	}
}

// Funktion zum Verarbeiten der Client-Verbindung
func handleConnection(conn *mint.Conn) {
	defer conn.Close()

	// Daten lesen
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Fehler beim Lesen der Verbindung: %v", err)
		return
	}

	// Empfangene Daten verarbeiten
	log.Printf("Empfangene Daten: %s", string(buffer[:n]))

	// Antwort senden
	_, err = conn.Write([]byte("Verbindung erfolgreich"))
	if err != nil {
		log.Printf("Fehler beim Senden der Antwort: %v", err)
	}
}
