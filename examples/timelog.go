package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/xialvjun/koa.go/koa"
)

func timeM(w http.ResponseWriter, r *http.Request, next koa.Next) {
	fmt.Println("begin\t" + r.URL.String())
	begin := time.Now()
	time.Sleep(time.Duration(10000000000))
	next()
	end := time.Now()
	ms := end.Sub(begin)
	fmt.Println(begin.String() + "\t" + end.String() + "\t" + ms.String())
	fmt.Println()
}

func logM(w http.ResponseWriter, r *http.Request, next koa.Next) {
	fmt.Println(r.Method + "\t" + r.URL.String())
	w.Write([]byte("logMB\t"))
	next()
	w.Write([]byte("logMA\t"))
}

// func timeMM(w http.ResponseWriter, r *http.Request) {
// 	begin := time.Now()
// 	time.Sleep(time.Duration(10000000000))
// 	end := time.Now()
// 	ms := end.Sub(begin)
// 	fmt.Println(begin.String() + "\t" + end.String() + "\t" + ms.String())
// }

func main() {
	var app koa.Application
	app.Use(timeM)
	app.Use(logM)
	app.Listen("0.0.0.0:8888")
	// http.HandleFunc("/", timeMM)
	// http.ListenAndServe("0.0.0.0:8888", nil)
}
