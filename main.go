package main

import (
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

// Song Struct
type Song struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	User      string        `bson:"user" json:"user"`
	Preferred bool          `bson:"preferred" json:"preferred"`
	Name      string        `bson:"name" json:"name"`
	Source    string        `bson:"source" json:"source"`
	Link      string        `bson:"link" json:"link"`
	Timestamp time.Time     `bson:"timestamp" json:"timestamp"`
}

// HeraldDAO struct - Herald Data Access Object
type HeraldDAO struct {
	Server   string
	Database string
}

// Global DB Object
var db *mgo.Database

// Collection name
const (
	COLLECTION = "songs"
)

// Connect - Establish a connection to the db
func (hDAO *HeraldDAO) Connect() {
	session, err := mgo.Dial(hDAO.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(hDAO.Database)
}

/////////////////
// CRUD functions for song collection in the Herald DB
/////////////////

/*
InsertHeraldSong - Insert song into the Herald DB songs collection
song - song to be inserted
*/
func (hDAO *HeraldDAO) InsertHeraldSong(song Song) error {
	err := db.C(COLLECTION).Insert(&song)
	return err
}

/*
GetHeraldSongs - Return list of herald songs
rpp - Number of results per page
page - Page number to start from - Pages are 0 indexed
*/
func (hDAO *HeraldDAO) GetHeraldSongs(rpp int, page int) ([]Song, error) {
	var songs []Song
	// Find all songs, sort by timestamp, skip to page, limit results
	err := db.C(COLLECTION).Find(bson.M{}).Sort("-timestamp").Skip(rpp * page).Limit(rpp).All(&songs)
	if err != nil {
		panic(err)
	}
	return songs, err
}

/*
GetHeraldUserSongs - Return list of herald songs the user owns
*/
func (hDAO *HeraldDAO) GetHeraldUserSongs(user string) ([]Song, error) {
	var songs []Song
	err := db.C(COLLECTION).Find(bson.M{"user": user}).Sort("-timestamp").All(&songs)
	if err != nil {
		panic(err)
	}
	return songs, err
}

/*
GetHeraldSong - Return information about herald song matching the id
id - id of the song to be returned
*/
func (hDAO *HeraldDAO) GetHeraldSong(id string) (Song, error) {
	var song Song
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&song)
	return song, err
}

/*
UpdateHeraldSong - Update information of song in the Herald DB
song - updated song information
*/
func (hDAO *HeraldDAO) UpdateHeraldSong(song Song) error {
	err := db.C(COLLECTION).UpdateId(song.ID, &song)
	return err
}

/*
DeleteHeraldSong - Remove song from Herald DB songs collection
song - song to be removed
*/
func (hDAO *HeraldDAO) DeleteHeraldSong(song Song) error {
	err := db.C(COLLECTION).Remove(&song)
	return err
}

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

func main() {
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
