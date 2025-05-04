package helper

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"

	"github.com/bifurcation/mint"
)

// MakeNewSelfSignedCert erstellt ein selbstsigniertes Zertifikat und den dazugehörigen privaten Schlüssel
func MakeNewSelfSignedCert() (*rsa.PrivateKey, *mint.Certificate, error) {
	// Erstelle den privaten Schlüssel
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	// Erstelle eine Seriennummer für das Zertifikat
	serialNumber, err := rand.Int(rand.Reader, big.NewInt(1<<62))
	if err != nil {
		return nil, nil, err
	}

	// Erstelle das Zertifikat
	tmpl := x509.Certificate{
		SerialNumber: serialNumber,
		Subject:      pkix.Name{CommonName: "localhost"}, // Setze den Common Name auf "localhost"
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(24 * time.Hour), // Gültigkeit des Zertifikats für 24 Stunden
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}

	// Erstelle das Zertifikat im DER-Format
	certDER, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, &priv.PublicKey, priv)
	if err != nil {
		return nil, nil, err
	}

	// Parse das Zertifikat
	certParsed, err := x509.ParseCertificate(certDER)
	if err != nil {
		return nil, nil, err
	}

	// Erstelle das Mint-Zertifikat (mit der Zertifikatskette und dem privaten Schlüssel)
	cert := &mint.Certificate{
		Chain:      []*x509.Certificate{certParsed},
		PrivateKey: priv,
	}

	return priv, cert, nil
}
