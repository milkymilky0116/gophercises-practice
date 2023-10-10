package models

import (
	"encoding/json"
	"fmt"
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

type AdventureModel struct {
	Stories map[string]*AdventureStory
}

func OpenFile(filename string) ([]byte, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}
func ParseJson(filename string) (map[string]*AdventureStory, error) {
	buff, err := OpenFile(filename)
	if err != nil {
		return nil, err
	}
	result := make(map[string]*AdventureStory)
	err = json.Unmarshal(buff, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m *AdventureModel) Get(arc string) (*AdventureStory, error) {
	page, ok := m.Stories[arc]
	if !ok {
		err := fmt.Errorf("the story arc %s is not exists", arc)
		return nil, err
	}
	return page, nil
}
