package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var l *log.Logger

func main() {
	logFile, err := os.OpenFile("./lesson17/task3/httpServer.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		l.Fatalln(fmt.Errorf("ошибка открытия файла логирования: %s", err))
	}
	defer logFile.Close()
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	l = log.New(logFile, "", log.LstdFlags)
	logHandler := logMiddleware()
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: authHandler(logHandler(mux)),
	}
	if err := httpServer.ListenAndServe(); err != nil {
		l.Fatalln(fmt.Errorf("Не удалось запустить сервер: %w", err))
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
