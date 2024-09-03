package main

import (
	"log"
	"net/http"

	poker "github.com/Dhairya3124/simple-game-dashboard.git"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server, err := poker.NewPlayerServer(store)
	if err != nil {
		log.Fatalf("problem creating player server %v", err)
	}

	log.Fatal(http.ListenAndServe(":5000", server))

}
