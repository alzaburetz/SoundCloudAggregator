package models

type SCUser struct {
	ID            int    `json:"id"`
	Kind          string `json:"kind"`
	Perm          string `json:"permalink"`
	Username      string `json:"username"`
	LModified     string `json:"last_modified"`
	URI           string `json:"uri"`
	PermURL       string `json:"permalink_url"`
	Avatar        string `json:"avatar_url"`
	Country       string `json:"country"`
	Fname         string `json:"first_name"`
	Lname         string `json:"last_name"`
	Fullname      string `json:"full_name"`
	Desc          string `json:"description"`
	City          string `json:"city"`
	Dname         string `json:"discogs_name"`
	Myspace       string `json:"myspace_name"`
	Website       string `json:"website"`
	WebsiteT      string `json:"website_title"`
	Tracks        int    `json:"track_count"`
	Playlists     int    `json:"playlist_count"`
	Online        bool   `json:"online"`
	Plan          string `json:"plan"`
	PFavourites   int    `json:"public_favorites_count"`
	Followers     int    `json:"followers_count"`
	Followings    int    `json:"followings_count"`
	Subscriptions []Subs `json:"subscriptions"`
	Likes         int    `json:"likes_count"`
	Reposts       int    `json:"reposts_count"`
	Comments      int    `json:"comments_count"`
}

type Subs struct {
	Product Prod `json:"product"`
}

type Prod struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
