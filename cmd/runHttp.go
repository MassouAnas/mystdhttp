package main

import (
	"fmt"
	"net/http"

	"github.com/MassouAnas/mystdhttp/router"
)

func runHttp(listenAddr string) error{
	s := http.Server{
		Addr: listenAddr,
		Handler: router.NewRouter(),
	}
	fmt.Printf("Starting HTTP listener at %s\n", listenAddr)
	return s.ListenAndServe()
}