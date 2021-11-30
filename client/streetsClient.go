package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	Timeout time.Duration
}

func NewClient() *Client {
	client := Client{
		Timeout: 1 * time.Second,
	}
	return &client
}

type Response struct {
	Streets []string
}

type jsonData []string

func (c *Client) ReadStreets(url string) (content *Response, err error) {
	client := http.Client{
		Timeout: c.Timeout,
	}

	resp, err := client.Get(url)

	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return nil, err
		}

		content := string(body)

		var arr jsonData
		_ = json.Unmarshal([]byte(content), &arr)
		arr.deleteEmptyStreets()

		return &Response{Streets: arr}, nil
	} else {
		return nil, fmt.Errorf("get `%s` with response code '%s'", url, resp.Status)
	}
}

func (l *jsonData) deleteEmptyStreets() {
	var r []string
	for _, str := range *l {
		if str != "" {
			r = append(r, str)
		}
	}
	*l = r
}
