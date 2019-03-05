package main

import (
	"./models"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

//global variables, some not used yet
var (
	mu      sync.Mutex
	cookie  string
	cookies []string
	client = "?client_id=Zju0N4zmkkZ60kMBebep9nP703ozMTpx"
	agent string
)

//struct for encoding credentials
type Cred struct {
	clientID string `json:"client_id"`
	Agent string `json:"user_agent"`
}

//get client_id and user_agent(browser) from json file
func getInfo() {
	var credentials Cred
	input, _ := ioutil.ReadFile("credentials.json")
	json.Unmarshal(input,&credentials)
	client += credentials.clientID
	agent = credentials.Agent
}

//get mil users with a starter id
func getUsers(n int64) {
	var user models.SCUser
	for i := n; i < n+999999; i++ {
		c := http.Client{
			Timeout: 10 * time.Second,
		}
		req, _ := c.Get("https://api.soundcloud.com/users/" + strconv.FormatInt(i, 10) + client)

		resp := req.Body
		defer func() {
			err := resp.Close()
			if err != nil {
				log.Println(err.Error())
			}
		}()

		if req.StatusCode == 200 || req.StatusCode == 404 {
			if req.StatusCode == 404 {
				continue
			} else {
				log.Println(i)
			}
		} else {
			log.Fatal("Blocked\n")
		}
		bytes, _ := ioutil.ReadAll(resp)
		json.Unmarshal(bytes,&user)
		create(&user,i)
		log.Println(string(bytes))
	}
}

//creates file with user information
func create(sc *models.SCUser, i int64) {
	mu.Lock()
	data, _ := json.Marshal(sc)
	err := ioutil.WriteFile(strconv.FormatInt(i, 10)+".json", data, 0666)
	if err != nil {
		log.Printf("%v", err)
	}
	mu.Unlock()
}

//get track by id
func getTrack(id int64) {
	var track models.Track
	req, er := http.Get("https://api.soundcloud.com/tracks/" + strconv.FormatInt(id, 10) + client)
	if er != nil {
		log.Println("Oops")
	}
	//set header to our agent and random cookie
	req.Header.Set("User-Agent", agent)
	req.Header.Set("Cookie",  cookies[rand.Intn(19)])
	if req.StatusCode != 404 {
		bytess, err := ioutil.ReadAll(req.Body)
		json.Unmarshal(bytess, &track)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(track.Genre)
		if track.Downloadable {
			_ = os.Mkdir(track.Genre + "-genre", 0777)
			req, _ := http.Get(track.DownloadURL + client)
			reader := req.Body
			info, _  := ioutil.ReadAll(reader)
			_ = ioutil.WriteFile(track.Genre + "-genre/" + track.Title + ".mp3", info, 0666)
		}
	}

}


func main() {
	//fetch cookie. 20 for starters
	mu.Lock()
	for i := 0; i < 20; i++ {
		c, _ := http.Get("https://soundcloud.com")
		c.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")
		cookie = c.Header.Get("Set-Cookie")
		cookies = append(cookies,cookie)
	}
	mu.Unlock()
	getInfo()
	var i int64
	for i = 0; i < 1000; i++ {
		getTrack(i)
	}
}
