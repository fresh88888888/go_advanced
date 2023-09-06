package git

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/context/ctxhttp"
	"golang.org/x/oauth2"
)

const (
	defaultBaseUrl      = "https://api.github.com"
	acceptVersionHeader = "application/vnd.github.v3+json" // https://developer.github.com/v3/#current-version
)

// A client manages communication with the github API
type Client struct {
	// HTTP client used to communicate with the API
	client *http.Client
	// Base url for API requests. Defaults to the public github API
	baseUrl *url.URL
	// Service used for handling the github API
	Repositories *RepositoriesService
}

type ErrorResponse struct {
	Response *http.Response
	Message  string `json:"message,omitempty"`
}

// NewClient returns a new gitLab API client. You must provide a vaild token.
func NewClient(ctx context.Context, token string) *Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := newClient(tc)

	return client
}

func newClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{client: httpClient}

	c.Repositories = &RepositoriesService{client: c}

	return c
}

// GetBaseURL return a copy of the baseURL.
func (c *Client) GetBaseURL() *url.URL {
	u := *c.baseUrl
	return &u
}

// SetBaseURL sets the base URL for API requests to a custom endpoint. urlStr should always be specified with a trailing slash
func (c *Client) SetBaseURL(urlStr string) error {
	// Make sure the given URL end with a slash
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}
	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	// Update the base URL of the client
	c.baseUrl = baseURL

	return nil
}

// NewRequest create an API request using a relative URL can be provide in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.baseUrl.ResolveReference(rel)
	buf, err := c.encodeRequestBody(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	// Set version header
	req.Header.Set("Accept", acceptVersionHeader)
	// req.Header.Set("User-Agent", )

	return req, nil
}

// Do send an API request and returns the API response. The API response is Json decoded and stored in the value pointed to by v,
// or returned as error if an API error has occurred.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := ctxhttp.Do(ctx, c.client, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		reer := resp.Body.Close()
		if reer != nil {
			err = reer
		}
	}()

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	err = c.decodeResponseBody(resp.Body, v)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (c *Client) decodeResponseBody(body io.ReadCloser, v interface{}) error {
	if v != nil {
		var err error
		err = json.NewDecoder(body).Decode(v)
		return err
	}

	return nil
}

func (c *Client) encodeRequestBody(body interface{}) (io.ReadWriter, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		} else {
			return buf, nil
		}

	} else {
		return nil, nil
	}
}

func (er *ErrorResponse) Error() string {
	path, _ := url.QueryUnescape(er.Response.Request.URL.Opaque)
	return fmt.Sprintf("%v %v: %d %v", er.Response.Request.Method, path, er.Response.StatusCode, er.Message)
}

// CheckResponse checks the API response for errors and returns them if present
func CheckResponse(r *http.Response) error {
	if r.StatusCode < 300 {
		return nil
	}

	errorResponse := &ErrorResponse{
		Response: r,
	}

	data, err := io.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}

	return errorResponse
}
