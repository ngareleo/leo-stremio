package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type MiddleWare func(http.Handler) http.Handler


type CustomResponseWriter struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func NewCustomResponseWriter(w http.ResponseWriter) *CustomResponseWriter {
	return &CustomResponseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK, 
		body:           new(bytes.Buffer),
	}
}

func (rw *CustomResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *CustomResponseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b) 
	return rw.ResponseWriter.Write(b)
}



func loggingMiddleware (next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			customRW := NewCustomResponseWriter(w)
			
			slog.Info("incoming request", "path", r.URL.Path)
			next.ServeHTTP(customRW, r)

			if customRW.statusCode != 200 {
				slog.Error("error during request", "message", customRW.body.String())
			}
		})
	}

func BootServer(volume Volume) {

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, volume)
	})

	router.HandleFunc("/index-section", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/sections/index-section.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, volume)
	})

	router.HandleFunc("/open/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Within answer block")
		vars := mux.Vars(r)
		id, found := vars["id"]
		if !found {
			http.Error(w, "couldn't find the media entry you've clicked", http.StatusInternalServerError)
			return
		}

		intVal, err := strconv.Atoi(id);

		if err != nil {
			http.Error(w, "received invalid id; expected numeral", http.StatusBadRequest)
			return
		}

		file, notFound := volume.FindFileById(int(intVal))

		if notFound {
			http.Error(w, "file not found", http.StatusBadRequest)
			return 
		}

		tmpl, err := template.ParseFiles("templates/sections/video-section.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, file)
	})

	router.HandleFunc("/stream/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		_, found := vars["id"]

		if !found {
			http.Error(w, "Missing ID Param", http.StatusBadRequest)
			return
		}
		// open up a connections
	})

	router.Use(loggingMiddleware)

	s := &http.Server{
		Addr:         ":3000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router,
	}

	fmt.Println("Server listening on http://127.0.0.1:3000")
	log.Fatal(s.ListenAndServe())
}


