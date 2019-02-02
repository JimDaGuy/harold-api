package main

import (
	dao "herald-api/dao"
	song "herald-api/models"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

///////////////////
// Helper Functions
///////////////////

// SonglistAudiophiler - Get global list of songs from Audiophiler
func SonglistAudiophiler() (string, string) {
	return "beep", "boop"
}

// SongAudiophiler - Get specific song info from Audiophiler
func SongAudiophiler() {}

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

}

// GetUserPreferred - GET /users/{username}/preferred
func GetUserPreferred(w http.ResponseWriter, r *http.Request) {

}

// GetSong - GET /song
func GetSong(w http.ResponseWriter, r *http.Request) {

}

// MakeClip - POST /clip
func MakeClip(w http.ResponseWriter, r *http.Request) {

}

// DeleteClip - DELETE /removeClip
func DeleteClip(w http.ResponseWriter, r *http.Request) {

}

/*
ID        bson.ObjectId `bson:"_id" json:"id"`
	User      string        `bson:"user" json:"user"`
	Preferred bool          `bson:"preferred" json:"preferred"`
	Name      string        `bson:"name" json:"name"`
	Source    string        `bson:"source" json:"source"`
	Link      string        `bson:"link" json:"link"`
	Timestamp time.Time     `bson:"timestamp" json:"timestamp"`
*/

func main() {
	song := song.Song{
		ID:        bson.NewObjectId(),
		User:      "JimDaGuy",
		Preferred: true,
		Name:      "Bamb",
		Source:    "audiophiler",
		Link:      "google.com",
		Timestamp: time.Now(),
	}

	songsDAO := dao.SongsDAO{Database: "go-harold", Server: "localhost"}
	songsDAO.Connect()
	songsDAO.InsertHeraldSong(song)

	router := mux.NewRouter()
	router.HandleFunc("/getApSongs", GetAudiophilerSonglist).Methods("GET")
	router.HandleFunc("/getApSong", GetAudiophilerSong).Methods("GET")
	router.HandleFunc("/users/{username}", GetUserSongs).Methods("GET")
	router.HandleFunc("/users/{username}/preferred", GetUserPreferred).Methods("GET")
	router.HandleFunc("/song", GetSong).Methods("GET")
	router.HandleFunc("/clip", MakeClip).Methods("POST")
	router.HandleFunc("/removeClip", DeleteClip).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
