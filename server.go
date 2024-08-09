package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Player struct{
	Name string
	Wins int
}
// PlayerStore stores score information about players.
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

// PlayerServer is a HTTP interface for player information.
// Example of Embedding
type PlayerServer struct {
	store PlayerStore
	http.Handler
}

// NewPlayerServer creates a PlayerServer with routing configured.
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router

	return p
}
const jsonContentType = "application/json"
func (p *PlayerServer)leagueHandler(w http.ResponseWriter, r *http.Request){
	
	//Adding header as content-type = application/json
	w.Header().Set("content-type",jsonContentType)

	// Encoding and Decoding json for best practices
	json.NewEncoder(w).Encode(p.store.GetLeague())

}

func(p *PlayerServer)playersHandler(w http.ResponseWriter,r *http.Request){
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
	}
func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}