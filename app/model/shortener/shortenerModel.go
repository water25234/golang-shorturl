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
	ShortenerID int    `json:"shortener_id"`
	CreateDate  string `json:"create_date"`
	ModifyDate  string `json:"modify_date"`
	URLEncode   string `json:"url_encode"`
	URL         string `json:"url"`
	IsClose     bool   `json:"is_close"`
}

func SaveShortenerUrl(args ...interface{}) (*Shortener, error) {

	var shortener Shortener
	err = server.DB.QueryRow("SELECT * FROM shortener WHERE url_encode = $1;", args).Scan(
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
	succ, err = server.DB.Exec(sql, "coconut", 300)
	if err != nil {
		panic(err)
	}

	err = server.DB.QueryRow("SELECT * FROM shortener WHERE shortener_id = $1;", succ.LastInsertId).Scan(
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
