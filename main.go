package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 关键修复：增加 *http.Request 参数（习惯上命名为 r）
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, My SaaS is running on Docker!")
	})

	fmt.Println("Server starting at :8080")
	http.ListenAndServe(":8080", nil)
}
