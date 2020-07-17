package servershortener

import (
	"fmt"
	"strings"

	modelshortener "github.com/water25234/golang-shorturl/app/model/shortener"
)

const (
	encodeParameter = "Flpvf70CsakVjqgeWUPXQxSyJizmNH6B1u3b8cAEKwTd54nRtZOMDhoG2YLrI"
	length          = int64(len(encodeParameter))
)

type ShortenerEncode struct {
	URLEncode string `json:"url_encode"`
	URL       string `json:"url"`
}

func SaveShortenerURL(Url string) interface{} {

	var num int64

	getLastShortenerID, err := modelshortener.GetLastShortenerID()

	num = getLastShortenerID.ShortenerID + 1000

	shortenerEncodeNum := Encode(num)

	fmt.Println(shortenerEncodeNum)

	if Url == "" {
		panic("SaveShortenerURL : url is null")
	}

	saveShortenerURL, err := modelshortener.SaveShortenerUrl(shortenerEncodeNum, Url)

	if err != nil {
		panic(err)
	}

	fmt.Println(saveShortenerURL)

	return saveShortenerURL
}

func Encode(n int64) string {
	if n == 0 {
		return string(encodeParameter[0])
	}

	s := ""
	for ; n > 0; n = n / length {
		s = string(encodeParameter[n%length]) + s
	}
	return s
}

func Decode(key string) (int64, error) {
	var n int64
	for _, c := range []byte(key) {
		i := strings.IndexByte(encodeParameter, c)
		if i < 0 {
			return 0, fmt.Errorf("unexpected character %c in base62 literal", c)
		}
		n = length*n + int64(i)
	}
	return n, nil
}
