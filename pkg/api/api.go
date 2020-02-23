package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
)

type Client struct {
	Adress string
}

type SearchResponse struct {
	Id   string
	Name string
}

func (c Client) Search(search string) []SearchResponse {
	url := fmt.Sprintf("%v/stations?query=%v", c.Adress, url2.QueryEscape(search))
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var res []SearchResponse
	json.Unmarshal(body, &res)
	return res
}

type LineResponse struct {
	Name string
}
type DeparturesResponse struct {
	When      string
	Direction string
	Line      LineResponse
	Platform  string
}

func (c Client) Departures(stationId string) []DeparturesResponse {
	url := fmt.Sprintf("%v/stations/%v/departures", c.Adress, stationId)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var res []DeparturesResponse
	json.Unmarshal(body, &res)
	return res
}
