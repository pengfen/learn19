package main

import (
"io"
"log"
"net/http"
)

const form = `<html><body><form action='#' method='post' name='bar'>
<input type='text' name='in'/>
<input type='text' name='in'/>
<input type='submit' value='submit'/>
</form></body></html>`

func SimpleServer(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "<h1>welcome to http world</h1>")
	panic("test test")
}

func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
		case "GET":
			io.WriteString(w, form)
		case "POST":
			request.ParseForm()
			io.WriteString(w, request.Form["in"][1])
			io.WriteString(w, "\n")
			io.WriteString(w, request.FormValue("in"))
	}
}

func main() {
	http.HandlerFunc("/test1", logPanics(SimpleServer))
	http.HandlerFunc("/test2", logPanics(FormServer))
	if err := http.ListenAndServe(":8088", nil); err != nil {

	}
}

func logPanics(handle http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}
		}()
		handle(writer, request)
	}
}
