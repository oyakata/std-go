package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/oyakata/std-go/prof/memeater"
)
import _ "net/http/pprof"

func main() {
	http.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	length := memeater.Eat()
	fmt.Fprintf(w, "[%v] size=%d\n", now, length)
}
