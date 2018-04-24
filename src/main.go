package main

import (
	"net/http"
	"io/ioutil"
	"net/url"
	"github.com/unknwon/goconfig"
	"log"
	"encoding/json"
	"net/http/cookiejar"
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
	// 保持cookie的可用性
	var client http.Client
	jar, _ := cookiejar.New(nil)
	client.Jar = jar
	//请求登录
	resp, err := client.PostForm(login_url, url.Values{"email": {email}, "passwd": {passwd}})
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("%s", err)
	}
	var login_msg map[string]interface{}
	// 解析返回参数是否登录成功
	if err := json.Unmarshal([]byte(string(body)), &login_msg); err == nil {
		if resp.StatusCode == 200 {
			log.Printf(email+" :%s", login_msg["msg"])
			checkin_url, _ := conf.GetValue("user1", "checkin_url")
			resp, _ = client.PostForm(checkin_url, url.Values{})
			if err != nil {
				log.Fatalf("签到失败:%s", err)
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("%s", err)
			}
			//log.Println(string(body))
			// 解析签到信息
			var sign_msg map[string]interface{}
			if err := json.Unmarshal([]byte(string(body)), &sign_msg); err == nil {
				//log.Println(resp.StatusCode)
				if resp.StatusCode == 200 {
					log.Println(sign_msg["msg"])
				} else {
					log.Println("登录失败")
				}
			}
		} else {
			log.Println("登录失败")
		}
	} else {
		log.Fatalf("%s", err)
	}
}
