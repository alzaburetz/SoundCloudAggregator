package main

import (
	"encoding/json"
	"github.com/alzaburetz/SoundCloud/models"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func getUser(n int64) (models.SCUser, error) {
	var user models.SCUser
	var err error
	c := http.Client{
		Timeout: 10 * time.Second,
	}
	req, _ := c.Get("https://api.soundcloud.com/users/" + strconv.FormatInt(n, 10) + client)

	resp := req.Body
	defer func() {
		err := resp.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}()

	if req.StatusCode == 200 || req.StatusCode == 404 {
		if req.StatusCode == 404 {
			err = new(url.Error)
		}
	} else {
		log.Fatal("Blocked\n")
	}
	bytes, _ := ioutil.ReadAll(resp)
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		err = new(json.SyntaxError)
	}
	log.Println(string(bytes))
	return user, err
}
