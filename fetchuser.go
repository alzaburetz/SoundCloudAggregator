package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func fetchuser(id int64) {
	user, err := getUser(id)
	if err != nil {
		log.Println(err)
	} else {
		mu.Lock()
		db, err := sql.Open("sqlite3", "./database.db3")
		defer func() {
			err := db.Close()
			if err != nil {
				log.Println(err)
			}
		}()
		if err != nil {
			log.Println(err)
		}

		_, err = db.Exec(`INSERT INTO users(id, 
                  kind, 
                  permalinl, 
                  username, 
                  last_modified, 
                  uri, 
                  permurl, 
                  avatar, 
                  country, 
                  first_name, 
                  last_name, 
                  full_name, 
                  "description ", 
                  city, 
                  discogs_name, 
                  myspace, 
                  website, 
                  website_title, 
                  tracks, 
                  playlists,
                  plan_, 
                  public_favorites, 
                  followers, 
                  followings, 
                  subscriptions, 
                  likes, 
                  reposts, 
                  comments) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
			&user.ID, &user.Kind, &user.Perm, &user.LModified, &user.URI, &user.PermURL,
			&user.Avatar, &user.Country, &user.Fname, &user.Lname, &user.Fullname, &user.Desc,
			&user.City, &user.Dname, &user.Myspace, &user.Website, &user.WebsiteT, &user.Tracks,
			&user.Playlists, &user.Plan, &user.PFavourites, &user.Followers, &user.Followings,
			&user.Subscriptions, &user.Likes, &user.Reposts, &user.Comments)
		if err != nil {
			log.Println(err)
		}
		mu.Unlock()
	}

}
