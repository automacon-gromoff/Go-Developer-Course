package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var l *log.Logger

func main() {
	mux := http.NewServeMux()
	l = log.New(os.Stdout, "", 0)
	mux.HandleFunc("/hello", hello)
	logHandler := logMiddleware()
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: authHandler(logHandler(mux)),
	}
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalln(fmt.Errorf("Не удалось запустить сервер: %w", err))
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	msg := "Hello, Go!"
	l.Println("resp:", msg)
	fmt.Fprint(res, msg)
}

func authHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xId := r.Header.Get("x-my-app-id")
		if xId != "my_secret" {
			http.Error(w, "пользователь не авторизован", http.StatusUnauthorized)
			l.Println("err: unauthorized")
			return
		}
		h.ServeHTTP(w, r)
	})
}

func logMiddleware() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Println("url:", r.URL)
			h.ServeHTTP(w, r)
		})
	}
}
