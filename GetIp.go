package main

import (
	"fmt"
	"net/http"
	"log"
	"net"
)

func getRemoteIp(w http.ResponseWriter, r *http.Request) {
	// get client ip address
	ip,_,_ := net.SplitHostPort(r.RemoteAddr)

	// print out the ip address
	fmt.Fprintf(w,ip)
}

func main() {
	http.HandleFunc("/", getRemoteIp) //设置访问的路由
	err := http.ListenAndServe(":56667", nil) //设置监听的端口
	if err != nil {
		log.Fatal("Server ERROR: ", err)
	}
}
