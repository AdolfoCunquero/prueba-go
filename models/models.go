package models

type Item struct {
	Id       string `json:"id"`
	Icon_url string `json:"icon_url"`
	Url      string `json:"url"`
	Value    string `json:"value"`
}
