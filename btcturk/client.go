package btcturk

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const baseURL = "https://www.btcturk.com"

type Client struct {
	client  *http.Client
	baseURL *url.URL
	params  url.Values
	body    io.Reader

	publicKey  string
	privateKey string
}

func NewBTCTurkClient() *Client {
	baseURL, _ := url.Parse(baseURL)
	return &Client{
		baseURL: baseURL,
		client:  &http.Client{},
		params:  url.Values{},
		body:    nil,
	}
}

func (c *Client) SetAuthKey(publicKey, privateKey string) {
	c.publicKey, c.privateKey = publicKey, privateKey
}

func (c *Client) SetBaseURL(basURL string) (err error) {
	c.baseURL, err = url.Parse(basURL)
	return
}

func (c *Client) newRequest(method, relURL string, body io.Reader) (*http.Request, error) {
	rel, err := url.Parse(relURL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, c.baseURL.ResolveReference(rel).String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(context.Background())

	return req, nil
}

func (c *Client) do(r *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(r)
	if err != nil {
		return nil, err
	}

	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		c.clearRequest()
	}()

	if json.NewDecoder(resp.Body).Decode(v) != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) auth(req *http.Request) error {
	if c.privateKey == "" {
		return errors.New("Private Key is not set")
	}

	if c.publicKey == "" {
		return errors.New("Public Key is not set")
	}

	key, err := base64.StdEncoding.DecodeString(c.privateKey)
	if err != nil {
		return err
	}

	stamp := fmt.Sprint(time.Now().Unix())
	message := c.publicKey + stamp

	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))

	req.Header.Set("X-PCK", c.publicKey)
	req.Header.Set("X-Stamp", stamp)
	req.Header.Set("X-Signature", base64.StdEncoding.EncodeToString(h.Sum(nil)))
	return nil
}

func (c *Client) clearRequest() {
	c.params = url.Values{}
	c.body = nil
}
