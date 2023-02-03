package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func TakeCards() []Artists {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var artist []Artists

	json.Unmarshal(body, &artist)

	return artist
}

func TakeArtist(s int) Artists {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var artist []Artists
	json.Unmarshal(body, &artist)
	return artist[s-1]
}

func TakeConcert(s int) Relations {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + strconv.Itoa(s))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var concert Relations
	json.Unmarshal(body, &concert)
	return concert
}
