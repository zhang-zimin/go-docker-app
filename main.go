package main

import (
	"fmt"
	"log"
	"net/http"
)

// 定义一个处理函数，用于处理根路径的HTTP请求
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 记录访问信息
	log.Printf("访问信息: 方法 %s, 路径 %s, 客户端地址 %s", r.Method, r.URL.Path, r.RemoteAddr)
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// 注册处理函数到根路径
	http.HandleFunc("/", helloHandler)

	// 启动HTTP服务器，监听8080端口
	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("Error starting server:", err)
	}
}
