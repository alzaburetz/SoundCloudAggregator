package models

type Track struct {
	Kind              string  `json:"kind"`
	ID                int     `json:"id"`
	CreatedAt         string  `json:"created_at"`
	User              int     `json:"user_id"`
	Duration          int     `json:"duration"`
	Commentable       bool    `json:"commentable"`
	State             string  `json:"state"`
	OriginalSize      int     `json:"original_content_size"`
	LastModified      string  `json:"last_modified"`
	Sharing           string  `json:"sharing"`
	TagList           string  `json:"tag_list"`
	Permalink         string  `json:"permalink"`
	Streamable        bool    `json:"streamable"`
	EmbeddableBy      string  `json:"embeddable_by"`
	PurchaseUrl       string  `json:"purchase_url"`
	PurchaseTitle     string  `json:"purchase_title"`
	LabelID           int     `json:"label_id"`
	Genre             string  `json:"genre"`
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	LabelName         string  `json:"label_name"`
	Release           string  `json:"release"`
	TrackType         string  `json:"track_type"`
	KeySignature      string  `json:"key_signature"`
	ISRC              string  `json:"isrc"`
	VideoUrl          string  `json:"video_url"`
	BPM               float32 `json:"bpm"`
	ReleaseYear       int     `json:"release_year"`
	ReleaseMounth     int     `json:"release_month"`
	ReleaseDay        int     `json:"release_day"`
	OriginalFormat    string  `json:"original_format"`
	License           string  `json:"license"`
	URI               string  `json:"uri"`
	Creator           User    `json:"user"`
	PermalinkURL      string  `json:"permalink_url"`
	ArtworkURL        string  `json:"artwork_url"`
	StreamURL         string  `json:"stream_url"`
	DownloadURL       string  `json:"download_url"`
	PlaybackCount     int     `json:"playback_count"`
	DownloadCount     int     `json:"download_count"`
	FavoritingsCount  int     `json:"favoritings_count"`
	RepostsCount      int     `json:"reposts_count"`
	CommentCount      int     `json:"comment_count"`
	Downloadable      bool    `json:"downloadable"`
	WaveformURL       string  `json:"waveform_url"`
	AttachmentsURI    string  `json:"attachments_uri"`
	Policy            string  `json:"policy"`
	MonetizationModel string  `json:"monetization_model"`
}

type User struct {
	ID           int    `json:"id"`
	Kind         string `json:"kind"`
	Permalink    string `json:"permalink"`
	Username     string `json:"username"`
	LastModified string `json:"last_modified"`
	URI          string `json:"uri"`
	PermalinkURI string `json:"permalink_uri"`
	AvatarURL    string `json:"avatar_url"`
}
