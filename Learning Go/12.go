package main

import (
	"fmt"
	"net/http"
)

type pair struct {
	x int
	y int
}

// ServeHTTP implements http.Handler.
func (p pair) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("LessGooo!"))
}

func server() {
	go func() {
		err := http.ListenAndServe(":6969", pair{})
		fmt.Println(err)
	}()

	reqServer()
}

func reqServer() {
	_, err := http.Get("http://localhost:6969")
	fmt.Println(err)
}

func main() {
	server()
}

// TODO: not working
