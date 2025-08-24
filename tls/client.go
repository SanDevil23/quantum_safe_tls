package tls

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	url = "https://localhost:8443/"
)

func Client() {
	cert, err := os.ReadFile("tls/ca.crt")
	if err != nil {
		log.Fatalf("Failed to read certificate file: %v", err)
	}

	// creating an empty certificate pool
	caCertPool := x509.NewCertPool()

	// appending the self-signed certificate to the pool
	if ok := caCertPool.AppendCertsFromPEM(cert); !ok {
		log.Fatal("Failed to append CA certificate")
	}

	// creating a TLS configuration object
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,						// self-signed CA certificate as the root CA
	}

	tr:= &http.Transport{
		TLSClientConfig: tlsConfig,	// setting the TLS configuration
	}
	client := &http.Client{Transport: tr}

	// make an http GET request to the url

	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Failed to make GET request: %v", err)
	}



	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	log.Printf("Response status: %s", resp.Status)
	log.Printf("Response body: %s", string(body))

}