package goRequests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// HttpSession 客户端请求封装
type HttpSession struct {
	Client *http.Client
}

// NewSession 新建session
func NewSession() *HttpSession {
	return &HttpSession{Client: http.DefaultClient}
}

// Get Get请求
func (h *HttpSession) Get(url_ string, headers map[string]string) (*Response, error) {
	parsedUrl, err := url.Parse(url_)
	if err != nil {
		return nil, err
	}
	header := make(http.Header)
	for k, v := range headers {
		header.Set(k, v)
	}
	request := &http.Request{
		Method: http.MethodGet,
		URL:    parsedUrl,
		Header: header,
	}
	resp, err := h.Client.Do(request)
	if err != nil {
		return nil, err
	}
	return NewResponse(resp), nil
}

// Post Post请求
func (h *HttpSession) Post(url_ string, headers map[string]string, body string) (*Response, error) {
	parsedUrl, err := url.Parse(url_)
	if err != nil {
		return nil, err
	}
	header := make(http.Header)
	for k, v := range headers {
		header.Set(k, v)
	}
	header.Set("Content-Type", "application/x-www-form-urlencoded")
	request := &http.Request{
		Method: http.MethodPost,
		URL:    parsedUrl,
		Header: header,
		Body:   io.NopCloser(bytes.NewReader([]byte(body))),
	}
	resp, err := h.Client.Do(request)
	if err != nil {
		return nil, err
	}
	return NewResponse(resp), nil
}

// PostByteBody Post请求， with byte body
func (h *HttpSession) PostByteBody(url_ string, headers map[string]string, body []byte) (*Response, error) {
	parsedUrl, err := url.Parse(url_)
	if err != nil {
		return nil, err
	}
	header := make(http.Header)
	for k, v := range headers {
		header.Set(k, v)
	}
	header.Set("Content-Type", "application/x-www-form-urlencoded")
	request := &http.Request{
		Method: http.MethodPost,
		URL:    parsedUrl,
		Header: header,
		Body:   io.NopCloser(bytes.NewReader(body)),
	}
	resp, err := h.Client.Do(request)
	if err != nil {
		return nil, err
	}
	return NewResponse(resp), nil
}

// PostForm PostForm请求
func (h *HttpSession) PostForm(url_ string, headers map[string]string, body url.Values) (*Response, error) {
	parsedUrl, err := url.Parse(url_)
	if err != nil {
		return nil, err
	}
	header := make(http.Header)
	for k, v := range headers {
		header.Set(k, v)
	}
	header.Set("Content-Type", "multipart/form-data")
	request := &http.Request{
		Method: http.MethodPost,
		URL:    parsedUrl,
		Header: header,
		Body:   io.NopCloser(bytes.NewReader([]byte((body.Encode())))),
	}
	resp, err := h.Client.Do(request)
	if err != nil {
		return nil, err
	}
	return NewResponse(resp), nil
}

// PostJson PostJson请求
func (h *HttpSession) PostJson(url_ string, headers map[string]string, body JsonT) (*Response, error) {
	parsedUrl, err := url.Parse(url_)
	if err != nil {
		return nil, err
	}
	header := make(http.Header)
	for k, v := range headers {
		header.Set(k, v)
	}
	header.Set("Content-Type", "application/json")
	bd, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request := &http.Request{
		Method: http.MethodPost,
		URL:    parsedUrl,
		Header: header,
		Body:   io.NopCloser(bytes.NewReader(bd)),
	}
	resp, err := h.Client.Do(request)
	if err != nil {
		return nil, err
	}
	return NewResponse(resp), nil
}

// Put Put请求
func (h *HttpSession) Put(url_ string, headers map[string]string, body string) (*Response, error) {
	parsedUrl, err := url.Parse(url_)
	if err != nil {
		return nil, err
	}
	header := make(http.Header)
	for k, v := range headers {
		header.Set(k, v)
	}
	request := &http.Request{
		Method: http.MethodPut,
		URL:    parsedUrl,
		Header: header,
		Body:   io.NopCloser(bytes.NewReader([]byte(body))),
	}
	resp, err := h.Client.Do(request)
	if err != nil {
		return nil, err
	}
	return NewResponse(resp), nil
}

// Delete Delete请求
func (h *HttpSession) Delete(url_ string, headers map[string]string) (*Response, error) {
	parsedUrl, err := url.Parse(url_)
	if err != nil {
		return nil, err
	}
	header := make(http.Header)
	for k, v := range headers {
		header.Set(k, v)
	}
	request := &http.Request{
		Method: http.MethodDelete,
		URL:    parsedUrl,
		Header: header,
	}
	resp, err := h.Client.Do(request)
	if err != nil {
		return nil, err
	}
	return NewResponse(resp), nil
}

// Options Options请求
func (h *HttpSession) Options(url_ string, headers map[string]string) (*Response, error) {
	parsedUrl, err := url.Parse(url_)
	if err != nil {
		return nil, err
	}
	header := make(http.Header)
	for k, v := range headers {
		header.Set(k, v)
	}
	request := &http.Request{
		Method: http.MethodOptions,
		URL:    parsedUrl,
		Header: header,
	}
	resp, err := h.Client.Do(request)
	if err != nil {
		return nil, err
	}
	return NewResponse(resp), nil
}
