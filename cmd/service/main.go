package main

import (
	"errors"
	"fmt"

	"github.com/WelverinRaven/SpotifyPlaylistAnalyzer/pkg/models"
	"github.com/go-resty/resty"
)

const baseURL = "https://spotify.com/api/"
const clientID = ""
const clientSecret = ""
const playlistID = ""

func main() {
	accessToken, err := authenticate()
	playlist := &models.Playlist{}
	if err != nil {
		fmt.Println("Authentication failed... Shutting down")
		return
	}
	fmt.Println(accessToken)

	resp, err := resty.New().R().SetResult(playlist).SetHeader("Authorization", accessToken).Get("https://api.spotify.com/v1/playlists/" + playlistID)
	if err != nil || resp.IsError() {
		fmt.Println(resp.Body())
		fmt.Println("Also:")
		panic(err)
	}
	fmt.Println("it worked!")
	fmt.Println(playlist)
}

func authenticate() (string, error) {
	authResponse := &models.AuthResponse{}
	key := "" // base64 encoded clientID:clientSecret
	body := models.AuthBody{GrantType: "client_credentials"}
	resp, err := resty.New().R().SetResult(authResponse).SetHeader("Authorization", key).SetBody(body).Post("https://accounts.spotify.com/api/token")
	if err != nil || resp.IsError() {
		fmt.Println(resp.Body())
		return "", errors.New("something wrong:")
	}
	return authResponse.AccessToken, nil
}
