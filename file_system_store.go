package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() League {
	var league []Player
	// Following function will always start reading from start of the file.
	f.database.Seek(0, io.SeekStart)
	json.NewDecoder(f.database).Decode(&league)
	return league
}
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	league := f.GetLeague()
	player := league.Find(name)
	if player != nil {
		return player.Wins
	}

	return 0
}
func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{name, 1})
	}
	f.database.Seek(0, io.SeekStart)
	json.NewEncoder(f.database).Encode(league)
}
