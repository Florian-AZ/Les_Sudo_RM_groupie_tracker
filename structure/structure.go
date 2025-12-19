package structure

type PageData struct {
	LogIn     bool
	Artist    string
	TrackData TrackData
	AlbumData Html_Album
	Track     string
}

type Api_Token struct {
	AccessToken      string `json:"access_token"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// Structure avec les données de la recherche
type Api_Recherche struct {
	Tracks struct {
		Items []Api_Track `json:"items"`
		/*JSON Structure
		tracks
			items[]
				album
					external_urls
					id
					images[1]
					name
					release_date
					total_tracks

				artists[]
					external_urls
					id
					name

				duration_ms
				external_urls
				id
				name
		*/
	} `json:"tracks"`
	Artists struct {
		Items []Api_Artist `json:"items"`
		/*JSON Structure
		artists
			items[]
				external_urls
				folowers
					total

				genres[]
				id
				images[1]
				name
		*/
	} `json:"artists"`
	Albums struct {
		Items []Api_Albums `json:"items"`
		/*JSON Structure
		albums
			items[]
				total_tracks
				external_urls
				id
				images[1]
				name
				release_date
				artists[]
					external_urls
					id
					name
		*/
	} `json:"albums"`
	Error Api_Error `json:"error"`
}

// structure avec les données de l'album que l'on veut récupérer
type Api_AllAlbums struct {
	Items []Api_Albums `json:"items"`
	Error Api_Error    `json:"error"`
	/*JSON Structure
	items[]
	*/
}

type Api_Albums struct {
	TotalTracks int              `json:"total_tracks"`
	URL         Api_ExternalUrls `json:"external_urls"`
	Id          string           `json:"id"`
	Image       []Api_Image      `json:"images"`
	Name        string           `json:"name"`
	ReleaseDate string           `json:"release_date"`
	Artists     []Api_Artist     `json:"artists"`
	/*JSON Structure
	total_tracks
	external_urls
	id
	images[1]
	name
	release_date
	artists[]
		external_urls
		id
		name
	*/
}

type Api_Track struct {
	Album      Api_Track_Album  `json:"album"`
	Artists    []Api_Artist     `json:"artists"`
	DurationMs int              `json:"duration_ms"`
	URL        Api_ExternalUrls `json:"external_urls"`
	Id         string           `json:"id"`
	Name       string           `json:"name"`
	Error      Api_Error        `json:"error"`
	/*JSON Structure
	Tracks
		album
		artists[]
		duration_ms
		external_urls
		id
		name
	*/
}

type Api_Track_Album struct {
	URL         Api_ExternalUrls `json:"external_urls"`
	Id          string           `json:"id"`
	Image       []Api_Image      `json:"images"`
	Name        string           `json:"name"`
	ReleaseDate string           `json:"release_date"`
	TotalTracks int              `json:"total_tracks"`
	/*JSON Structure
	album
		external_urls
		id
		images[1]
		name
		release_date
		total_tracks
	*/
}

type Api_ExternalUrls struct {
	Spotify string `json:"spotify"`
	/*JSON Structure
	external_urls
	    spotify
	*/
}

type Api_Image struct {
	URL string `json:"url"`
	/*JSON Structure
	images[]
		url
	*/
}

type Api_Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Api_Artist struct {
	URL       Api_ExternalUrls `json:"external_urls"`
	Followers Api_Followers    `json:"followers"`
	Genres    []string         `json:"genres"`
	Id        string           `json:"id"`
	Images    []Api_Image      `json:"images"`
	Name      string           `json:"name"`
	/*JSON Structure
	Arists
		external_urls
		folowers
		genres[]
		id
		images[1]
		name
	*/
}

type Api_Followers struct {
	Total int `json:"total"`
	/*JSON Structure
	folowers
		total
	*/
}

// Donnée que l'on envoie au html
type Html_Album struct {
	Data []AlbumData
}
type AlbumData struct {
	Image       string
	Name        string
	ReleaseDate string
	TotalTracks int
	URL         string
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
