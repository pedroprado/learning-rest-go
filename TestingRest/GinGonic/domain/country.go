package domain

type Country struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	Locale         string         `json:"locale"`
	GeoInformation GeoInformation `json:"geo_information"`
	States         []State        `json:"states"`
}

type GeoInformation struct {
	Location Location `json:"location"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:""longitude`
}

type State struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
