package backend

type ErrorPage struct {
	ErrStatus string
	ErrMsg    string
}

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	LocationData LocationData
	DatesData    DatesData
	RelationData RelationData
}

type LocationData struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type DatesData struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type RelationData struct {
	ID            int                 `json:"id"`
	DatesLocation map[string][]string `json:"datesLocations"`
}

var (
	artists       []Artist
	locationsData []LocationData
	datesData     []DatesData
	relationData  []RelationData
	Err           error
)