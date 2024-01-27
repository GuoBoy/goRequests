package example

import (
	"fmt"
	"github.com/GuoBoy/goRequests"
	"log"
)

func Run() {
	session := goRequests.NewSession()
	accessApi := "http://localhost:8080"
	headers := map[string]string{
		"User-Agent": "Mozilla Chrome 111.1.1",
	}
	body := goRequests.JsonT{"username": "admin", "password": "admin"}
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
