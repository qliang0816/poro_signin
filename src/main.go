package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"net/url"
	"github.com/unknwon/goconfig"
	"log"
)

func main() {
	// 读取配置文件
	conf, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件：%s", err)
	}
	login_url, _ := conf.GetValue("user1", "login_url")
	email, _ := conf.GetValue("user1", "email")
	passwd, _ := conf.GetValue("user1", "passwd")
	//请求登录
	resp, err := http.PostForm(login_url, url.Values{"email": {email}, "passwd": {passwd}})
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(string(body))
}
