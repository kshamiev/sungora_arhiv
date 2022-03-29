package request

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Подготовка данных для отправки запроса
func (r *Request) requestSendData(method, query string, requestBody interface{}) (string, *bytes.Buffer, error) {
	var (
		err  error
		body = new(bytes.Buffer)
		data = []byte("")
	)

	check := r.Header.Get("Content-Type")
	if check == "" {
		check = headerTypeJSON
	}

	if method == http.MethodPost || method == http.MethodPut {
		switch strings.Split(check, ";")[0] {
		case strings.Split(headerTypeJSON, ";")[0]:
			data, err = json.Marshal(requestBody)
		case strings.Split(headerTypeXML, ";")[0]:
			data, err = xml.Marshal(requestBody)
		case strings.Split(headerTypeFormURLEncoded, ";")[0]:
			if p, ok := requestBody.(map[string]interface{}); ok {
				data = []byte(uriParamsCompile(p))
			}
		}

		if err != nil {
			return "", nil, err
		}

		if _, err = body.Write(data); err != nil {
			return "", nil, err
		}
	}

	if p, ok := requestBody.(map[string]interface{}); ok {
		query += "?" + uriParamsCompile(p)
	}

	return query, body, nil
}

// Разбор данных ответа на запрос
func (r *Request) requestResiveData(res *http.Response, responseBody interface{}) (err error) {
	check := res.Header.Get("Accept")
	if check == "" {
		check = res.Header.Get("Content-Type")
		if check == "" {
			check = headerTypeJSON
		}
	}
	if responseBody != nil {
		switch strings.Split(check, ";")[0] {
		case strings.Split(headerTypeJSON, ";")[0]:
			err = json.Unmarshal(r.ResponseBody, responseBody)
		case strings.Split(headerTypeXML, ";")[0]:
			err = xml.Unmarshal(r.ResponseBody, responseBody)
		}
	}
	return
}

// uriParamsCompile
func uriParamsCompile(postData map[string]interface{}) string {
	q := &url.Values{}

	for k, v := range postData {
		switch v1 := v.(type) {
		case uint64:
			q.Add(k, strconv.FormatUint(v1, 10))
		case int64:
			q.Add(k, strconv.FormatInt(v1, 10))
		case int:
			q.Add(k, strconv.Itoa(v1))
		case float64:
			q.Add(k, strconv.FormatFloat(v1, 'f', -1, 64))
		case bool:
			q.Add(k, strconv.FormatBool(v1))
		case string:
			q.Add(k, v1)
		}
	}

	return q.Encode()
}
