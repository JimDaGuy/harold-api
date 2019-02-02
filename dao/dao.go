package dao

import (
	song "herald-api/models"

	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SongsDAO struct - Songs Data Access Object
type SongsDAO struct {
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
func (hDAO *SongsDAO) Connect() {
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
func (hDAO *SongsDAO) InsertHeraldSong(song song.Song) error {
	err := db.C(COLLECTION).Insert(&song)
	return err
}

/*
GetHeraldSongs - Return list of herald songs
rpp - Number of results per page
page - Page number to start from - Pages are 0 indexed
*/
func (hDAO *SongsDAO) GetHeraldSongs(rpp int, page int) ([]song.Song, error) {
	var songs []song.Song
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
func (hDAO *SongsDAO) GetHeraldUserSongs(user string) ([]song.Song, error) {
	var songs []song.Song
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
func (hDAO *SongsDAO) GetHeraldSong(id string) (song.Song, error) {
	var song song.Song
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&song)
	return song, err
}

/*
UpdateHeraldSong - Update information of song in the Herald DB
song - updated song information
*/
func (hDAO *SongsDAO) UpdateHeraldSong(song song.Song) error {
	err := db.C(COLLECTION).UpdateId(song.ID, &song)
	return err
}

/*
DeleteHeraldSong - Remove song from Herald DB songs collection
song - song to be removed
*/
func (hDAO *SongsDAO) DeleteHeraldSong(song song.Song) error {
	err := db.C(COLLECTION).Remove(&song)
	return err
}
