package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {

	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	// check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		//cache hit
		fmt.Println("cache hit!")
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(dat, &locationAreasResp)

		if err != nil {
			return LocationAreasResp{}, err
		}

		return locationAreasResp, nil

	}
	fmt.Println("cache miss")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(dat, &locationAreasResp)

	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationAreasResp, nil

}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {

	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	// check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		//cache hit
		fmt.Println("cache hit!")
		locationArea := LocationArea{}
		err := json.Unmarshal(dat, &locationArea)

		if err != nil {
			return LocationArea{}, err
		}

		return LocationArea{}, nil

	}
	fmt.Println("cache miss")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(dat, &locationArea)

	if err != nil {
		return locationArea, err
	}

	c.cache.Add(fullURL, dat)

	return locationArea, nil

}
