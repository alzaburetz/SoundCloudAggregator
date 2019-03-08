package main

import (
	"encoding/json"
	"github.com/alzaburetz/SoundCloud/models"
	"github.com/mikkyang/id3-go"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

//global variables, some not used yet
var (
	mu      sync.Mutex
	cookie  string
	cookies []string
	client  = "?client_id="
	agent   string
)

//struct for encoding credentials
type Cred struct {
	ClientID string `json:"client_id"`
	Agent    string `json:"user_agent"`
}

//get client_id and user_agent(browser) from json file
func getInfo() {
	var credentials Cred
	input, _ := ioutil.ReadFile("credentials.json")
	err := json.Unmarshal(input, &credentials)
	if err != nil {
		log.Fatal(err)
	}
	client += credentials.ClientID
	agent = credentials.Agent
}

//get mil users with a starter id

func changeName(str string) string {
	if strings.Contains(str, "/") {
		slash := strings.Index(str, "/")
		str = str[0:slash-1] + str[slash+1:]
	}
	return str
}

func Download(track models.Track) {
	mu.Lock()
	_ = os.Mkdir(track.Genre+"-genre", 0777)
	req, _ := http.Get(track.DownloadURL + client)
	reader := req.Body
	info, _ := ioutil.ReadAll(reader)
	_ = ioutil.WriteFile(track.Genre+"-genre/"+track.Title+"."+track.OriginalFormat, info, 0666)
	if track.OriginalFormat == "mp3" {
		mp3file, err := id3.Open(track.Genre + "-genre/" + track.Title + "." + track.OriginalFormat)
		if err != nil {
			log.Println(err)
		}
		defer mp3file.Close()
		mp3file.SetTitle(track.Title)
		mp3file.SetGenre(track.Genre)
		mp3file.SetArtist(track.Creator.Username)
		mp3file.SetYear(strconv.Itoa(track.ReleaseYear))
	}
	mu.Unlock()
}

func Playlist(dir string) {
	file, _ := os.OpenFile(dir+".pls", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	files, _ := ioutil.ReadDir(dir)
	file.WriteString("[playlist]\nNumberOfEntries=" + strconv.Itoa(len(files)) + "\n")
	defer file.Close()
	for key, f := range files {
		filep, _ := filepath.Abs(dir + "/" + f.Name())

		file.WriteString("File" + strconv.Itoa(key+1) + "=" + filep + "\nTitle" + strconv.Itoa(key+1) + "=" + f.Name() + "\n")
	}

}

func main() {
	//fetch cookie. 20 for starters
	getInfo()
	for i := 0; i < 20; i++ {
		c, _ := http.Get("https://soundcloud.com")
		c.Header.Set("User-Agent", agent)
		cookie = c.Header.Get("Set-Cookie")
		cookies = append(cookies, cookie)
	}
	os.Setenv("CLIENT", client)
	os.Setenv("AGENT", agent)
	var i int64
	for i = 0; i < 50000; i++ {
		fetchTrack(i)
	}
	directories, _ := ioutil.ReadDir("./")
	for _, directory := range directories {
		matched, _ := regexp.Match("-genre", []byte(directory.Name()))
		if matched {
			Playlist(directory.Name())
		}
	}

}
