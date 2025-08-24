package main

import (
	"sync"

	"github.com/sankalp/quantum_safe_tls/tls"
)

func main(){
    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        defer wg.Done()
        tls.Server()
    }()

    go func() {
        defer wg.Done()
        tls.Client()
    }()

    wg.Wait()
}