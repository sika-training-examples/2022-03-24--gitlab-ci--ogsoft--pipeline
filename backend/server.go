package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	j, _ := json.Marshal(map[string]string{
		"hostname": hostname,
	})
	w.Write(j)
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server started.")
	http.ListenAndServe(":80", nil)
}
