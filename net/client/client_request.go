package client

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/cuisi521/go-cs-core/tool/util/conv"
)

func (c *Client) Get(url string, body []byte) (resp *Response, err error) {
	resp, err = c.httpDo(url, http.MethodGet, body)
	return
}

func (c *Client) Post(url string, body []byte) (resp *Response, err error) {
	resp, err = c.httpDo(url, http.MethodPost, body)
	return
}

func (c *Client) Put(url string, body []byte) (resp *Response, err error) {
	resp, err = c.httpDo(url, http.MethodPut, body)
	return
}

func (c *Client) PATCH(url string, body []byte) (resp *Response, err error) {
	resp, err = c.httpDo(url, http.MethodPatch, body)
	return
}

func (c *Client) Delete(url string, body []byte) (resp *Response, err error) {
	resp, err = c.httpDo(url, http.MethodDelete, body)
	return
}

func (c *Client) httpDo(url, method string, body []byte) (resp *Response, err error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return
	}
	for k, v := range c.header {
		req.Header.Set(k, v)
	}
	resp = &Response{
		request: req,
	}
	resp.Response, err = c.Client.Do(req)
	return
}

func (c *Client) parseRequest(method, url string, data ...interface{}) (req *http.Request, err error) {
	method = strings.ToUpper(method)
	var (
		params     string
		bodyBuffer *bytes.Buffer
	)
	if len(data) > 0 {
		switch c.header[httpHeaderContentType] {
		case httpHeaderContentTypeJson:
			switch data[0].(type) {
			case string, []byte:
				params = conv.Str(data[0])
			default:
				if b, err := json.Marshal(data[0]); err != nil {
					return nil, err
				} else {
					params = string(b)
				}
			}
		case httpHeaderContentTypeXml:
			switch data[0].(type) {
			case string, []byte:
				params = conv.Str(data[0])
			default:
				if b, err := xml.Marshal(data[0]); err != nil {
					return nil, err
				} else {
					params = string(b)
				}
			}
		default:
			params = c.BuildParams(data[0])
		}
	}

	if method == http.MethodGet {
		if params != "" {
			switch c.header[httpHeaderContentType] {
			case httpHeaderContentTypeJson, httpHeaderContentTypeXml:
				bodyBuffer = bytes.NewBuffer([]byte(params))
			default:
				if strings.Contains(url, "?") {
					url = url + "&" + params
				} else {
					url = url + "?" + params
				}
				bodyBuffer = bytes.NewBuffer(nil)
			}
		} else {
			bodyBuffer = bytes.NewBuffer(nil)
		}
		if req, err = http.NewRequest(method, url, bodyBuffer); err != nil {
			err = errors.New(fmt.Sprintf(`http.NewRequest failed with method "%s" and URL "%s"`, method, url))
			return nil, err
		}
	}
	return
}

func (c *Client) BuildParams(params interface{}, noUrlEncode ...bool) (encodedParamStr string) {
	// If given string/[]byte, converts and returns it directly as string.
	switch v := params.(type) {
	case string, []byte:
		return conv.Str(params)
	case []interface{}:
		if len(v) > 0 {
			params = v[0]
		} else {
			params = nil
		}
	}
	return
}
