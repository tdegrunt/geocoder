// Package geocoder provides support for interacting with Google Geocoding API
package geocoder

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"net/url"
)

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

type Viewport struct {
	Northeast Location `json:"north_east"`
	Southwest Location `json:"south_west"`
}

type Geometry struct {
	Location     Location `json:"location"`
	LocationType string   `json:"location_type"`
}

type AddressItem struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

type Result struct {
	AddressComponents []AddressItem `json:"address_components"`
	FormattedAddress  string        `json:"formatted_address"`
	Geometry          Geometry      `json:"geometry"`
	Types             []string      `json:"types"`
}

type Response struct {
	Status  string
	Results []Result
}

func (r *Response) BestMatch() (*Result, error) {
	if len(r.Results) > 0 {
		return &r.Results[0], nil
	}
	return nil, errors.New("ZERO_RESULTS")
}

type Geocoder struct {
	baseURL    string
	path       string
	language   string
	client     *string
	privateKey *string
}

// Creates a new google Geocoder client
func NewGeocoder() (*Geocoder, error) {
	var gc Geocoder

	gc.baseURL = "https://maps.googleapis.com"
	gc.path = "/maps/api/geocode/json"
	gc.language = "en"

	return &gc, nil
}

// Sets the Client for Enterprise accounts
func (gc *Geocoder) SetClient(c string) {
	gc.client = &c
}

// Sets the Private Key (or Crypto Key) for Enterprise accounts, use the value as is from your Enterprise Support Portal
func (gc *Geocoder) SetPrivateKey(pk string) {
	gc.privateKey = &pk
}

// Returns the full URL including client and signature (if PrivateKey is set)
func (gc *Geocoder) GetFullUrl(v url.Values) (*string, error) {

	if gc.client != nil {
		v.Add("client", *gc.client)
	}

	if gc.privateKey != nil {

		pathToEncode := gc.path + "?" + v.Encode()

		decodedKey, err := base64.URLEncoding.DecodeString(*gc.privateKey)
		if err != nil {
			return nil, err
		}

		h := hmac.New(sha1.New, decodedKey)
		h.Write([]byte(pathToEncode))

		signature := base64.URLEncoding.EncodeToString(h.Sum(nil))

		v.Add("signature", signature)
	}

	url := gc.baseURL + gc.path + "?" + v.Encode()
	return &url, nil
}
