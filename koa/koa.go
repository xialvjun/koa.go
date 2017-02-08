package koa

import (
	"net/http"
)

// Next assaf
type Next func()

// Middleware asdasf
type Middleware func(http.ResponseWriter, *http.Request, Next)

// Application asf
type Application struct {
	// http.HandleFunc
	middlewares []Middleware
}

// Use asd
func (app *Application) Use(middleware Middleware) {
	app.middlewares = append(app.middlewares, middleware)
}

// Callback asdasf
func (app *Application) Callback(w http.ResponseWriter, r *http.Request) {
	// app.middlewares
	current, count := 0, len(app.middlewares)
	var next func()
	next = func() {
		if current < count {
			middleware := app.middlewares[current]
			current = current + 1
			middleware(w, r, next)
		}
	}
	next()
	// app.middlewares[current](w, r)
}

// Listen asdsaf
func (app *Application) Listen(addr string) error {
	http.HandleFunc("/", app.Callback)
	return http.ListenAndServe(addr, nil)
}

// func (app *Application) HandleFunc(w http.ResponseWriter, r *http.Request)  {
//   app.Callback(w, r)
// }
