package main

import (
	"encoding/json"
	"github.com/alzaburetz/SoundCloud/models"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func getTrack(id int64) (models.Track, error) {
	var track models.Track
	req, err := http.Get("https://api.soundcloud.com/tracks/" + strconv.FormatInt(id, 10) + os.Getenv("CLIENT"))
	if err != nil {
		log.Println(err)
	}
	if req.StatusCode != 404 {
		bytess, err := ioutil.ReadAll(req.Body)
		json.Unmarshal(bytess, &track)
		if err != nil {
			log.Fatal(err)
		}
		return track, nil
	} else {
		return track, new(url.Error)
	}

}
