package models

// SimpleAlbum represents a Spotify Album Object
type SimpleAlbum struct {
	AlbumGroup           string         `json:"album_group,omitempty"`            //The type of the album: one of "album" , "single" , or "compilation".
	AlbumType            string         `json:"album_type,omitempty"`             //The type of the album: one of “album”, “single”, or “compilation”.
	Artists              []SimpelArtist `json:"artists,omitempty"`                //array of simplified artist objects 	The artists of the album. Each artist object includes a link in href to more detailed information about the artist.
	AvailableMarkets     []string       `json:"available_markets,omitempty"`      //array of strings 	The markets in which the album is available: ISO 3166-1 alpha-2 country codes. Note that an album is considered available in a market when at least 1 of its tracks is available in that market.
	ExternalURLs         []ExternalURL  `json:"external_urls,omitempty"`          //an external URL object 	Known external URLs for this album.
	HREF                 string         `json:"href,omitempty"`                   //A link to the Web API endpoint providing full details of the album.
	ID                   string         `json:"id,omitempty"`                     //The Spotify ID for the album.
	Images               []Image        `json:"images,omitempty"`                 //array of image objects 	The cover art for the album in various sizes, widest first.
	Name                 string         `json:"name,omitempty"`                   //The name of the album. In case of an album takedown, the value may be an empty string.
	ReleaseDate          string         `json:"release_date,omitempty"`           //The date the album was first released, for example 1981. Depending on the precision, it might be shown as 1981-12 or 1981-12-15.
	ReleaseDatePrecision string         `json:"release_date_precision,omitempty"` //The precision with which release_date value is known: year , month , or day.
	Restrictions         Restrictions   `json:"restrictions,omitempty"`           //a restrictions object 	Part of the response when Track Relinking is applied, the original track is not available in the given market, and Spotify did not have any tracks to relink it with. The track response will still contain metadata for the original track, and a restrictions object containing the reason why the track is not available: "restrictions" : {"reason" : "market"}
	ObjectType           string         `json:"type,omitempty"`                   //The object type: “album”
	URI                  string         `json:"uri,omitempty"`                    //The Spotify URI for the album.
}

// SimpleArtist represents a Spotify Artist Object (simple)
type SimpelArtist struct {
	ExternalURLs []ExternalURL `json:"external_urls,omitempty"` //an external URL object 	Known external URLs for this artist.
	HREF         string        `json:"href,omitempty"`          //A link to the Web API endpoint providing full details of the artist.
	ID           string        `json:"id,omitempty"`            //The Spotify ID for the artist.
	Name         string        `json:"name,omitempty"`          //The name of the artist.
	ObjectType   string        `json:"type,omitempty"`          //The object type: "artist"
	URI          string        `json:"uri,omitempty"`           //The Spotify URI for the artist.
}

// ExternalURL represents a Spotify External URL Object (simple)
type ExternalURL struct {
	//TODO: look at to implement
}

// Image represents a Spotify Image Object
type Image struct {
	Height int32  `json:"height"`          //The image height in pixels. If unknown: null or not returned.
	URL    string `json:"url,omitempty"`   //The source URL of the image.
	Width  int32  `json:"width,omitempty"` //The image width in pixels. If unknown: null or not returned.
}

// Restrictions represents a Spotify Restrictions Object
type Restrictions struct {
	Seeds  []RecommendationSeed `json:"seeds,omitempty"`  //An array of recommendation seed objects.
	Tracks []SimpleTrack        `json:"tracks,omitempty"` //An array of track object (simplified) ordered according to the parameters supplied.
}

// Seed represents a Spotify Seeds Object
type RecommendationSeed struct {
	AfterFilteringSize int32  `json:"afterFilteringSize,omitempty"` //The number of tracks available after min_* and max_* filters have been applied.
	AfterRelinkingSize int32  `json:"afterRelinkingSize,omitempty"` //The number of tracks available after relinking for regional availability.
	HREF               string `json:"href,omitempty"`               //A link to the full track or artist data for this seed. For tracks this will be a link to a Track Object. For artists a link to an Artist Object. For genre seeds, this value will be null.
	ID                 string `json:"id,omitempty"`                 //The id used to select this seed. This will be the same as the string used in the seed_artists , seed_tracks or seed_genres parameter.
	InitialPoolSize    int32  `json:"initialPoolSize,omitempty"`    //The number of recommended tracks available for this seed.
	ObjectType         string `json:"type,omitempty"`               //The entity type of this seed. One of artist , track or genre.
}

// Track represents a Spotify Tracks Object
type SimpleTrack struct {
	Artists          []SimpelArtist `json:"artists,omitempty"`           //an array of simple artist objects 	The artists who performed the track. Each artist object includes a link in href to more detailed information about the artist.
	AvailableMarkets []string       `json:"available_markets,omitempty"` //array of strings 	A list of the countries in which the track can be played, identified by their ISO 3166-1 alpha-2 code.
	DiscNumber       int32          `json:"disc_number,omitempty"`       //The disc number (usually 1 unless the album consists of more than one disc).
	Durationms       int32          `json:"duration_ms,omitempty"`       //The track length in milliseconds.
	Explicit         bool           `json:"explicit"`                    //Whether or not the track has explicit lyrics ( true = yes it does; false = no it does not OR unknown).
	ExternalURLs     []ExternalURL  `json:"external_urls,omitempty"`     //an external URL object 	External URLs for this track.
	HREF             string         `json:"href,omitempty"`              //A link to the Web API endpoint providing full details of the track.
	ID               string         `json:"id,omitempty"`                //The Spotify ID for the track.
	IsPlayable       bool           `json:"is_playable"`                 //Part of the response when Track Relinking is applied. If true , the track is playable in the given market. Otherwise false.
	LinkedFrom       LinkedTrack    `json:"linked_from,omitempty"`       //a linked track object 	Part of the response when Track Relinking is applied and is only part of the response if the track linking, in fact, exists. The requested track has been replaced with a different track. The track in the linked_from object contains information about the originally requested track.
	Restrictions     []Restrictions `json:"restrictions,omitempty"`      //a restrictions object 	Part of the response when Track Relinking is applied, the original track is not available in the given market, and Spotify did not have any tracks to relink it with. The track response will still contain metadata for the original track, and a restrictions object containing the reason why the track is not available: "restrictions" : {"reason" : "market"}
	Name             string         `json:"name,omitempty"`              //The name of the track.
	PreviewURL       string         `json:"preview_url,omitempty"`       //A URL to a 30 second preview (MP3 format) of the track.
	TrackNumber      int32          `json:"track_number,omitempty"`      //The number of the track. If an album has several discs, the track number is the number on the specified disc.
	ObjectType       string         `json:"type,omitempty"`              //The object type: “track”.
	URI              string         `json:"uri,omitempty"`               //The Spotify URI for the track.
	IsLocal          bool           `json:"is_local"`                    //Whether or not the track is from a local file.
}

// LinkedTrack represents a Spotify Linked Track Object
type LinkedTrack struct {
	ExternalURLs []ExternalURL `json:"external_urls"`  //an external URL object 	Known external URLs for this track.
	HREF         string        `json:"href,omitempty"` //A link to the Web API endpoint providing full details of the track.
	ID           string        `json:"id,omitempty"`   //The Spotify ID for the track.
	ObjectType   string        `json:"type,omitempty"` //The object type: “track”.
	URI          string        `json:"uri,omitempty"`  //The Spotify URI for the track.
}

// Playlist represents a Spotify Playlist Object
type Playlist struct {
	Collaborative bool            `json:"collaborative,omitempty"` //Returns true if context is not search and the owner allows other users to modify the playlist. Otherwise returns false.
	Description   string          `json:"description,omitempty"`   //The playlist description. Only returned for modified, verified playlists, otherwise null.
	ExternalURLs  []ExternalURL   `json:"external_urls,omitempty"` //an external URL object 	Known external URLs for this playlist.
	Followers     Follower        `json:"followers,omitempty"`     //a followers object 	Information about the followers of the playlist.
	HREF          string          `json:"href,omitempty"`          //A link to the Web API endpoint providing full details of the playlist.
	ID            string          `json:"id,omitempty"`            //The Spotify ID for the playlist.
	Images        []Image         `json:"images,omitempty"`        //an array of image objects 	Images for the playlist. The array may be empty or contain up to three images. The images are returned by size in descending order. See Working with Playlists.Note: If returned, the source URL for the image ( url ) is temporary and will expire in less than a day.
	Name          string          `json:"name,omitempty"`          //The name of the playlist.
	Owner         PublicUser      `json:"owner,omitempty"`         //a public user object 	The user who owns the playlist
	Public        bool            `json:"public,omitempty"`        //or null 	The playlist’s public/private status: true the playlist is public, false the playlist is private, null the playlist status is not relevant. For more about public/private status, see Working with Playlists.
	SnapshotID    string          `json:"snapshot_id,omitemtpy"`   //The version identifier for the current playlist. Can be supplied in other requests to target a specific playlist version: see Remove tracks from a playlist
	Tracks        []PlaylistTrack `json:"tracks,omitempty"`        //array of playlist track objects inside a paging object 	Information about the tracks of the playlist.
	ObjectType    string          `json:"type,omitempty"`          //The object type: “playlist”
	URI           string          `json:"uri,omitempty"`           //The Spotify URI for the playlist.
}

// Follower represents a Spotify Folloer Object
type Follower struct {
	//TODO: implement this? ATM not supported by web API
}

// PublicUser represents a Spotify Public User Object
type PublicUser struct {
	DisplayName  string      `json:"display_name,omitempty"`  //The name displayed on the user’s profile. null if not available.
	ExternalURLs ExternalURL `json:"external_urls,omitempty"` //an external URL object 	Known public external URLs for this user.
	Followers    Follower    `json:"follwers,omitempty"`      //A followers object 	Information about the followers of this user.
	HREF         string      `json:"href,omitempty"`          //A link to the Web API endpoint for this user.
	ID           string      `json:"id,omitempty"`            //The Spotify user ID for this user.
	Images       []Image     `json:"images,omitempty"`        //array of image objects 	The user’s profile image.
	ObjectType   string      `json:"type,omitempty"`          //The object type: “user”
	URI          string      `json:"uri,omitempty"`           //The Spotify URI for this user.
}

// PlaylistTrack represents a Spotify Playlist Track Object
type PlaylistTrack struct {
	AddedAt string `json:"added_at,omitempty"` //a timestamp 	The date and time the track was added.
	//Note that some very old playlists may return null in this field.
	AddedBy PublicUser `json:"added_by,omitemtpy"` //a user object 	The Spotify user who added the track.
	//Note that some very old playlists may return null in this field.
	IsLocal bool  `json:"is_local,omitempty"` //a Boolean 	Whether this track is a local file or not.
	Track   Track `json:"track,omitempty"`    //a track object 	Information about the track.
}

// Track represents a Spotify Track Object
type Track struct {
	Album            SimpleAlbum    `json:"album,omitempty"`             //a simplified album object 	The album on which the track appears. The album object includes a link in href to full information about the album.
	Artists          []SimpelArtist `json:"artists,omitempty"`           //an array of simplified artist objects 	The artists who performed the track. Each artist object includes a link in href to more detailed information about the artist.
	AvailableMarkets []string       `json:"available_markets,omitempty"` //array of strings 	A list of the countries in which the track can be played, identified by their ISO 3166-1 alpha-2 code.
	DiscNumber       int32          `json:"disc_number,omitempty"`       //The disc number (usually 1 unless the album consists of more than one disc).
	DurationMS       int32          `json:"duration_ms,omitempty"`       //The track length in milliseconds.
	Explicit         bool           `json:"explicit,omitempty`           //Whether or not the track has explicit lyrics ( true = yes it does; false = no it does not OR unknown).
	ExternalIDs      []ExternalID   `json:"external_ids,omitempty"`      //an external ID object 	Known external IDs for the track.
	ExternalURLs     []ExternalURL  `json:"external_urls,omitempty"`     //an external URL object 	Known external URLs for this track.
	HREF             string         `json:"href,omitempty"`              //A link to the Web API endpoint providing full details of the track.
	ID               string         `json:"id,omitempty"`                //The Spotify ID for the track.
	IsPlayable       bool           `json:"is_playable`                  //Part of the response when Track Relinking is applied. If true , the track is playable in the given market. Otherwise false.
	LinkedFrom       LinkedTrack    `json:"linked_from,omitempty"`       //a linked track object 	Part of the response when Track Relinking is applied, and the requested track has been replaced with different track. The track in the linked_from object contains information about the originally requested track.
	Restrictions     Restrictions   `json:"restrictions,omitempty"`      //a restrictions object 	Part of the response when Track Relinking is applied, the original track is not available in the given market, and Spotify did not have any tracks to relink it with. The track response will still contain metadata for the original track, and a restrictions object containing the reason why the track is not available: "restrictions" : {"reason" : "market"}
	Name             string         `json:"name,omitempty"`              //The name of the track.
	PreviewURL       string         `json:"preview_url,omitempty"`       //A link to a 30 second preview (MP3 format) of the track. Can be null
	TrackNumber      int32          `json:"track_number,omitempty"`      //The number of the track. If an album has several discs, the track number is the number on the specified disc.
	ObjectType       string         `json:"type,omitempty"`              //The object type: “track”.
	URI              string         `json:"uri,omitempty"`               //The Spotify URI for the track.
	IsLocal          bool           `json:"is_local,omitempty"`          //Whether or not the track is from a local file.
	Popularity       int32          `json:"Popularity,omitempty"`        //The popularity of the track. The value will be between 0 and 100, with 100 being the most popular.
	//The popularity of a track is a value between 0 and 100, with 100 being the most popular. The popularity is calculated by algorithm and is based, in the most part, on the total number of plays the track has had and how recent those plays are.
	//Generally speaking, songs that are being played a lot now will have a higher popularity than songs that were played a lot in the past. Duplicate tracks (e.g. the same track from a single and an album) are rated independently. Artist and album popularity is derived mathematically from track popularity. Note that the popularity value may lag actual popularity by a few days: the value is not updated in real time.
}

// AuthResponse represents a Spotify AccessToken
type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
}

// AuthBody represents a Spotify AuthBody
type AuthBody struct {
	GrantType string `json:"grant_type"`
}
