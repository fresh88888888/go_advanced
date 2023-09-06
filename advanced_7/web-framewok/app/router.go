package app

import (
	url2 "net/url"
	"regexp"
	"strings"
)

type Router struct {
	routes []*RoutePath
}

type RoutePath struct {
	regex  *regexp.Regexp
	method string
	h      HandleFunc
}

// NewRouter returns new router instance
func NewRouter() *Router {
	r := Router{
		routes: make([]*RoutePath, 0),
	}
	return &r
}

func newRouter() *RoutePath {
	route := RoutePath{}
	return &route
}

func (rt *Router) MatchRoute(url string, method string) (*RoutePath, error) {
	url = url2.QueryEscape(url)
	if !strings.HasSuffix(url, "%2F") {
		url += "%2F"
	}
	url = strings.Replace(url, "%2F", "/", -1)
	for _, route := range rt.routes {
		matched, _ := regexp.MatchString(route.regex.String(), url)
		if matched {
			return route, nil
		}
	}

	return nil, nil
}

func (r *Router) Route(pattern string, method string, handler HandleFunc) *RoutePath {
	route := newRouter()
	route.regex = parsePattern(pattern)
	route.h = handler
	route.method = method

	return route
}

func parsePattern(pattern string) *regexp.Regexp {
	pattern = url2.QueryEscape(pattern)
	if !strings.HasSuffix(pattern, "%2F") {
		pattern += "%2F"
	}
	pattern = strings.Replace(pattern, "%2F", "/", -1)
	regex := regexp.MustCompile(pattern)

	return regex
}
