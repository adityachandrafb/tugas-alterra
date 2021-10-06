package database

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Audio struct {
	ID   string `json:"id"`
	Audio string `json:"audio"`
}

func GetAudio() Audio {
	// akses url
	resp, err := http.Get("https://musicbrainz.org/instrument/d5cc3c69-218e-449a-b80d-8bd7a61311a1fmt=json")
	if err != nil {
		panic(err)
	}

	// ambil data dari response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// masukkan data ke struct
	var res Audio
	err = json.Unmarshal(body, &res)
	if err != nil {
		panic(err)
	}

	return res
}