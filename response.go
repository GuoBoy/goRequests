package goRequests

import (
	"encoding/json"
	"io"
)

// Response Session响应
type Response struct {
	Data []byte
}

// NewResponse 新建响应对象
func NewResponse(reader io.ReadCloser) *Response {
	dt, _ := io.ReadAll(reader)
	return &Response{Data: dt}
}

// Bytes 获取bytes
func (s *Response) Bytes() []byte {
	return s.Data
}

// 获取字符串
func (s *Response) String() string {
	return string(s.Data)
}

// Json 获取json
func (s *Response) Json(result any) error {
	return json.Unmarshal(s.Data, &result)
}
