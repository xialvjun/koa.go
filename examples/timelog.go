package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/xialvjun/koa.go/koa"
)

type StringReader struct {
	current, length int
}

func (sr *StringReader) Read(p []byte) (n int, err error) {
	if sr.current < sr.length {
		n, err = rand.Read(p)
		sr.current += n
		return
	}
	return 0, io.EOF
}

func timeM(r *http.Request, next koa.Next) interface{} {
	fmt.Println("begin\t" + r.URL.String())
	begin := time.Now()
	time.Sleep(time.Duration(5 * time.Second))
	next()
	end := time.Now()
	ms := end.Sub(begin)
	fmt.Println(begin.String() + "\t" + end.String() + "\t" + ms.String())
	fmt.Println()
	return &StringReader{0, 100}
}

func logM(r *http.Request, next koa.Next) interface{} {
	fmt.Println(r.Method + "\t" + r.URL.String())
	next()
	return nil
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

// func abc(fn func() interface{}) {}

// func ab() string { return "" }

// var a = abc(ab)
