package main

import (
	"database/sql"
	"encoding/json"
	"github.com/alzaburetz/SoundCloud/models"
	"log"
)

func fetchTrack(id int64) {
	db, err := sql.Open("sqlite3", "./database.db3")
	if err != nil {
		log.Println(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	var track models.Track

	if err != nil {
		log.Println(err)
	} else {
		var user models.User
		user = track.Creator
		byteuser, _ := json.Marshal(user)
		_, err := db.Exec(`INSERT INTO tracks(KIND, 
						ID, 
					  	CREATED_AT, 
					  	USER_ID, 
					  	DURATION, 
					  	COMMENTABLE, 
					  	STATE, 
					  	ORIGINALSIZE, 
					  	LASTMODIFIED, 
					  	SHARING, 
					  	TEGLIST, 
					  	PERMALINK, 
-- 					  	STREAMABLE, 
					  	EMBEDDABLEBY, 
					  	PURCHASETITLE, 
					  	LABELID, 
					  	GENRE, 
					  	TITLE, 
					  	DESCRIPTION, 
					  	LABELNAME, 
					  	RELEASETRACK, 
					  	TRACKTYPE, 
					  	KEY_SIGNATURE, 
					  	ISRC, 
					  	VIDEOURL, 
					  	BPM, 
					  	RELEASE_YEAR, 
					  	RELEASE_MONTH, 
					  	RELEASE_DAY, 
					  	ORIGINAL_FORMAT, 
					  	LICENSE, 
					  	URI, 
					  	CREATOR, 
					  	PERMALINKURL, 
					  	ARTWORKURL, 
					  	STREAMURL, 
					  	DOWNLOADURL, 
					  	PLAYBACK, 
					  	DOWNLOADS, 
					  	FAVORITINGS, 
					  	REPOSTS, 
					  	COMMENTS, 
					  	DOWNLOADABLE, 
					  	WAVEFORM, 
					  	ATTACHMENTSURI, 
					  	POLICY, 
					  	MONETIZATION) 
						VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
			&track.Kind, &track.ID, &track.CreatedAt, &track.User, &track.Duration, &track.Commentable,
			&track.State, &track.OriginalSize, &track.LastModified, &track.Sharing, &track.TagList,
			&track.Permalink, &track.EmbeddableBy, &track.PurchaseTitle, &track.LabelID,
			&track.Genre, &track.Title, &track.Description, &track.LabelName, &track.Release, &track.TrackType,
			&track.KeySignature, &track.ISRC, &track.VideoUrl, &track.BPM, &track.ReleaseYear, &track.ReleaseMounth,
			&track.ReleaseDay, &track.OriginalFormat, &track.License, &track.URI, string(byteuser), &track.PermalinkURL,
			&track.ArtworkURL, &track.StreamURL, &track.DownloadURL, &track.PlaybackCount, &track.DownloadCount, &track.FavoritingsCount,
			&track.RepostsCount, &track.CommentCount, &track.Downloadable, &track.WaveformURL, &track.AttachmentsURI, &track.Policy, &track.MonetizationModel)
		if err != nil {
			log.Println(err)
		}

	}

}
