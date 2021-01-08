package fbi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// List all fugitives from the Most Wanted API
func (f *Fugitives) List() error {
	url := BaseURL + "/list"
	method := "GET"
	c := http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "fbi-most-wanted-golang-client")
	req.Header.Add("Content-Type", "application/json")
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &f); err != nil {
		return err
	}

	return nil
}

// ListPretty lists all fugitives from the Most Wanted API and then
// auto formats results as pretty-printed JSON
func (f *Fugitives) ListPretty() ([]byte, error) {
	err := f.List()
	if err != nil {
		return nil, err
	}

	j, err := json.MarshalIndent(&f, "", "\t")
	if err != nil {
		return nil, err
	}

	return j, nil
}

// Fugitives is the base structure for the
// FBI most wanted REST API response and
// contains a list of individuals as well as the
// total items found
type Fugitives struct {
	Total int          `json:"total"`
	Items []Individual `json:"items"`
}

// Options contains the available
// query params for filter API request
// results
type Options struct {
	Title        string `url:"title"`
	FieldOffices string `url:"field_offices"`
	Page         int    `url:"page"`
}

// Find all fugitives matching the parameters provided in the
// opt struct
func (f *Fugitives) Find(opt *Options) error {
	if opt == nil {
		return fmt.Errorf("no options supplied")
	}

	q, _ := query.Values(opt)
	url := BaseURL + "/list" + q.Encode()
	method := "GET"
	c := http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "fbi-most-wanted-golang-client")
	req.Header.Set("Content-Type", "application/json")
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &f); err != nil {
		return err
	}

	return nil
}

// Individual represents a single person in the FBI
// most wanted list
type Individual struct {
	Images                []Images      `json:"images"`
	Occupations           interface{}   `json:"occupations"`
	DatesOfBirthUsed      interface{}   `json:"dates_of_birth_used"`
	WeightMax             interface{}   `json:"weight_max"`
	AgeMax                interface{}   `json:"age_max"`
	AgeRange              interface{}   `json:"age_range"`
	AdditionalInformation interface{}   `json:"additional_information"`
	WeightMin             interface{}   `json:"weight_min"`
	RewardText            string        `json:"reward_text"`
	Aliases               interface{}   `json:"aliases"`
	UID                   string        `json:"uid"`
	HairRaw               interface{}   `json:"hair_raw"`
	Race                  interface{}   `json:"race"`
	RaceRaw               interface{}   `json:"race_raw"`
	Eyes                  interface{}   `json:"eyes"`
	Hair                  interface{}   `json:"hair"`
	Sex                   interface{}   `json:"sex"`
	Languages             interface{}   `json:"languages"`
	WarningMessage        interface{}   `json:"warning_message"`
	PossibleCountries     interface{}   `json:"possible_countries"`
	Remarks               interface{}   `json:"remarks"`
	Build                 interface{}   `json:"build"`
	HeightMin             interface{}   `json:"height_min"`
	Suspects              interface{}   `json:"suspects"`
	Publication           string        `json:"publication"`
	Status                string        `json:"status"`
	Subjects              []string      `json:"subjects"`
	Title                 string        `json:"title"`
	ScarsAndMarks         interface{}   `json:"scars_and_marks"`
	Path                  string        `json:"path"`
	PossibleStates        interface{}   `json:"possible_states"`
	Nationality           interface{}   `json:"nationality"`
	Files                 []Files       `json:"files"`
	Coordinates           []interface{} `json:"coordinates"`
	PersonClassification  string        `json:"person_classification"`
	Description           string        `json:"description"`
	PlaceOfBirth          interface{}   `json:"place_of_birth"`
	HeightMax             interface{}   `json:"height_max"`
	Locations             interface{}   `json:"locations"`
	Caution               interface{}   `json:"caution"`
	EyesRaw               interface{}   `json:"eyes_raw"`
	Ncic                  interface{}   `json:"ncic"`
	LegatNames            interface{}   `json:"legat_names"`
	Complexion            interface{}   `json:"complexion"`
	FieldOffices          []string      `json:"field_offices"`
	RewardMin             int           `json:"reward_min"`
	RewardMax             int           `json:"reward_max"`
	AgeMin                interface{}   `json:"age_min"`
	URL                   string        `json:"url"`
	Details               string        `json:"details"`
	Modified              time.Time     `json:"modified"`
	Weight                interface{}   `json:"weight"`
	ID                    string        `json:"@id"`
}

// Images are the pictures available
// of the individual
type Images struct {
	Large    string      `json:"large"`
	Caption  interface{} `json:"caption"`
	Original string      `json:"original"`
	Thumb    string      `json:"thumb"`
}

// Download files to a file from the URLs provided
func (f Files) Download(filename string) error {
	var url string
	if f.URL != "" {
		url = f.URL
	} else {
		return fmt.Errorf("no image urls in struct")
	}

	//Get the response bytes from the url
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("received non 200 response code")
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

// Download images to a file from the URLs provided
func (i Images) Download(filename string) error {
	var url string
	if i.Large != "" {
		url = i.Large
	} else if i.Original != "" {
		url = i.Original
	} else if i.Thumb != "" {
		url = i.Thumb
	} else {
		return fmt.Errorf("no image urls in struct")
	}

	//Get the response bytes from the url
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("received non 200 response code")
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

// Files are any related  files available
// pertaining to the individual
type Files struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
