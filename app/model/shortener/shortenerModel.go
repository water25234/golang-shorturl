package modelshortener

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/water25234/golang-shorturl/server"
)

var err error
var tx *sql.Tx
var succ sql.Result

type Shortener struct {
	ShortenerID int64  `json:"shortener_id"`
	CreateDate  string `json:"create_date"`
	ModifyDate  string `json:"modify_date"`
	URLEncode   string `json:"url_encode"`
	URL         string `json:"url"`
	IsClose     bool   `json:"is_close"`
}

func GetShortenerID(URLEncode string) (*Shortener, error) {
	var shortener Shortener
	err = server.DB.QueryRow("SELECT url_encode, url FROM shortener Where url_encode = $1;", URLEncode).Scan(
		&shortener.URLEncode,
		&shortener.URL)
	if err != nil {
		panic(err)
	}
	return &shortener, err
}

func GetLastShortenerID() (*Shortener, error) {

	var shortener Shortener
	err = server.DB.QueryRow("SELECT shortener_id FROM shortener ORDER BY shortener_id DESC LIMIT 1;").Scan(
		&shortener.ShortenerID)
	if err != nil {
		shortener.ShortenerID = 0
	}
	return &shortener, err
}

func SaveShortenerUrl(URLEncode string, URL string) (*Shortener, error) {

	var shortener Shortener
	err = server.DB.QueryRow("SELECT * FROM shortener WHERE url_encode = $1;", URLEncode).Scan(
		&shortener.ShortenerID,
		&shortener.CreateDate,
		&shortener.ModifyDate,
		&shortener.URLEncode,
		&shortener.URL,
		&shortener.IsClose)

	if shortener.ShortenerID != 0 {
		return &shortener, err
	}

	tx, err = server.DB.Begin()
	if err != nil {
		panic(err)
	}

	sql := "INSERT INTO shortener (create_date, modify_date, url_encode, url, is_close) VALUES (now(), now(), $1, $2, false);"
	_, err = server.DB.Exec(sql, URLEncode, URL)
	if err != nil {
		panic(err)
	}

	err = server.DB.QueryRow("SELECT * FROM shortener WHERE url_encode = $1;", URLEncode).Scan(
		&shortener.ShortenerID,
		&shortener.CreateDate,
		&shortener.ModifyDate,
		&shortener.URLEncode,
		&shortener.URL,
		&shortener.IsClose)

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		if err != nil {
			panic(err)
		}
	}

	return &shortener, err
}
