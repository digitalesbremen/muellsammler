package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	Timeout               time.Duration
	RetryTimeAfterTimeout time.Duration
}

func NewClient() *Client {
	client := Client{
		Timeout:               1 * time.Second,
		RetryTimeAfterTimeout: 5 * time.Second,
	}
	return &client
}

type Response struct {
	Addresses []string
}

type jsondata []string

func (c *Client) GetContent(url string) (content *Response, err error) {
	client := http.Client{
		Timeout: c.Timeout,
	}

	resp, err := client.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return nil, err
		}

		content := string(body)

		var arr jsondata
		_ = json.Unmarshal([]byte(content), &arr)
		arr.deleteEmpty()

		return &Response{Addresses: arr}, nil
	} else {
		return nil, fmt.Errorf("Try to load `%s`. Response code is '%s'", url, resp.Status)
	}
}

func (l *jsondata) deleteEmpty() {
	var r []string
	for _, str := range *l {
		if str != "" {
			r = append(r, str)
		}
	}
	*l = r
}
