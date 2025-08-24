package tls

import (
	"crypto/tls"
	"log"
	"net/http"
)

const (
	port         = ":8443"
	responseBody = "Hello, TLS!"
)

func Server() {
	log.Default().Println("Starting TLS server ....")
	cert, err := tls.LoadX509KeyPair("tls/server.crt", "tls/server.key")
	if err != nil {
		log.Fatalf("Failed to load X509 key pair: %v", err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	router := http.NewServeMux()
	router.HandleFunc("/", handleRequest)

	server := &http.Server{
		Addr:      port,
		Handler:   router,
		TLSConfig: config,
	}


	log.Printf("Listening on %s...", port)
	err = server.ListenAndServeTLS("", "")
	if err != nil {
  		log.Fatalf("Failed to start server: %v", err)
	}	
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseBody))
}