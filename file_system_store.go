package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	var league []Player
	// Following function will always start reading from start of the file.
	f.database.Seek(0, io.SeekStart)
	json.NewDecoder(f.database).Decode(&league)
	return league
}
