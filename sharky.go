// This file was auto-generated from the
// Grooveshark API Extractor
// package sharky
package sharky

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
	"unicode"
)

// Get your own KEY and SECTER here http://developers.grooveshark.com/api
const API_HOST = "api.grooveshark.com"
const API_ENDPOIT = "/ws3.php"
const SIG_GET_KEY = "?sig="
const HTTPS = "https://"
const HTTP = "http://"
const CONTENT_TYPE = "application/json;charset=utf-8"

type RequestData struct {
	Method     string                 `json:"method"`
	Parameters map[string]interface{} `json:"parameters"`
	Header     map[string]string      `json:"header"`
}

type Response struct {
	Header map[string]string        `json:"header"`
	Result map[string]interface{}   `json:"result"`
	Errors []map[string]interface{} `json:"errors"`
}

type Country struct {
	ID  int64
	CC1 int64
	CC2 int64
	CC3 int64
	CC4 int64
	DMA int64
	IPR int64
}

func getCountryElem(country *Country) reflect.Value {
	return reflect.ValueOf(country).Elem()
}

type AutoplayState struct {
	// TODO fill
}

type Song struct {
	SongID                int64
	SongName              string
	ArtistID              int64
	ArtistName            string
	AlbumID               int64
	AlbumName             string
	CoverArtFilename      string
	Popularity            string
	IsLowBitrateAvailable bool
	IsVerified            bool
	Flags                 int64
	TSFavorited           string
}

func getSongElem(song *Song) reflect.Value {
	return reflect.ValueOf(song).Elem()
}

type LibSong struct {
	SongID   int64
	AlbumID  int64
	ArtistID int64
	TrackID  int64
}

func getLibSongElem(libSong *LibSong) reflect.Value {
	return reflect.ValueOf(libSong).Elem()
}

type Playlist struct {
	PlaylistID   int64
	PlaylistName string
	TSAdded      string
}

func getPlaylistElem(playlist *Playlist) reflect.Value {
	return reflect.ValueOf(playlist).Elem()
}

type UserInfo struct {
	UserID     int64
	Email      string
	FName      string
	LName      string
	IsPlus     bool
	IsAnywhere bool
	IsPremium  bool
}

func getUserInfoElem(userInfo *UserInfo) reflect.Value {
	return reflect.ValueOf(userInfo).Elem()
}

type UserSubscriptionInfo struct {
	// TODO fill
}

type PlaylistInfo struct {
	PlaylistName        string
	TSModified          string
	UserID              int64
	PlaylistDescription string
	CoverArtFilename    string
}

func getPlaylistInfoElem(plInfo *PlaylistInfo) reflect.Value {
	return reflect.ValueOf(plInfo).Elem()
}

type ServiceDescription struct {
	// TODO fill
}

type AlbumInfo struct {
	// TODO fill
}

type ArtistInfo struct {
	// TODO fill
}

type SongInfo struct {
	// TODO fill
}

type Album struct {
	// TODO fill
}

type Artist struct {
	// TODO fill
}

type StreamDetails struct {
	StreamKey      string
	Url            string
	StreamServerID int64
	USecs          int64
}

func getStreamDetailsElem(streamDetails *StreamDetails) reflect.Value {
	return reflect.ValueOf(streamDetails).Elem()
}

type SongUrl struct {
	// TODO fill
}

type PlaylistUrl struct {
	// TODO fill
}

type TinysongUrl struct {
	// TODO fill
}

type TrialInfo struct {
	// TODO fill
}

type StreamKey struct {
	// TODO fill
}

type Tag struct {
	// TODO fill
}

// end structs definitions

func New(key, secret string) *Sharky {
	return new(Sharky).Init(key, secret)
}

// ######################  Sharky's methods  ######################
type Sharky struct {
	SessionID string
	Key       string
	Secret    string
	Username  string
	Password  string
	UserInfo  *UserInfo
}

// Initializes Sharky with key and secret needed for communication with
// GS API and username and password, needed for some specific functionality
func (sharky *Sharky) Init(key, secret string) *Sharky {
	sharky.Key = key
	sharky.Secret = secret

	return sharky
}

func (sharky *Sharky) NoSessionCallHttp(method string, params map[string]interface{}) map[string]interface{} {
	return makeCall(method, params, "", HTTP, sharky.Key, sharky.Secret)
}

func (sharky *Sharky) NoSessionCallHttps(method string, params map[string]interface{}) map[string]interface{} {
	return makeCall(method, params, "", HTTPS, sharky.Key, sharky.Secret)
}

func (sharky *Sharky) SessionCallHttps(method string, params map[string]interface{}) map[string]interface{} {
	return makeCall(method, params, sharky.SessionID, HTTPS, sharky.Key, sharky.Secret)
}

func (sharky *Sharky) SessionCallHttp(method string, params map[string]interface{}) map[string]interface{} {
	return makeCall(method, params, sharky.SessionID, HTTP, sharky.Key, sharky.Secret)
}

// Use addUserLibrarySongsEx instead. Add songs to a user's library.
// Song metadata should be spread across all 3 params. albumIDs[0] should
// be the respective albumID for songIDs[0] and same with artistIDs.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) AddUserLibrarySongs(songIDs, albumIDs, artistIDs string) {
	log.Fatal("Use AddUserLibrarySongsEx instead.")
}

// Get user library songs. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetUserLibrarySongs(limit, page int) []*Song {
	if page <= 0 {
		return nil
	}
	return sharky.getSongs(limit, "getUserLibrarySongs")
}

// Add songs to a user's library. Songs should be an array of objects
// representing each song with keys: songID, albumID, artistID, trackNum.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) AddUserLibrarySongsEx(songs []LibSong) {
	// TODO check it
	params := make(map[string]interface{})
	params["songs"] = songs

	sharky.SessionCallHttp("addUserLibrarySongsEx", params)
}

// Remove songs from a user's library.
// Returns true if everything is OK.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) RemoveUserLibrarySongs(songs []LibSong) bool {
	// TODO check it
	params := make(map[string]interface{})
	params["songs"] = songs

	result := sharky.SessionCallHttp("removeUserLibrarySongs", params)

	if suc, ok := result["success"].(bool); ok {
		return suc
	}

	return false
}

// Get subscribed playlists of the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetUserPlaylistsSubscribed() []*Playlist {
	result := sharky.SessionCallHttp("getUserPlaylistsSubscribed", nil)
	return sharky.getPlaylists(result)
}

// Get playlists of the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetUserPlaylists(limit int) []*Playlist {
	params := make(map[string]interface{})
	params["limit"] = limit

	result := sharky.SessionCallHttp("getUserPlaylists", params)
	return sharky.getPlaylists(result)
}

func (sharky *Sharky) getPlaylists(result map[string]interface{}) []*Playlist {
	if playlists, ok := result["playlists"].([]interface{}); ok {
		plArr := make([]*Playlist, 0)
		for _, plParams := range playlists {
			if plParam, ok := plParams.(map[string]interface{}); ok {
				playlist := new(Playlist)
				elem := getPlaylistElem(playlist)
				mapToStruct(plParam, &elem)
				plArr = append(plArr, playlist)
			}
		}

		return plArr
	}
	return nil
}

// Get user favorite songs. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetUserFavoriteSongs(limit int) []*Song {
	return sharky.getSongs(limit, "getUserFavoriteSongs")
}

// Remove a set of favorite songs for a user. Must provide a logged-in sessionID.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) RemoveUserFavoriteSongs(songIDs string) {
	log.Fatal("Not implemented due to non-clean spec")
}

// Logout a user using an established session.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) Logout() {
	sharky.SessionCallHttp("logout", nil)
	sharky.UserInfo = nil
	sharky.Username = ""
	sharky.Password = ""
	sharky.SessionID = ""
	sharky.Key = ""
	sharky.Secret = ""
}

// Authenticate a user using a token from http://grooveshark.com/auth/.
// See Overview for documentation.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) AuthenticateToken(token string) {
	params := make(map[string]interface{})
	params["token"] = token

	result := sharky.SessionCallHttps("authenticateToken", params)

	sharky.auth(result)
}

// Get logged-in user info from sessionID
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetUserInfo() *UserInfo {
	if sharky.UserInfo == nil {
		result := sharky.SessionCallHttp("getUserInfo", nil)
		userInfo := new(UserInfo)
		elem := getUserInfoElem(userInfo)
		mapToStruct(result, &elem)
		sharky.UserInfo = userInfo
	}
	return sharky.UserInfo
}

// Get logged-in user subscription info. Returns type of subscription
// and either dateEnd or recurring.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetUserSubscriptionDetails() *UserSubscriptionInfo {
	log.Fatal("Not impelemented: Service does not have access to this method.")
	return nil
}

// Add a favorite song for a user. Must provide a logged-in sessionID.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) AddUserFavoriteSong(songID int) {
	params := make(map[string]interface{})
	params["songID"] = songID

	result := sharky.SessionCallHttp("addUserFavoriteSong", params)

	logMsg(result, "Song added to Favorites successfully.", "Error adding song to Favorites.")
}

// Subscribe to a playlist for the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) SubscribePlaylist(playlistID int) {
	params := make(map[string]interface{})
	params["playlistID"] = playlistID
	result := sharky.SessionCallHttp("subscribePlaylist", params)

	sucMsg := "Subscribtion to playlist finished successfully"
	errMsg := "Cannot subscribe to playlist"
	logMsg(result, sucMsg, errMsg)
}

// Unsubscribe from a playlist for the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) UnsubscribePlaylist(playlistID int) {
	params := make(map[string]interface{})
	params["playlistID"] = playlistID
	result := sharky.SessionCallHttp("unsubscribePlaylist", params)

	sucMsg := "Unsubscribed from playlist."
	errMsg := "Cannot upsubscribe from playlist."
	logMsg(result, sucMsg, errMsg)
	// TODO impelemnt
}

// Get country from IP. If an IP is omitted, it will use the request's IP.
func (sharky *Sharky) GetCountry(ip string) *Country {
	params := make(map[string]interface{})
	params["ip"] = ip
	result := sharky.SessionCallHttp("getCountry", params)

	country := new(Country)
	elem := getCountryElem(country)
	mapToStruct(result, &elem)

	return country
}

// Get playlist information. To get songs as well, call getPlaylist.
func (sharky *Sharky) GetPlaylistInfo(playlistID string) *PlaylistInfo {
	params := make(map[string]interface{})
	params["playlistID"] = playlistID
	result := sharky.SessionCallHttp("getPlaylistInfo", params)

	playlistInfo := new(PlaylistInfo)
	elem := getPlaylistInfoElem(playlistInfo)
	mapToStruct(result, &elem)

	return playlistInfo
}

// Get a subset of today's popular songs, from the Grooveshark popular billboard.
func (sharky *Sharky) GetPopularSongsToday(limit int) []*Song {
	return sharky.getSongs(limit, "getPopularSongsToday")
}

// Get a subset of this month's popular songs, from the Grooveshark popular billboard.
func (sharky *Sharky) GetPopularSongsMonth(limit int) []*Song {
	return sharky.getSongs(limit, "getPopularSongsMonth")
}

// Get a subset of this month's popular songs, from the Grooveshark popular billboard.
func (sharky *Sharky) getSongs(limit int, method string) []*Song {
	if limit <= 0 {
		return nil
	}
	params := make(map[string]interface{})
	params["limit"] = limit
	result := sharky.SessionCallHttp(method, params)

	if songs, ok := result["songs"].([]interface{}); ok {
		songArr := make([]*Song, 0)
		for _, songParams := range songs {
			if sParams, ok := songParams.(map[string]interface{}); ok {
				song := new(Song)
				elem := getSongElem(song)
				mapToStruct(sParams, &elem)
				songArr = append(songArr, song)
			}
		}

		return songArr
	}

	return nil
}

// Useful for testing if the service is up. Returns "Hello, World" in various languages.
func (sharky *Sharky) PingService() string {
	// TODO impelemnt
	// http.Get()
	return ""
}

// Describe service methods
func (sharky *Sharky) GetServiceDescription() *ServiceDescription {
	// TODO impelemnt
	return nil
}

// Undeletes a playlist.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) UndeletePlaylist(playlistID int) {
	// TODO impelemnt
}

// Deletes a playlist.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) DeletePlaylist(playlistID int) {
	// TODO impelemnt
}

// Get songs on a playlist. Use getPlaylist instead.
func (sharky *Sharky) GetPlaylistSongs(playlistID string, limit int) []Song {
	// TODO impelemnt
	return nil
}

// Get playlist info and songs.
func (sharky *Sharky) GetPlaylist(playlistID string, limit int) *Playlist {
	// TODO impelemnt
	return nil
}

// Set playlist songs, overwrites any already saved
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) SetPlaylistSongs(playlistID int, songIDs string) {
	// TODO impelemnt
}

// Create a new playlist, optionally adding songs to it.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) CreatePlaylist(name, songIDs string) {
	// TODO impelemnt
}

// Renames a playlist.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) RenamePlaylist(playlistID int, name string) {
	// TODO impelemnt
}

// Authenticate a user using an established session, a login and an md5 of their password.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) Authenticate(login, password string) {
	params := make(map[string]interface{})
	params["login"] = login
	params["password"] = md5sum(password)

	result := sharky.SessionCallHttps("authenticate", params)

	sharky.auth(result)
}

func (sharky *Sharky) auth(result map[string]interface{}) {
	if suc, ok := result["success"].(bool); ok {
		if suc {
			log.Println("Authentication successful.")
		} else {
			log.Fatal("Invalid credentials. Check your username and password then try again.")
		}
	}
	userInfo := new(UserInfo)
	elem := getUserInfoElem(userInfo)
	mapToStruct(result, &elem)

	sharky.UserInfo = userInfo
}

// Get userID from username
func (sharky *Sharky) GetUserIDFromUsername(username string) string {
	// TODO impelemnt
	return ""
}

// Get meta-data information about one or more albums
func (sharky *Sharky) GetAlbumsInfo(albumIDs string) []AlbumInfo {
	// TODO impelemnt
	return nil
}

// Get songs on an album. Returns all songs, verified and unverified
func (sharky *Sharky) GetAlbumSongs(albumID, limit int) []Song {
	// TODO impelemnt
	return nil
}

// Get meta-data information about one or more artists
func (sharky *Sharky) GetArtistsInfo(artistIDs string) []ArtistInfo {
	// TODO impelemnt
	return nil
}

// Get information about a song or multiple songs.
// The songID(s) should always be passed in as an array.
func (sharky *Sharky) GetSongsInfo(songIDs string) []SongInfo {
	// TODO impelemnt
	return nil
}

// Check if an album exists
func (sharky *Sharky) GetDoesAlbumExist(albumID int) bool {
	// TODO impelemnt
	return false
}

// Check if a song exists
func (sharky *Sharky) GetDoesSongExist(songID int) bool {
	// TODO impelemnt
	return false
}

// Check if an artist exists
func (sharky *Sharky) GetDoesArtistExist(artistID int) bool {
	// TODO impelemnt
	return false
}

// Authenticate a user (login) using an established session.
// Please use the authenticate method instead.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) AuthenticateUser(username, token string) {
	// TODO impelemnt
}

// Get an artist's verified albums
func (sharky *Sharky) GetArtistVerifiedAlbums(artistID int) []Album {
	// TODO impelemnt
	return nil
}

// Get an artist's albums, verified and unverified
func (sharky *Sharky) GetArtistAlbums(artistID int) []Album {
	// TODO impelemnt
	return nil
}

// Get 100 popular songs for an artist
func (sharky *Sharky) GetArtistPopularSongs(artistID int) []Song {
	// TODO impelemnt
	return nil
}

// ================= Search =================

// Perform a playlist search.
func (sharky *Sharky) GetPlaylistSearchResults(query string, limit int) []Playlist {
	// TODO impelemnt
	return nil
}

// Perform an album search.
func (sharky *Sharky) GetAlbumSearchResults(query string, limit int) []Album {
	// TODO impelemnt
	return nil
}

// Perform a song search.
func (sharky *Sharky) GetSongSearchResults(query string, country Country, limit, offset int) []Song {
	// TODO impelemnt
	return nil
}

// Perform an artist search.
func (sharky *Sharky) GetArtistSearchResults(query string, limit int) []Artist {
	// TODO impelemnt
	return nil
}

// ================= Streams =================

// Get stream key, ID, etc. from songID. Requires country object obtained from getCountry
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetStreamKeyStreamServer(songID int64, country *Country, lowBitrate bool) *StreamDetails {
	params := make(map[string]interface{})
	params["songID"] = songID
	params["country"] = country
	params["lowBitrate"] = lowBitrate

	result := sharky.SessionCallHttp("getStreamKeyStreamServer", params)

	streamDetails := new(StreamDetails)
	elem := getStreamDetailsElem(streamDetails)
	mapToStruct(result, &elem)

	return streamDetails
}

// ================= URLS =================

// Get Grooveshark URL for tinysong base 62.
func (sharky *Sharky) GetSongURLFromTinysongBase62(base62 string) *SongUrl {
	// TODO impelemnt
	return nil
}

// Get playable song URL from songID
func (sharky *Sharky) GetSongURLFromSongID(songID int) *SongUrl {
	// TODO impelemnt
	return nil
}

// Get playlist URL from playlistID
func (sharky *Sharky) GetPlaylistURLFromPlaylistID(playlistID int) *PlaylistUrl {
	// TODO impelemnt
	return nil
}

// Get a song's Tinysong.com url.
func (sharky *Sharky) GetTinysongURLFromSongID(songID int) *TinysongUrl {
	// TODO impelemnt
	return nil
}

// ================= Users (no auth) =================

// Get playlists created by a userID. Does not require an authenticated session.
func (sharky *Sharky) GetUserPlaylistsByUserID(userID, limit int) []Playlist {
	// TODO impelemnt
	return nil
}

// Get user info from userID
func (sharky *Sharky) GetUserInfoFromUserID(userID int) *UserInfo {
	// TODO impelemnt
	return nil
}

// ================= Recs =================

// Get similar artist for a given artistID.
func (sharky *Sharky) GetSimilarArtists(artistID, limit, page int) []Artist {
	// TODO impelemnt
	return nil
}

// ================= Sessions =================

// Start a session
func (sharky *Sharky) StartSession() {
	result := sharky.NoSessionCallHttps("startSession", nil)
	if val, ok := result["sessionID"].(string); ok {
		sharky.SessionID = val
	}
	if isEmpty(sharky.SessionID) {
		log.Fatalln("Cannot obtain session ID.")
	}
}

// ================= Trials =================

// Gets a trial for an application and the provided uniqueID or logged in user.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetTrialInfo(uniqueID string) *TrialInfo {
	// TODO impelemnt
	return nil
}

// Starts a trial for a user bound to your application and the provided uniqueID.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) CreateTrial(uniqueID string) {
	// TODO impelemnt
}

// ================= Autocomplete =================

// Autocomplete search. Type parameter is 'music', 'playlist', or 'user'. Returns an array of words.
func (sharky *Sharky) GetAutocompleteSearchResults(query, typeParam string, limit int) []string {
	// TODO impelemnt
	return nil
}

// ================= Subscriber streams =================

// Get stream key, ID, etc. from songID for a subscriber account.
// Requires country object obtained from getCountry and a logged-in
// sessionID from a Grooveshark Anywhere subscriber.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetSubscriberStreamKey(songID int, country Country, lowBitrate bool, uniqueID string) *StreamKey {
	// TODO impelemnt
	return nil
}

// Mark a song as having been played for greater than or equal to 30 seconds.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) MarkStreamKeyOver30Secs(streamKey string, streamServerID int, uniqueID string) {
	// TODO impelemnt
}

// ================= Subscriber streams =================

// Mark a song as complete (played for greater than or equal to 30 seconds,
// and having reached the last second either through seeking or normal playback).
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) MarkSongComplete(songID int, streamKey string, streamServerID int, autoplayState AutoplayState) {
	// TODO impelemnt
}

// ================= Autoplay =================

// Grab a relevant song for autoplay
func (sharky *Sharky) GetAutoplaySong(autoplayState AutoplayState) *Song {
	// TODO impelemnt
	return nil
}

// Gets a list of tags (stations)
func (sharky *Sharky) GetAutoplayTags() []Tag {
	// TODO impelemnt
	return nil
}

// Start autoplay using a tag and grab a relevant song
func (sharky *Sharky) StartAutoplayTag(tagID int) {
	// TODO impelemnt
}

// Start autoplay and grab a relevant song
// TODO check if the params are right
func (sharky *Sharky) StartAutoplay(artistIDs, songIDs []string) {
	// TODO impelemnt
}

// Remove a vote up for a song
func (sharky *Sharky) RemoveVoteUpAutoplaySong(song Song, autoplayState AutoplayState) {
	// TODO impelemnt
}

// Vote up a song
func (sharky *Sharky) VoteUpAutoplaySong(song Song, autoplayState AutoplayState) {
	// TODO impelemnt
}

// Remove a song from the autoplay state
func (sharky *Sharky) RemoveSongFromAutoplay(song Song, autoplayState AutoplayState) {
	// TODO impelemnt
}

// Add a song to the autoplay state
func (sharky *Sharky) AddSongToAutoplay(song Song, autoplayState AutoplayState) {
	// TODO impelemnt
}

// Vote down a song
func (sharky *Sharky) VoteDownAutoplaySong(song Song, autoplayState AutoplayState) {
	// TODO imApelemnt
}

// Remove a vote down for a song
func (sharky *Sharky) RemoveVoteDownAutoplaySong(song Song, autoplayState AutoplayState) {
	// TODO impelemnt
}

// ================= Tinysong =================

// Get Grooveshark songID for tinysong base 62.
func (sharky *Sharky) GetSongIDFromTinysongBase62(base62 string) string {
	// TODO impelemnt
	return ""
}

// ================= Register =================

// Register and authenticate a user using an established session.
// The username is alpha-numeric with a period, dash or underscore allowed
// in the middle. The username can be blank or 5-32 characters.
// Passwords must be between 5 and 32 characters.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) RegisterUser(emailAddress, password, fullName, username, gender, birthDate string) {
	// TODO impelemnt
}

// ==================================== Reflection section ==================================

func mapToStruct(params map[string]interface{}, elem *reflect.Value) {
	for k, v := range params {
		setFieldOfElem(elem, k, v)
	}
}

func setFieldOfElem(elem *reflect.Value, key string, val interface{}) {
	field := elem.FieldByName(firstToUpper(key))

	if !field.CanSet() {
		return
	}

	switch field.Kind() {
	case reflect.String:
		if v, ok := val.(string); ok {
			field.SetString(v)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v := getInt64(val)
		field.SetInt(v)
	case reflect.Float32, reflect.Float64:
		v := getFloat64(val)
		field.SetFloat(v)
	case reflect.Bool:
		if v, ok := val.(bool); ok {
			field.SetBool(v)
		}
	}
}

func firstToUpper(value string) string {
	uni := []rune(value)
	uni[0] = unicode.ToUpper(uni[0])
	value = string(uni)
	return value
}

func getFloat64(value interface{}) float64 {
	if v, ok := value.(float32); ok {
		return float64(v)
	}
	if v, ok := value.(float64); ok {
		return v
	}
	return 0
}

func getInt64(value interface{}) int64 {
	if v, ok := value.(int); ok {
		return int64(v)
	}
	if v, ok := value.(int8); ok {
		return int64(v)
	}
	if v, ok := value.(int16); ok {
		return int64(v)
	}
	if v, ok := value.(int32); ok {
		return int64(v)
	}
	if v, ok := value.(rune); ok {
		return int64(v)
	}
	if v, ok := value.(int64); ok {
		return v
	}
	if v, ok := value.(float32); ok {
		return int64(v)
	}
	if v, ok := value.(float64); ok {
		return int64(v)
	}

	log.Fatal("!!!!!!!!!!! CANNOT PARSE INT VALUE !!!!!!!!!!!!")
	return 0
}

// ==================================== Util section ==================================

// Makes POST request to the API's method with params. SessionID should also
// be provided for some of the methods. You should also provide protocol (HTTP or HTTPS)
func makeCall(method string, params map[string]interface{}, sessionId, protocol, key, secret string) map[string]interface{} {
	reqData := buildRequestData(key, method, sessionId, params)
	buf, _ := json.Marshal(&reqData)
	signature := generateSignature(buf, []byte(secret))
	url := buildApiURL(signature, protocol)
	body := bytes.NewReader(buf)
	r, err := http.Post(url, CONTENT_TYPE, body)
	if err != nil {
		log.Fatal(err)
	}
	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	var resp Response
	json.Unmarshal(response, &resp)

	if resp.Errors != nil {
		error(resp.Errors, method)
	}

	return resp.Result
}

func error(errors []map[string]interface{}, method string) {
	line := "======================="
	errMessage := fmt.Sprintf("\n%v\nError while executing %v()\n%v\n", line, method, line)
	for _, err := range errors {
		code := err["code"]
		msg := err["message"]
		data := err["data"]
		errMessage += fmt.Sprintf("Error Code: %v, %v [%v]\n", code, msg, data)
	}
	log.Fatal(errMessage)
}

func buildRequestData(key, method, sessionID string, params map[string]interface{}) *RequestData {
	data := new(RequestData)
	data.Method = method
	data.Parameters = params

	header := make(map[string]string)
	header["wsKey"] = key
	header["sessionID"] = sessionID
	data.Header = header

	return data
}

// The signature is generated via HMAC using MD5 and the
// secret provided by Grooveshark team.
func generateSignature(postData, secret []byte) string {
	mac := hmac.New(md5.New, secret)
	mac.Write(postData)
	signature := fmt.Sprintf("%x", mac.Sum(nil))

	return signature
}

// Build the entire URL to the API. For some calls HTTPS
// protocol is not mandatory.
func buildApiURL(sig, protocol string) string {
	return protocol + API_HOST + API_ENDPOIT + SIG_GET_KEY + sig
}

// Util method to check empty values
func isEmpty(value string) bool {
	if len(strings.Trim(value, " ")) == 0 {
		return true
	} else {
		return false
	}
}

func md5sum(value string) string {
	h := md5.New()
	io.WriteString(h, value)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func logMsg(result map[string]interface{}, sucMsg, errMsg string) {
	if suc, ok := result["success"].(bool); ok {
		if suc {
			log.Println(sucMsg)
		} else {
			log.Fatal(errMsg)
		}
	}
}
