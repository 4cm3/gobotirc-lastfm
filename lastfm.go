package lastfm

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// RecentTracks represents the recent tracks JSON structure
type RecentTracks struct {
	RecentTracks struct {
		Track []Track `json:"track"`
	} `json:"recenttracks"`
}

type Track struct {
	Name string `json:"name"`
	Artist struct {
		Name string `json:"#text"`
	} `json:"artist"`
}


func (s *Track) String() string {
	if s.Name == "" || s.Artist.Name == "" {
		return fmt.Sprintf("song not found")
	}
	return fmt.Sprintf("%v - %v",s.Name, s.Artist.Name)
}


func GetRecentTrack(apiKey string, username string) (*Track, error) {
	dest := fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&format=json&user=%s&api_key=%s&limit=1", username, apiKey)
	resp, err := http.DefaultClient.Get(dest)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var r RecentTracks
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}

	if len(r.RecentTracks.Track) == 0 {
		return nil, fmt.Errorf("user/song not found")
	}

	return &r.RecentTracks.Track[0], nil
}
