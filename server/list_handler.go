package server

import (
	"fmt"
	"net/http"
)

// TODO: using golang html/template to render the output

// ListURLHandler ...
func (app *App) ListURLHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: add panic handler
	userID := 2 // TODO: get user id from query param

	// searching db to get all url belongs to userID
	var recs []URLHash
	err := app.DB.Model(&recs).Where("user_id=?", userID).Select()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Fprintf(w, "UserID: %d", userID)
	fmt.Fprintln(w, "\nList of all the shorten URLs:")
	for _, rec := range recs {
		fmt.Fprintf(w,
			"\nLink: %s %s",
			rec.URL,
			"127.0.0.1:8000/url/"+rec.Hash,
		)
	}
}
