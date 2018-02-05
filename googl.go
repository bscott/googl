package googl

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

// Googl struct
type Googl struct {
	Key string
}

// ShortMsg struct
type ShortMsg struct {
	Kind    string `json:"kind"`
	ID      string `json:"id"`
	LongURL string `json:"longUrl"`
}

// LongMsg struct
type LongMsg struct {
	Kind    string `json:"kind"`
	ID      string `json:"id"`
	LongURL string `json:"longUrl"`
	Status  string `json:"status"`
}

// NewClient returns a Client of Googl
func NewClient(key string) *Googl {
	return &Googl{Key: key}
}

// Shorten function takes a longUrl and return a struct with the shorten data
func (c *Googl) Shorten(url string) (ShortMsg, error) {
	request := gorequest.New()
	var response ShortMsg
	var err string
	if c.Key == "" {
		err := "You need to set the Google Url Shortener API Key"
		return response, errors.New(err)
	} else if url == "" {
		err := "You need to set the url to be shortened"
		return response, errors.New(err)
	} else {
		// Create Request URL
		gURL := "https://www.googleapis.com/urlshortener/v1/url?key=" + c.Key
		resp, _, _ := request.Post(gURL).
			Set("Accept", "application/json").
			Set("Content-Type", "application/json").
			Send(`{"longUrl":"` + url + `"}`).End()
		if resp.Status == "200 OK" {
			json.NewDecoder(resp.Body).Decode(&response)

		} else {
			err := "Some error occurred, please try again later"
			return response, errors.New(err)
		}
	}

	return response, nil
}

// Expand takes a shortURL & returns the expanded URL
func (c *Googl) Expand(shortURL string) string {
	request := gorequest.New()
	var response string

	gURL := "https://www.googleapis.com/urlshortener/v1/url?key=" + c.Key + "&shortUrl=" + shortURL
	if c.Key == "" {
		response = "You need to set the Google Url Shortener API Key"
	} else if shortURL == "" {
		response = "You need to set the url to be expanded"
	} else {
		resp, body, _ := request.Get(gURL).
			Set("Accept", "application/json").
			Set("Content-Type", "application/json").End()
		if resp.Status == "200 OK" {
			fmt.Println(body)
			response = "Done! Ok!"
		} else {
			response = "Some error occurred, please try again later"
		}
	}

	return response
}
