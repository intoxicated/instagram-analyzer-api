package dto

type LocationRequest struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Radius    float64 `json:"radius"`
}

type HashtagRequest struct {
	Hashtag string `json:"hashtag"`
}
