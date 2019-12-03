package main

type ApiUrl struct {
	Url     string `json:"url"`
	Name    string `json:"name"`
	Id      string `json:"id,omitempty"`
	Version string `json:"version,omitempty"`
}
