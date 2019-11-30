package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

// Index home
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	html, err := os.Open("index.html")
	if err != nil {
		fmt.Fprintf(w, "Error %d", err)
		return
	}
	fi, err := html.Stat()
	if err != nil {
		fmt.Fprintf(w, "Error %d", err)
		return
	}

	r.Header.Set("Content-Length", fmt.Sprint(fi.Size()))
	r.Header.Set("Content-Type", "text/html; charset=utf-8")
	if _, err = io.Copy(w, html); err != nil {
		fmt.Fprintf(w, "Error %d", err)
	}
	html.Close()
}

// ImageHandler image handle
func ImageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// queryValues := r.URL.Query()

	output := fmt.Sprintf("./image/%s_%s.png", ps.ByName("name"), ps.ByName("size"))

	reqImg, err := os.Open(output)
	if err != nil {
		fmt.Fprintf(w, "Error %d", err)
		return
	}
	fi, err := reqImg.Stat()
	if err != nil {
		fmt.Fprintf(w, "Error %d", err)
		return
	}

	r.Header.Set("Content-Length", fmt.Sprint(fi.Size()))
	r.Header.Set("Content-Type", "image/png")

	if _, err = io.Copy(w, reqImg); err != nil {
		fmt.Fprintf(w, "Error %d", err)
	}
	reqImg.Close()
}

func main() {
	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/image/:name/:size", ImageHandler)
	fmt.Println("Listen 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
