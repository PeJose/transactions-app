package helpers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func BuildUrl(base string, query map[string]string) (url.URL, error) {
	u, err := url.Parse(base)
	if err != nil {
		return url.URL{}, err
	}

	v := url.Values{}
	for key, value := range query {
		v.Set(key, value)
	}

	u.RawQuery = v.Encode()

	return *u, nil
}

func ReadBody(res *http.Response, v interface{}) error {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, v); err != nil {
		return err
	}

	return nil
}
