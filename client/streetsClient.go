package client

import (
	"fmt"
	"net/http"
)

type Response struct {
	Streets []string
}

type jsonData []string

func (c *Client) ReadStreets(contextPath string) (response *Response, err error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", c.BaseURL, contextPath), nil)
	if err != nil {
		return nil, err
	}

	streets := make(jsonData, 0)
	if err := c.sendRequest(req, &streets); err != nil {
		return nil, err
	}

	streets.deleteEmptyStreets()

	return &Response{Streets: streets}, nil
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
