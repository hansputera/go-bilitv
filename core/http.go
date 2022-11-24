package core

import (
	"go-bilitv/config"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/exp/maps"
)

type BiliHttp struct {
	Base           *url.URL
	DefaultHeaders map[string]string
	httpClient     *http.Client
}

func (h *BiliHttp) CombineHeaders(headers map[string]string) *http.Header {
	maps.Copy(headers, h.DefaultHeaders)
	headers_net := http.Header{}

	for head, headValue := range headers {
		headers_net.Add(head, headValue)
	}

	return &headers_net
}

func (h *BiliHttp) Request(method string, path string, queries *url.Values, postform *url.Values, body *io.ReadCloser, additional_headers *map[string]string) *http.Request {
	tmpUrl := *h.Base
	tmpUrl.Path = path
	if queries != nil {
		tmpUrl.RawQuery = queries.Encode()
	}

	headers := &http.Header{}
	if additional_headers != nil {
		headers = h.CombineHeaders(*additional_headers)
	} else {
		headers = h.CombineHeaders(map[string]string{})
	}

	r := http.Request{
		URL:    &tmpUrl,
		Header: *headers,
		Method: method,
	}

	if postform != nil {
		r.PostForm = *postform
	} else if body != nil {
		r.Body = *body
	}

	return &r
}

func (h *BiliHttp) Send(request *http.Request) *http.Response {
	res, err := h.httpClient.Do(request)
	if err != nil {
		res = &http.Response{}
		res.Status = err.Error()
		res.StatusCode = 505
	} else {
		defer res.Body.Close()
	}

	return res
}

func (h *BiliHttp) Get(path string, queries url.Values) *http.Response {
	return h.Send(h.Request("GET", path, &queries, nil, nil, nil))
}

func (h *BiliHttp) PostForm(path string, fields url.Values) *http.Response {
	return h.Send(h.Request("POST", path, nil, &fields, nil, nil))
}

func (h *BiliHttp) Delete(path string) *http.Response {
	return h.Send(h.Request("DELETE", path, nil, nil, nil, nil))
}

func (h *BiliHttp) PostPayload(path string, body io.ReadCloser) *http.Response {
	return h.Send(h.Request("POST", path, nil, nil, &body, nil))
}

func (h *BiliHttp) PostStr(path string, str string) *http.Response {
	p := io.NopCloser(strings.NewReader(str))

	return h.PostPayload(path, p)
}

func GetBiliHttp(httpType int, webConfig *config.WebConfig) *BiliHttp {
	if webConfig == nil {
		webConfig = config.GetWebConfig()
	}

	switch httpType {
	case 0:
		uri, err := url.Parse(webConfig.ApiGateway)
		if err != nil {
			return nil
		}

		return &BiliHttp{
			Base: uri,
			DefaultHeaders: map[string]string{
				"Origin":  "https://www.bilibili.tv",
				"Referer": "https://www.bilibili.tv/",
			},
			httpClient: &http.Client{
				Timeout: 5 * time.Second,
			},
		}
	case 1:
		uri, err := url.Parse(webConfig.WebUrl)
		if err != nil {
			return nil
		}

		return &BiliHttp{
			Base: uri,
			DefaultHeaders: map[string]string{
				"Origin":  "https://www.bilibili.tv",
				"Referer": "https://www.bilibili.tv/",
			},
			httpClient: &http.Client{
				Timeout: 15 * time.Second,
			},
		}
	}

	return nil
}
