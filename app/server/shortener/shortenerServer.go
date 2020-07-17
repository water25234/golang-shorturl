package servershortener

import (
	modelshortener "github.com/water25234/golang-shorturl/app/model/shortener"
)

func SaveShortenerUrl() interface{} {
	item, err := modelshortener.SaveShortenerUrl()

	if err != nil {
		panic(err)
	}

	return item
}
