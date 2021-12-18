package request

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	headerTypeFormURLEncoded = "application/x-www-form-urlencoded; charset=utf-8"
	headerTypeJSON           = "application/json; charset=utf-8"
	headerTypeXML            = "text/xml; charset=utf-8"
)

// Структура для работы с исходящими запросами
type Request struct {
	url          string
	ResponseBody []byte
	Header       http.Header
	Transport    *http.Transport
}

// New Функционал по работе с исходящими запросами к внешним ресурсам
func New(link string) *Request {
	return &Request{
		url:    link,
		Header: http.Header{},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,            // ignore expired SSL certificates
				MinVersion:         tls.VersionTLS13, //
			},
		},
	}
}

func (r *Request) ContentTypeFormURLEncoded(link string) *Request {
	r.Header.Add("Content-Type", headerTypeFormURLEncoded)
	return r
}

func (r *Request) ContentTypeJSON(link string) *Request {
	r.Header.Add("Content-Type", headerTypeJSON)
	return r
}

func (r *Request) ContentTypeXML(link string) *Request {
	r.Header.Add("Content-Type", headerTypeXML)
	return r
}

// AuthorizationBasic
func (r *Request) AuthorizationBasic(login, passw string) {
	r.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(login+":"+passw)))
}

func (r *Request) Transports(proxy string) {
	if proxy != "" {
		proxyURL, _ := url.Parse(proxy)
		r.Transport.Proxy = http.ProxyURL(proxyURL)
	}
}

// GET запрос
func (r *Request) GET(uri string, requestBody, responseBody interface{}) (*http.Response, error) {
	return r.request(http.MethodGet, uri, requestBody, responseBody)
}

// POST запрос
func (r *Request) POST(uri string, requestBody, responseBody interface{}) (*http.Response, error) {
	return r.request(http.MethodPost, uri, requestBody, responseBody)
}

// PUT запрос
func (r *Request) PUT(uri string, requestBody, responseBody interface{}) (*http.Response, error) {
	return r.request(http.MethodPut, uri, requestBody, responseBody)
}

// DELETE запрос
func (r *Request) DELETE(uri string, requestBody, responseBody interface{}) (*http.Response, error) {
	return r.request(http.MethodDelete, uri, requestBody, responseBody)
}

// OPTIONS запрос
func (r *Request) OPTIONS(uri string, requestBody, responseBody interface{}) (*http.Response, error) {
	return r.request(http.MethodOptions, uri, requestBody, responseBody)
}

func (r *Request) request(method, uri string, requestBody, responseBody interface{}) (*http.Response, error) {
	// Данные запроса
	query, body, err := r.requestSendData(method, r.url+uri, requestBody)
	if err != nil {
		return nil, err
	}

	// Запрос
	request, err := http.NewRequest(method, query, body)
	if err != nil {
		return nil, err
	}
	request.Header = r.Header

	// Транспорт // r.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: false}}
	c := http.Client{Transport: r.Transport}

	// Ответ
	response, err := c.Do(request)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	if r.ResponseBody, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}

	// Данные ответа
	err = r.requestResiveData(response, responseBody)

	if response.StatusCode >= http.StatusBadRequest {
		err = fmt.Errorf("%s:[%d]:%s", method, response.StatusCode, query)
	}

	return response, err
}
