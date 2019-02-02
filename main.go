package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Song Struct
type song struct {
	ID     uint32 `json:"id,omitempty"`
	User   string `json:"user,omitempty"`
	Source string `json:"source,omitempty"`
	Link   string `json:"link,omitempty"`
}

var songs []song

// Get User Songs
func GetUserSongs(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(songs)
}

func main() {
	songs = append(songs, song{ID: 1, User: "jimdaguy", Source: "Audiophiler", Link: "https://google.com"})
	songs = append(songs, song{ID: 28389, User: "jimdaguy", Source: "Spotify", Link: "https://google.com"})
	songs = append(songs, song{ID: 99494, User: "jimbus", Source: "Spotify", Link: "https://google.com"})
	songs = append(songs, song{ID: 300403, User: "jimdaguy", Source: "Audiophiler", Link: "https://google.com"})

	router := mux.NewRouter()
	router.HandleFunc("/users", GetUserSongs).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
