package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request)  {
	v := req.FormValue("q")
	io.WriteString(w, "Do my search: "+v)
}
