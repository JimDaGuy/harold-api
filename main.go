package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	dao "github.com/jimdaguy/herald-api/dao"

	"github.com/gorilla/mux"
)

var songsDAO dao.SongsDAO

///////////////////
// Helper Functions
///////////////////

// SonglistAudiophiler - Get global list of songs from Audiophiler
func SonglistAudiophiler() (string, string) {
	return "beep", "boop"
}

// SongAudiophiler - Get specific song info from Audiophiler
func SongAudiophiler() {}

// DownloadFile - Download file from the given url
func DownloadFile(localpath string, fileurl string) error {
	resp, err := http.Get(fileurl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath.Join(localpath))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

/*
CreateSong -
id: Song ID from Audiophiler
user: Username of user that the song is saved to
link: S3 link to unsplit song
startSec: beginning of song split
endSec: end of song split
*/
func CreateSong(id string, user string, source string, link string, startSec int, endSec int) {
	// Download song from s3 link
	// Check params
	// Split downloaded song to parameter times
	// Upload split song to s3
	// Create new song object with id, user, and link to split song
	// Store song object in db
}

/////////////////////
// API Route Handlers
/////////////////////

// GetAudiophilerSonglist - GET /getApSongs
func GetAudiophilerSonglist(w http.ResponseWriter, r *http.Request) {

}

// GetAudiophilerSong - GET /getApSong
func GetAudiophilerSong(w http.ResponseWriter, r *http.Request) {

}

// GetUserSongs - GET /users/{username}/
func GetUserSongs(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	song, err := songsDAO.GetHeraldUserSongs(params["username"])
	if err != nil {
		panic(err)
	}

	songsJSON, _ := json.Marshal(song)

	w.Header().Set("Content-Type", "application/json")
	w.Write(songsJSON)
}

// GetUserPreferred - GET /users/{username}/preferred
func GetUserPreferred(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	song, err := songsDAO.GetHeraldUserPreferredSongs(params["username"])
	if err != nil {
		panic(err)
	}

	songsJSON, _ := json.Marshal(song)

	w.Header().Set("Content-Type", "application/json")
	w.Write(songsJSON)
}

// GetSong - GET /song
func GetSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	song, err := songsDAO.GetHeraldSong(params["id"])
	if err != nil {
		panic(err)
	}

	songsJSON, _ := json.Marshal(song)

	w.Header().Set("Content-Type", "application/json")
	w.Write(songsJSON)
}

// MakeClip - POST /clip
func MakeClip(w http.ResponseWriter, r *http.Request) {
	// CreateSong(id string, user string, name string, source string, link string, startSec int, endSec int)
	filename, ok := r.URL.Query()["name"]
	if !ok || len(filename[0]) < 1 {
		println("Missing name parameter")
		return
	}

	link, ok2 := r.URL.Query()["link"]
	if !ok2 || len(link[0]) < 1 {
		println("Missing link parameter")
		return
	}

	name := filename[0]
	url := link[0]
	filepath := fmt.Sprintf("temp/%s", name)

	err := DownloadFile(filepath, url)
	if err != nil {
		panic(err)
	}
}

// DeleteClip - DELETE /removeClip
func DeleteClip(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	song, err := songsDAO.GetHeraldSong(params["id"])
	if err != nil {
		panic(err)
	}

	err2 := songsDAO.DeleteHeraldSong(song)
	if err2 != nil {
		panic(err2)
	}

	success, _ := json.Marshal("Successfully removed")

	w.Header().Set("Content-Type", "application/json")
	w.Write(success)
}

func main() {
	songsDAO := dao.SongsDAO{Database: "go-harold", Server: "localhost"}
	songsDAO.Connect()

	router := mux.NewRouter()
	router.HandleFunc("/getApSongs", GetAudiophilerSonglist).Methods("GET")
	router.HandleFunc("/getApSong", GetAudiophilerSong).Methods("GET")
	router.HandleFunc("/users/{username}", GetUserSongs).Methods("GET")
	router.HandleFunc("/users/{username}/preferred", GetUserPreferred).Methods("GET")
	router.HandleFunc("/song/{id}", GetSong).Methods("GET")
	router.HandleFunc("/clip", MakeClip).Methods("POST")
	router.HandleFunc("/removeClip/{id}", DeleteClip).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
