package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)
type PlayerStore interface{
	GetPlayerScore(name string)int
}
type PlayerServer struct{
	store PlayerStore
}
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		player:= strings.TrimPrefix(r.URL.Path,"/players/")
		score:=p.store.GetPlayerScore(player)
		if score == 0{
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprint(w,score)
		
}
func GetPlayerScore(name string)string{
	if name == "Pepper"{
		return "20"
	}
	if name == "Floyd"{
		return "10"
	}
	return ""

}
type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}
func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
