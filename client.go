package bhdapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

type Client struct {
	cl        *http.Client
	AddRssKey bool
	ApiKey    string
	RssKey    string
	Transport http.RoundTripper
}

func New(opts ...Option) *Client {
	cl := &Client{}
	for _, o := range opts {
		o(cl)
	}
	if cl.cl == nil {
		cl.cl = &http.Client{
			Transport: cl.Transport,
		}
	}
	return cl
}

func (cl *Client) Do(ctx context.Context, action string, params, result interface{}) error {
	if cl.ApiKey == "" {
		return errors.New("must supply api key")
	}
	m := map[string]interface{}{
		"action": action,
	}
	if cl.AddRssKey && cl.RssKey != "" {
		m["rsskey"] = cl.RssKey
	}
	v := reflect.ValueOf(params)
	if v.Kind() != reflect.Struct {
		return errors.New("must past a pointer to a struct")
	}
	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		tag := strings.SplitN(typ.Field(i).Tag.Get("json"), ",", 2)[0]
		if tag == "-" || tag == "" {
			continue
		}
		vv, ok, err := val(v.Field(i).Interface())
		if err != nil {
			return fmt.Errorf("invalid field %d: %w", i, err)
		}
		if ok {
			m[tag] = vv
		}
	}
	buf, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", "https://beyond-hd.me/api/torrents/"+cl.ApiKey, bytes.NewReader(buf))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := cl.cl.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid http status %d", res.StatusCode)
	}
	dec := json.NewDecoder(res.Body)
	dec.DisallowUnknownFields()
	return dec.Decode(result)
}

// Search searches for a query.
func (cl *Client) Search(ctx context.Context, query ...string) (*SearchResponse, error) {
	return Search(query...).Do(ctx, cl)
}

// Torrent retrieves a torrent for the id.
func (cl *Client) Torrent(ctx context.Context, id int) ([]byte, error) {
	if cl.RssKey == "" {
		return nil, errors.New("must supply rss key")
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://beyond-hd.me/torrent/download/auto.%d.%s", id, cl.RssKey), nil)
	if err != nil {
		return nil, err
	}
	res, err := cl.cl.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid http status %d", res.StatusCode)
	}
	return ioutil.ReadAll(res.Body)
}

type Option func(cl *Client)

func WithApiKey(apiKey string) Option {
	return func(cl *Client) {
		cl.ApiKey = apiKey
	}
}

func WithRssKey(rssKey string, addRssKey bool) Option {
	return func(cl *Client) {
		cl.RssKey, cl.AddRssKey = rssKey, addRssKey
	}
}

func WithTransport(transport http.RoundTripper) Option {
	return func(cl *Client) {
		cl.Transport = transport
	}
}

func val(v interface{}) (interface{}, bool, error) {
	switch x := v.(type) {
	case []string:
		return strings.Join(x, ","), len(x) != 0, nil
	case string:
		return x, len(x) != 0, nil
	case int64:
		return x, x != 0, nil
	case int:
		return x, x != 0, nil
	case Bool:
		return x.Int(), x != false, nil
	case fmt.Stringer:
		y := x.String()
		return y, len(y) != 0, nil
	}
	return "", false, fmt.Errorf("unknown type %T", v)
}
