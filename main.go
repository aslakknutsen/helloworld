package main

import "net/http"

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	http.HandleFunc("/", ping)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
