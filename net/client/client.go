// package client
// @Author cuisi
// @Date 2024/6/12 11:06:00
// @Desc 客户端请求类
package client

import (
	"crypto/rand"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	go_cs_core "github.com/cuisi521/go-cs-core"
)

type Client struct {
	http.Client
	header  map[string]string // Custom header map.
	cookies map[string]string // Custom cookie map.
}

const (
	MaxIdleConnsPerHost int = 20000
	RequestTimeout      int = 30
	MaxIdleConns        int = 20000
)

const (
	httpProtocolName          = `http`
	httpParamFileHolder       = `@file:`
	httpRegexParamJson        = `^[\w\[\]]+=.+`
	httpRegexHeaderRaw        = `^([\w\-]+):\s*(.+)`
	httpHeaderHost            = `Host`
	httpHeaderCookie          = `Cookie`
	httpHeaderUserAgent       = `User-Agent`
	httpHeaderContentType     = `Content-Type`
	httpHeaderContentTypeJson = `application/json`
	httpHeaderContentTypeXml  = `application/xml`
	httpHeaderContentTypeForm = `application/x-www-form-urlencoded`
)

var (
	hostname, _        = os.Hostname()
	defaultClientAgent = fmt.Sprintf(`Client %s at %s`, go_cs_core.VERSION, hostname)
)

func New() *Client {
	c := &Client{
		Client: http.Client{
			Transport: &http.Transport{
				// No validation for https certification of the server in default.
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				DisableKeepAlives:   true,
				MaxIdleConnsPerHost: MaxIdleConnsPerHost,
				MaxIdleConns:        MaxIdleConns,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
					DualStack: true,
				}).DialContext,
				IdleConnTimeout:       90 * time.Second,
				ResponseHeaderTimeout: 30 * time.Second,
				// TLSHandshakeTimeout:   10 * time.Second,
			},
			Timeout: time.Duration(RequestTimeout) * time.Second,
		},
		header:  make(map[string]string),
		cookies: make(map[string]string),
	}
	return c
}

// Clone deeply clones current client and returns a new one.
func (c *Client) Clone() *Client {
	newClient := New()
	*newClient = *c
	if len(c.header) > 0 {
		newClient.header = make(map[string]string)
		for k, v := range c.header {
			newClient.header[k] = v
		}
	}
	if len(c.cookies) > 0 {
		newClient.cookies = make(map[string]string)
		for k, v := range c.cookies {
			newClient.cookies[k] = v
		}
	}
	return newClient
}

// LoadKeyCrt creates and returns a TLS configuration object with given certificate and key files.
func LoadKeyCrt(crtFile, keyFile string) (*tls.Config, error) {

	crtPath, err := filepath.Abs(crtFile)
	if err != nil {
		return nil, err
	}
	keyPath, err := filepath.Abs(crtFile)
	if err != nil {
		return nil, err
	}
	crt, err := tls.LoadX509KeyPair(crtPath, keyPath)
	if err != nil {
		err = errors.New(fmt.Sprint(err, `tls.LoadX509KeyPair failed for certFile "%s", keyFile "%s"`, crtPath, keyPath))
		return nil, err
	}
	tlsConfig := &tls.Config{}
	tlsConfig.Certificates = []tls.Certificate{crt}
	tlsConfig.Time = time.Now
	tlsConfig.Rand = rand.Reader
	return tlsConfig, nil
}
