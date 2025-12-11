package structure

type Token struct {
	AccessToken      string `json:"access_token"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// structure avec les données de l'album que l'on veut récupérer
type AllAlbums struct {
	AlbumItems       []Items `json:"items"`
	Error            string  `json:"error"`
	ErrorDescription string  `json:"error_description"`
}

type Items struct {
	TotalTracks int         `json:"total_tracks"`
	URL         ExternalURL `json:"external_urls"`
	Image       []Image     `json:"images"`
	Name        string      `json:"name"`
	ReleaseDate string      `json:"release_date"`
}

type ExternalURL struct {
	Spotify string `json:"spotify"`
}

type Image struct {
	URL string `json:"url"`
}

// Donnée que l'on envoie au html
type AlbumData struct {
	Data []Data
}
type Data struct {
	Image       string
	Name        string
	ReleaseDate string
	TotalTracks int
	URL         string
}

type Track struct {
	Name    string   `json:"name"`
	Album   Album    `json:"album"`
	Artists []Artist `json:"artists"`
	Error   Error    `json:"error"`
}

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Album struct {
	Name        string      `json:"name"`
	URL         ExternalURL `json:"external_urls"`
	Image       []Image     `json:"images"`
	ReleaseDate string      `json:"release_date"`
}

type Artist struct {
	Name string `json:"name"`
}

// Data de la track que l'on envoie au HTML
type TrackData struct {
	TrackName    string
	AlbumName    string
	AlbumRelease string
	AlbumURL     string
	AlbumImage   string
	ArtistName   string
}
