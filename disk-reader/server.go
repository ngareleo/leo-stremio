package main

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type MiddleWare func(http.Handler) http.Handler

func logger() MiddleWare {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			slog.Info("incoming request", "path", r.URL.Path)
			h.ServeHTTP(w, r)
		})
	}
}

func BootServer(volume Volume) {

	httpMux := mux.NewRouter()

	httpMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, volume)
	})

	httpMux.HandleFunc("/index-section", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/sections/index-section.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, volume)
	})

	httpMux.HandleFunc("/open/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		_, found := vars["id"]
		if !found {
			http.Error(w, "Couldn't find the media entry you've clicked", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("templates/sections/video-section.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, volume)
	})

	httpMux.HandleFunc("/stream/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		_, found := vars["id"]

		if !found {
			http.Error(w, "Missing ID Param", http.StatusBadRequest)
			return
		}
		// open up a connections
	})

	s := &http.Server{
		Addr:         ":3000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      logger()(httpMux),
	}

	fmt.Println("Server listening on http://127.0.0.1:3000")
	err := s.ListenAndServe()

	if err != nil {
		slog.Error("error booting server", "message", err.Error())
	}
}


