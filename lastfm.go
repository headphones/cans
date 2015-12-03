package main
import (
	"time"
	"net/url"
	"fmt"
	"net/http"
	"errors"
	"io"
)

const apiRoot = "http://ws.audioscrobbler.com/2.0/"
const apiKey = "94381a2146e972044b745b93575ddeac"

//var imageSizes = map[string]ImageSize

type Tag struct {
	Name string
	Url string
}

type ImageSize struct {
	Width int
	Height int
}

type Image struct {
	Url string
	Width int
	Height int
}

type ArtistStats struct {
	Listeners int
	Plays int
}

type Artist struct {
	Name string
	Mbid string
	Url	 string
	Images []Image
	Streamable bool

	Similar []Artist
}

type ArtistBio struct {
	Published time.Time
	Summary string
	Content string
}

func buildUrl(base string, args ...string) (string, error) {
	if len(args) < 1 {
		return "", errors.New("no arguments passed")
	} else if (len(args) % 2) != 0 {
		return "", errors.New("bad number of arguments")
	}
	qp := url.Values{}
	qp.Add("api_key", apiKey)
	for i := 0; i < len(args); i+=2 {
		qp.Add(args[i], args[i+1])
	}
	url := base + "?" + qp.Encode()
	return url, nil
}

func getArtist(w http.ResponseWriter, artist *Artist) error {
	url, err := buildUrl(apiRoot, "method", "artist.getInfo", "artist.name", artist.Name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("requesting: " + url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
		return err
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
	return nil
}

func testArtistSearch(w http.ResponseWriter, r *http.Request) {
	artist := Artist{Name:"The Black Keys"}
	getArtist(w, &artist)
}
