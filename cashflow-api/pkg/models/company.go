package models

import (
	"encoding/json"
	"strconv"
	"strings"
)

type StringSlice []string

func (s *StringSlice) UnmarshalJSON(data []byte) error {
	// Handle null values
	if string(data) == "null" {
		*s = nil
		return nil
	}

	// Remove surrounding quotes and unescape the string
	unquoted, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	// Replace single quotes with double quotes
	unquoted = strings.ReplaceAll(unquoted, "'", "\"")

	// Parse the cleaned string as JSON array
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
