package geocoder

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// Latitude/Longitude Lookup
func (gc *Geocoder) Geocode(address string) (*Response, error) {

	v := url.Values{}
	v.Set("address", address)
	v.Add("language", gc.language)
	
	url, err := gc.GetFullUrl(v)
	if err != nil {
		return nil, err
	}
		
	resp, err := http.Get(*url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r Response

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&r)
	if err != nil {
		return nil, err
	}
		
	return &r, nil
}