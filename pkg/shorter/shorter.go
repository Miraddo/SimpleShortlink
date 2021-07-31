package shorter

import (
	"database/sql"
	"encoding/base64"
	"github.com/google/uuid"
	"log"
	"strings"
)

type ShorterFunc struct {
	DB *sql.DB
}

// MainUrl get the key and return main url from db
func (s *ShorterFunc) MainUrl(hash string) (string, error) {
	query := `SELECT url FROM urls WHERE key=$1`
	rows, err := s.DB.Query(query, hash)

	if err != nil {
		return "", err
	}

	var url string
	for rows.Next() {
		err = rows.Scan(&url)

		if err != nil {
			return "", err
		}
	}

	return url, nil
}

// ShortUrl generate key for each url and save into database
func (s *ShorterFunc) ShortUrl(u string) (string, error) {
	// generate random string -> hash
	// store url in db with that hash
	key := RandomGenerator(u)

	query := `INSERT INTO urls (url, key) VALUES ($1, $2)`

	_, err := s.DB.Exec(query, u, key)

	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	return key, nil

}

func RandomGenerator(s string) string {
	// encode with base64 on the main url + uuid and return the hash
	// with uuid
	hash := base64.StdEncoding.EncodeToString([]byte(s + strings.Replace(uuid.New().String(), "-", "", -1)))

	// return 6 character
	return hash[len(hash)-8 : len(hash)-2]
}
