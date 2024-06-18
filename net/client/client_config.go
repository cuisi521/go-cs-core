package client

import (
	"crypto/tls"
	"errors"
	"net/http"
)

func (c *Client) SetHeader(key, value string) *Client {
	c.header[key] = value
	return c
}

func (c *Client) SetHeaderMap(m map[string]string) *Client {
	for k, v := range m {
		c.header[k] = v
	}
	return c
}

func (c *Client) SetAgent(agent string) *Client {
	c.header[httpHeaderUserAgent] = agent
	return c
}

func (c *Client) SetContentType(contentType string) *Client {
	c.header[httpHeaderContentType] = contentType
	return c
}

func (c *Client) SetTLSKeyCrt(crtFile, keyFile string) error {
	tlsConfig, err := LoadKeyCrt(crtFile, keyFile)
	if err != nil {
		return errors.New(err.Error() + "LoadKeyCrt failed")
	}
	if v, ok := c.Transport.(*http.Transport); ok {
		tlsConfig.InsecureSkipVerify = true
		v.TLSClientConfig = tlsConfig
		return nil
	}
	return errors.New(`cannot set TLSClientConfig for custom Transport of the client`)
}

func (c *Client) SetTLSConfig(tlsConfig *tls.Config) error {
	if v, ok := c.Transport.(*http.Transport); ok {
		v.TLSClientConfig = tlsConfig
		return nil
	}
	return errors.New(`cannot set TLSClientConfig for custom Transport of the client`)
}
