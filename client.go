package gortt

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

var ErrNoUsername = errors.New("no username supplied; did you mean to supply RTTAPIUSER?")
var ErrNoPassword = errors.New("no password supplied; ; did you mean to supply RTTAPIPASS?")

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	Username   string
	Password   string
}

// Returns a new client using the default HTTP Client
// and base URL.  If username or password are
// empty, it attempts to get them out of the
// RTTAPIUSER and RTTAPIPASS
func NewClient(user string, pass string) (*Client, error) {
	c := &Client{
		HTTPClient: http.DefaultClient,
		BaseURL:    "https://api.rtt.io/api/v1",
		Username:   user,
		Password:   pass,
	}

	if user == "" {
		c.Username = os.Getenv("RTTAPIUSER")
	}

	if pass == "" {
		c.Password = os.Getenv("RTTAPIPASS")
	}

	if c.Username == "" {
		return nil, ErrNoUsername
	}

	if c.Password == "" {
		return nil, ErrNoPassword
	}

	return c, nil
}

func (c *Client) GetServicesFromStationForDate(ctx context.Context, station string, date time.Time) (*RTTLocationLineup, error) {
	dateString := date.Format("2006/01/02")
	url := fmt.Sprintf("%s/json/search/%s/%s", c.BaseURL, url.PathEscape(station), dateString)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Username, c.Password)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r RTTLocationLineup
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
