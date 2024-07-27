package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
		player:= strings.TrimPrefix(r.URL.Path,"/players/")
		if player == "Pepper"{
			fmt.Fprint(w,"20")
		}
		if player == "Floyd"{
			fmt.Fprint(w,"10")
		}
}
func main(){
	handler := http.HandlerFunc(PlayerServer)
	log.Fatal(http.ListenAndServe(":5000",handler))
}
