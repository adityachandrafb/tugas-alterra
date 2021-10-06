package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Artist struct {
	ID   string `json:"id"`
	AAAA string `json:"sort-name"`
	Type string `json:"type"`
}

func GetArtist() Artist {
	// akses url
	resp, err := http.Get("https://musicbrainz.org/ws/2/artist/cc197bad-dc9c-440d-a5b5-d52ba2e14234?fmt=json")
	if err != nil {
		panic(err)
	}

	// ambil data dari response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// masukkan data ke struct
	var res Artist
	err = json.Unmarshal(body, &res)
	if err != nil {
		panic(err)
	}

	return res
}


// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"io/ioutil"
// )

// func music() {

// 	url := "https://theaudiodb.p.rapidapi.com/playlist.php?format=track"

// 	req, _ := http.NewRequest("GET", url, nil)

// 	req.Header.Add("x-rapidapi-host", "theaudiodb.p.rapidapi.com")
// 	req.Header.Add("x-rapidapi-key", "36825d9b32msh9525364d5c88032p1d1e68jsn74b19b2c82d8")

// 	res, _ := http.DefaultClient.Do(req)

// 	defer res.Body.Close()
// 	body, _ := ioutil.ReadAll(res.Body)

// 	fmt.Println(res)
// 	fmt.Println(string(body))

// }