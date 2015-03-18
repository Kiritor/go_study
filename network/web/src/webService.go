package main

import (
	"net/http"
)

func main() {
	//设置静态路由
	http.Handle("/css/", http.FileServer(http.Dir("template")))
	http.Handler("/js/", http.FileServer(http.Dir("template")))

}
