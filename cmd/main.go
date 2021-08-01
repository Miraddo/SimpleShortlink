package main

import (
	"database/sql"
	"fmt"
	"github.com/Miraddo/SimpleShortlink/config"
	"github.com/Miraddo/SimpleShortlink/pkg/handlers"
	"github.com/Miraddo/SimpleShortlink/pkg/shorter"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var rootCmd = &cobra.Command{
	Use:   "serv",
	Short: "Serv HTTP Server",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var servCmd = &cobra.Command{
	Use:   "serv",
	Short: "Return Short Url to Main Url",
	Long:  "Return Short Url to Main Url",
	Run: func(cmd *cobra.Command, args []string) {
		// connect to database
		data, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s database=%s sslmode=%s",
			config.DBHost,
			config.DBUser,
			config.DBPass,
			config.DBName,
			config.DBSSLMode),
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
		logger, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}

		hf := &handlers.HTTPHandler{
			Shorter: st,
			Logger:  logger.Sugar(),
		}

		// create two handle requests
		http.HandleFunc("/short", hf.ShortUrlFunc)
		http.HandleFunc("/url", hf.MainUrlFunc)

		// basic webserver
		err = http.ListenAndServe(":8080", nil)
		log.Println("Server is UP!!! \n Please run 127.0.0.1:8080 in your system!")
		if err != nil {
			return
		}

	},
}

var shortCmd = &cobra.Command{
	Use:   "short",
	Short: "Make Url Shorter",
	Long:  "Make Url Shorter",
	Run: func(cmd *cobra.Command, args []string) {
		// connect to database
		data, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s database=%s sslmode=%s",
			config.DBHost,
			config.DBUser,
			config.DBPass,
			config.DBName,
			config.DBSSLMode),
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

		if len(args) < 1 {
			log.Println("Please Pass the Key to return Main URL!")
			return
		}

		if handlers.AvailableUrl(args[0]) {
			s, err := st.ShortUrl(args[0])

			if err != nil {
				panic(err)
			}

			fmt.Println("Key: ", s)
		} else {
			fmt.Println("Ops! Url is not correct!")
		}

	},
}
var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "Return Short Url to Main Url",
	Long:  "Return Short Url to Main Url",
	Run: func(cmd *cobra.Command, args []string) {
		// connect to database
		data, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s database=%s sslmode=%s",
			config.DBHost,
			config.DBUser,
			config.DBPass,
			config.DBName,
			config.DBSSLMode),
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

		if len(args) < 1 {
			log.Println("Please Pass an URL to make it short!")
			return
		}
		s, err := st.MainUrl(args[0])
		if err != nil {
			panic(err)
		}
		if s == "" {
			s = "Not Found!"
		}
		fmt.Println("Url: ", s)
	},
}

func main() {
	rootCmd.AddCommand(shortCmd, urlCmd, servCmd)
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}

}
