package models

import (
	"encoding/json"
	"strconv"
	"strings"
)

type StringSlice []string

func (s *StringSlice) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*s = nil
		return nil
	}

	unquoted, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	unquoted = strings.ReplaceAll(unquoted, "'", "\"")

	var result []string
	if err := json.Unmarshal([]byte(unquoted), &result); err != nil {
		return err
	}

	*s = result
	return nil
}

type Company struct {
	Id      int         `json:"id"`
	Ibans   StringSlice `json:"ibans" gorm:"serializer:json"`
	Name    string      `json:"name"`
	Address string      `json:"address"`
}
