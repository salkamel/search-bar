package backend

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func InitData() error { // initialize the global variables for convenience purposes
	artists, locationsData, datesData, relationData, Err = FetchAllData()
	if Err != nil {
		return Err
	}
	artists = MergeStructs(artists, locationsData, datesData, relationData)
	return nil
}

func FetchAllData() ([]Artist, []LocationData, []DatesData, []RelationData, error) {
	var artists []Artist
	var locationsData struct {
		Index []LocationData `json:"index"`
	}
	var datesData struct {
		Index []DatesData `json:"index"`
	}
	var relationData struct {
		Index []RelationData `json:"index"`
	}

	err := FetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	err = FetchData("https://groupietrackers.herokuapp.com/api/locations", &locationsData)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	err = FetchData("https://groupietrackers.herokuapp.com/api/dates", &datesData)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	err = FetchData("https://groupietrackers.herokuapp.com/api/relation", &relationData)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return artists, locationsData.Index, datesData.Index, relationData.Index, nil
}

func FetchData(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch data from %s: %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch data from %s: status code %d", url, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body from %s: %v", url, err)
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON from %s: %v", url, err)
	}

	return nil
}

func MergeStructs(artists []Artist, locationsData []LocationData, datesData []DatesData, relationData []RelationData) []Artist {
	locationDataMap := make(map[int]LocationData)
	for _, loc := range locationsData {
		locationDataMap[loc.ID] = loc
	}

	datesDataMap := make(map[int]DatesData)
	for _, date := range datesData {
		datesDataMap[date.ID] = date
	}

	relationDataMap := make(map[int]RelationData)
	for _, relation := range relationData {
		relationDataMap[relation.ID] = relation
	}

	for i, artist := range artists {
		artists[i].LocationData = locationDataMap[artist.ID]
		artists[i].DatesData = datesDataMap[artist.ID]
		artists[i].RelationData = relationDataMap[artist.ID]
	}

	return artists
}

func GetURLPort(url1 string) (string) {
	parsedURL, err := url.Parse(url1)
	if err != nil {
		fmt.Println("Error parsing URL:", err);	return "Error parsing URL to get port"
	}
	port := parsedURL.Port()
	if port == "" {
		port = "80" // Default HTTP port
	}
	return port
}
