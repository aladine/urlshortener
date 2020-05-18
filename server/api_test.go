package server

import (
	"net/http"
	"testing"
	"net/http/httptest"
	"bytes"
	"time"

	"github.com/gorilla/mux"
	"github.com/go-pg/pg/v10"
)

func startDatabaseTest() *pg.DB {

	cnf := &pg.Options{
		Addr:     "rosie.db.elephantsql.com:5432",
		User:     "xzcomuke",
		Database: "xzcomuke",
		Password: "H9zmv3mPJlFszhqoUjbOWK7UG-ag4Nde",
	}

	return pg.Connect(cnf)
}

func initServer() App{
	app := App{}
	app.DB = startDatabaseTest()
	app.Router = mux.NewRouter()
	app.initializeRoutes()

	srv := &http.Server{
		Handler:      app.Router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	go srv.ListenAndServe()
	return app
}

func (a *App) executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    a.Router.ServeHTTP(rr, req)

	return rr
}
func TestCreateURLHandler(t *testing.T){
	app:=initServer()

	var testBody = `{
        "title": "aldi",
        "user_id": 2,
        "url": "https://www.aldi.com.au/en/special-buys/special-buys-sat-23-may/saturday-detail-wk21/ps/p/delonghi-argento-kettle-1/?pk_campaign=au_product_newsletter&pk_kwd=2020-05-18_11-58"
	}`

	var jsonStr = []byte(testBody)
    req, _ := http.NewRequest("POST", "/url", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    response := app.executeRequest(req)
	if response.Code != http.StatusCreated{
		t.Errorf("failed to call create shorten url %v", response)
	}
}
