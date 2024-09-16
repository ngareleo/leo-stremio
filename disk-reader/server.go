package main

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"time"
)

type Dir struct {
	Files []string
}

type MiddleWare func(http.Handler) http.Handler

func BootServer(dir Dir) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, dir)
	})


	mux.HandleFunc("/stream/{name}", func(w http.ResponseWriter, r *http.Request) {})

	s := &http.Server{
		Addr:         ":3000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      logger()(mux),
	}

	fmt.Println("Server listening on http://127.0.0.1:3000")
	s.ListenAndServe()
}

func logger() MiddleWare {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			slog.Info("incoming request", "path", r.URL.Path)
			h.ServeHTTP(w, r)
		})
	}
}
