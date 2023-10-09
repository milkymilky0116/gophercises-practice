package models

import (
	"encoding/json"
	"os"
)

type AdventureStory struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func OpenFile(filename string) ([]byte, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}
func ParseJson(filename string) (map[string]AdventureStory, error) {
	buff, err := OpenFile(filename)
	if err != nil {
		return nil, err
	}
	result := make(map[string]AdventureStory)
	err = json.Unmarshal(buff, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
