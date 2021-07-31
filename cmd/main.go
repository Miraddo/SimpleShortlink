package main

import (
	"database/sql"
	"fmt"
	"github.com/Miraddo/SimpleShortlink/config"
	"github.com/Miraddo/SimpleShortlink/pkg/handlers"
	"github.com/Miraddo/SimpleShortlink/pkg/shorter"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	// connect to database
	data, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s database=%s sslmode=%s",
		config.DatabaseHost,
		config.DatabaseUser,
		config.DatabasePass,
		config.DatabaseName,
		config.DatabaseSSLMode),
	)

	if err != nil {
		panic(err)
	}

	err = data.Ping()
	if err != nil {
		panic(err)
	}

	st := &shorter.ShorterFunc{
		DB: data,
	}

	hf := &handlers.HTTPHandler{
		Shorter: st,
	}
	// create two handle requests
	http.HandleFunc("/short", hf.ShortUrlFunc)
	http.HandleFunc("/url", hf.MainUrlFunc)

	// basic webserver
	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		return
	}

}
