package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/webdav"
)

func logRequest(req *http.Request, err error) {
	if req != nil {
		fmt.Printf("%s [%v] %q %q\n", req.RemoteAddr, time.Now().Format(time.RFC3339), req.Method, req.RequestURI)
	}
}

func main() {
	var addr string
	flag.StringVar(&addr, "addr", ":80", "listen address")
	flag.Parse()

	fs := webdav.NewMemFS()
	ls := webdav.NewMemLS()

	handler := &webdav.Handler{
		FileSystem: fs,
		LockSystem: ls,
		Logger:     logRequest,
	}

	log.Fatal(http.ListenAndServe(addr, handler))
}
