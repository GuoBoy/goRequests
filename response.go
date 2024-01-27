package goRequests

import (
	"encoding/json"
	"io"
	"net/http"
)

// Response Session响应
type Response struct {
	ParsedBytes []byte
	Reader      io.ReadCloser
	Parsed      bool
	Header      http.Header
}

// NewResponse 新建响应对象
func NewResponse(resp *http.Response) *Response {
	return &Response{Reader: resp.Body, Header: resp.Header.Clone()}
}

// Bytes 获取bytes
func (s *Response) Bytes() []byte {
	if !s.Parsed {
		s.ParsedBytes, _ = io.ReadAll(s.Reader)
		s.Parsed = true
	}
	return s.ParsedBytes
}

// 获取字符串
func (s *Response) String() string {
	if !s.Parsed {
		return string(s.Bytes())
	}
	return string(s.ParsedBytes)
}

// Json 获取json
func (s *Response) Json(result any) error {
	if !s.Parsed {
		s.Bytes()
	}
	return json.Unmarshal(s.ParsedBytes, &result)
}
