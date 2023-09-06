package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/context/ctxhttp"
)

type headers []string

type config struct {
	httpMethod      string
	body            string
	followRedirects bool
	httpHeaders     []string
	saveOutput      bool
	outputFile      string
}

func isHTTPRedirect(resp *http.Response) bool {
	return resp.StatusCode > 299 && resp.StatusCode < 400
}

// Set implements flag.Value.
func (h *headers) Set(v string) error {
	*h = append(*h, v)
	return nil
}

// String implements flag.Value.
func (h *headers) String() string {
	var o []string
	for _, v := range *h {
		o = append(o, fmt.Sprintf("-H %s", v))
	}

	return strings.Join(o, " ")
}

var (
	// Command line flags config
	httpMethod      string
	body            string
	followRedirects bool
	httpHeaders     headers
	saveOutput      bool
	outputFile      string

	// number of redirects followed
	redirectsFollowedCount int
)

const (
	defaultUrlScheme = "http"
	maxRedirects     = 10
)

func init() {
	flag.Var(&httpHeaders, "H", "set HTTP headers")
	flag.StringVar(&httpMethod, "X", "GET", "HTTP method to use")
	flag.StringVar(&body, "d", "", "the body of a POST or PUT request")
	flag.BoolVar(&followRedirects, "L", false, "follow redirects")
	flag.BoolVar(&saveOutput, "0", false, "save body as remote filename")
	flag.StringVar(&outputFile, "o", "", "output file for body")
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	if (httpMethod == http.MethodPost || httpMethod == http.MethodPut) && body == "" {
		log.Fatal("httpMethod: must supply body using -d flag")
	}

	url := parseUrl(args[0])
	c := config{
		body:            body,
		followRedirects: followRedirects,
		httpHeaders:     httpHeaders,
		saveOutput:      saveOutput,
		outputFile:      outputFile,
	}

	performRequest(context.Background(), url, &c)
}

func parseUrl(urlStr string) *url.URL {
	if !strings.Contains(urlStr, "://") && !strings.HasPrefix(urlStr, "//") {
		urlStr = fmt.Sprintf("//%v", urlStr)
	}

	url, err := url.Parse(urlStr)
	if err != nil {
		log.Fatalf("parseUrl: could not parse url %q: %v", urlStr, err)
	}

	if url.Scheme == "" {
		url.Scheme = defaultUrlScheme
	}

	return url
}

func performRequest(ctx context.Context, url *url.URL, c *config) {
	req := newRequest(c.httpMethod, url, c.body, c.httpHeaders)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := ctxhttp.Do(ctx, client, req)
	if err != nil {
		log.Fatalf("executeRequest: failed to read response: %v", err)
	}

	err = readResponseBody(req, resp, c)
	if err != nil {
		log.Fatalf("executeRequest: failed to read body: %v", err)
	}

	defer resp.Body.Close()

	if c.followRedirects && isHTTPRedirect(resp) {
		loc, err := resp.Location()
		if err != nil {
			if err == http.ErrNoLocation {
				log.Fatalf("redirect: unable to follow redirect")
			}
		}

		redirectsFollowedCount++
		if redirectsFollowedCount > maxRedirects {
			log.Fatalf("redirect: maximum number of redirects followed")
		}

		performRequest(ctx, loc, c)
	}
}

func newRequest(method string, url *url.URL, body string, headers headers) *http.Request {
	req, err := http.NewRequest(method, url.String(), strings.NewReader(body))
	if err != nil {
		log.Fatalf("newRequest: unable to create request: %v", err)
	}

	for _, header := range headers {
		k, v := parseHeader(header)
		req.Header.Add(k, v)
	}

	return req
}

func parseHeader(h string) (string, string) {
	i := strings.IndexRune(h, ':')
	if i == -1 {
		log.Fatalf("parserHeader: '%s' has invalid format", h)
	}

	return strings.TrimRight(h[:i], " "), strings.TrimLeft(h[:i], " :")
}

func readResponseBody(req *http.Request, resp *http.Response, c *config) error {
	if isHTTPRedirect(resp) {
		return nil
	}

	var out io.Writer
	if c.saveOutput || c.outputFile != "" {
		filename := outputFile
		if filename == "" {
			tmpFile, err := os.CreateTemp(".", req.URL.Path)
			if err != nil {
				return errors.New("readResponseBody: unable to create output file")
			}
			filename = tmpFile.Name()
		}

		f, err := os.Create(filename)
		if err != nil {
			return errors.New("readResponseBody: unable to create output file")
		}
		defer f.Close()
		out = f
	} else {
		out = os.Stdout
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New("readResponseBody: failed to read response body")
	}

	out.Write(respBytes)

	return nil
}
