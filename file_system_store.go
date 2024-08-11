package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	var league []Player
	// Following function will always start reading from start of the file.
	f.database.Seek(0, io.SeekStart)
	json.NewDecoder(f.database).Decode(&league)
	return league
}
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var score int
	for _, player := range f.GetLeague() {
		if player.Name == name {
			score = player.Wins
			break

		}
	}

	return score
}
func (f *FileSystemPlayerStore)RecordWin(name string){
	league:=f.GetLeague()
	for i, player := range  league{
		if player.Name == name {
			league[i].Wins++

		}
	}
	f.database.Seek(0,io.SeekStart)
	json.NewEncoder(f.database).Encode(league)
}