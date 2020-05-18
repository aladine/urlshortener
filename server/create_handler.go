package server

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// GetRandomString ...
func GetRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// CreateURLHandler ...
func (app *App) CreateURLHandler(w http.ResponseWriter, r *http.Request) {

	var rec *URLHash
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rec); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// validate url
	if err := validateURL(rec.URL); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// TODO: search for current database whether url already submitted
	// var recs []URLHash
	// err := app.DB.Model(&recs).
	// 	Where("user_id=?", userID).
	// 	Where("url=?", rec.URL).
	// 	Select()
	// if err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// hash is 4 random string which was designed to be unique for all the active url
	rec.Hash = GetRandomString(4)
	rec.Created = time.Now().UTC()

	if err := app.DB.Insert(rec); err != nil {
		// TODO: check if err is duplicate key, then retry insertion with another random hash
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// return 201 response code with new shorten url
	respondWithJSON(w, http.StatusCreated, rec)
}

// URLHash ...
type URLHash struct {
	UserID int64  `json:"user_id"`
	Title  string `json:"title"`
	URL    string `json:"url"`
	Hash   string `json:"hash"`
	// ExpiredTime time.Time
	Created time.Time `json:"created"`
}
