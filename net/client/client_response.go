package client

import (
	"io"
	"net/http"
)

type Response struct {
	*http.Response
	request *http.Request
}

func (r *Response) ReadAll() []byte {
	// Response might be nil.
	if r == nil || r.Response == nil {
		return []byte{}
	}

	body, err := io.ReadAll(r.Response.Body)
	if err != nil {
		return nil
	}
	return body
}

func (r *Response) ReadAllString() string {
	return string(r.ReadAll())
}

func (r *Response) Close() error {
	if r == nil || r.Response == nil {
		return nil
	}
	return r.Response.Body.Close()
}
