package models

type MonumentAttribute struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	IsOpen  bool   `json:"isOpen"`
}
