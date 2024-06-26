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

	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit")
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(data, &locationAreasResp)
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
		return LocationAreasResp{}, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreasResp, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (AreaResp, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit")
		areaResp := AreaResp{}
		err := json.Unmarshal(data, &areaResp)
		if err != nil {
			return AreaResp{}, err
		}

		return areaResp, nil
	}
	fmt.Println("cache miss")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return AreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AreaResp{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return AreaResp{}, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return AreaResp{}, err
	}

	areaResp := AreaResp{}
	err = json.Unmarshal(data, &areaResp)
	if err != nil {
		return AreaResp{}, err
	}

	c.cache.Add(fullURL, data)

	return areaResp, nil
}
