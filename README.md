# GoRequests

> A golang library Based on standard lib, the name comes from that Python requests.

## Design

It's request in a session, you can do:
- get
- post
  - form data
  - json
- put
- delete
- options

## Usage

```go
package main

import (
	"fmt"
	"github.com/GuoBoy/goRequests"
	"log"
)

func main() {
	session := goRequests.NewSession()
	accessApi := "http://localhost:8080"
	headers := map[string]string{
		"User-Agent": "Mozilla Chrome 111.1.1",
	}
	body := fmt.Sprintf(`{"username": "admin", "password": "admin"}`)
	accessResp, err := session.PostJson(accessApi, headers, body)
	if err != nil {
		log.Fatal("获取access_token失败", err)
	}
	// 获取access响应结构体
	type accessRespModel struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
	var result accessRespModel
	err = accessResp.Json(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
```

## Get more

TODO