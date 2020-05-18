package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// RedirectHandler ...
func (app *App) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: add panic handler

	vars := mux.Vars(r)
	hash := vars["hash"]

	var rec URLHash
	err := app.DB.Model(&rec).
		Where("hash=?", hash).Select()
	if err != nil {
		// TODO: return 500 error response
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.Redirect(w, r, rec.URL, http.StatusSeeOther)
}
