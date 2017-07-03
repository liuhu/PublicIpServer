package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"regexp"
)

func publicIpController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, getPublicIp()) //这个写入到w的是输出到客户端的
}

func getPublicIp() string {
	resp, err := http.Get("http://1212.ip138.com/ic.asp")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	// 正则匹配出IP
	reg, err := regexp.Compile(`((?:(?:25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\.){3}(?:25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d))))`)
	return reg.FindString(string(body))
}


func main() {
	http.HandleFunc("/", publicIpController) //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("Server ERROR: ", err)
	}
}
