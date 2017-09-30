package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

func publicIpController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, getPublicIp()) //这个写入到w的是输出到客户端的
}

func getPublicIp() string {
	resp, err := http.Get("http://47.52.66.195:56667")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	// 正则匹配出IP
	return string(body)
}


func main() {
	http.HandleFunc("/", publicIpController) //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("Server ERROR: ", err)
	}
}
