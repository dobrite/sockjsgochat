package main

import (
	"github.com/igm/sockjs-go/sockjs"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.Handle("/echo/", sockjs.NewHandler("/echo", sockjs.DefaultOptions, echoHandler))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func echoHandler(session sockjs.Session) {
	for {
		if msg, err := session.Recv(); err == nil {
			session.Send(msg)
			continue
		}
		break
	}
}

func StaticServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "YO!\n")
}

func Index(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	contents, err := ioutil.ReadFile("static/index.html")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(contents)
}
