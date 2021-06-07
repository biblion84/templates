package main

import (
	"net/http"
	"time"
)

func FetchArtists() []Artist {
	//http.Get("http://api.artists.com")
	//json.Unmarshal()
	time.Sleep(time.Millisecond * 60)
	return nil
}

func failOnError(err error, w http.ResponseWriter) bool {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return true
	}
	return false
}
