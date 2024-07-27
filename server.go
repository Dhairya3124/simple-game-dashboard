package main

import (
	"fmt"
	"log"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w,"20");
}
func main(){
	handler := http.HandlerFunc(PlayerServer)
	log.Fatal(http.ListenAndServe(":5000",handler))
}
