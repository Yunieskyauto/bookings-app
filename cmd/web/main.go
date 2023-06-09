package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/tsawler/bookings-app/internal/config"
	"github.com/tsawler/bookings-app/internal/driver"
	"github.com/tsawler/bookings-app/internal/handlers"
	"github.com/tsawler/bookings-app/internal/models"
	"github.com/tsawler/bookings-app/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {

	db, err := run()
	defer db.SQL.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB, error) {
	// what am I going to store in the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Restriction{})
	gob.Register(models.Room{})

	// change this to true when in production
	app.InProduction = false

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// Connect to database
	log.Println("Connecting to database")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=postgres password=freire16")
	if err != nil {
		log.Fatal("Cannot conect to database! Dying....")
	}

	log.Println("Connected to database")
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	return db, nil
}
