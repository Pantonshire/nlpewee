// +build debug

package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func init() {
	addr := "localhost:6060"
	go func() {
		fmt.Println("Starting debug server at", addr)
		fmt.Println(http.ListenAndServe(addr, nil))
	}()
}
