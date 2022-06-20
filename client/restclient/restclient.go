package restclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Post(url, authorization string, body interface{}, headers http.Header) (*http.Response, error) {

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", authorization)
	client := http.Client{}
	return client.Do(request)
}
