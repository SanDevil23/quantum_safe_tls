package main

import (
	"github.com/sankalp/quantum_safe_tls/tls"
)

func main(){
	go tls.Server()
	go tls.Client()
}