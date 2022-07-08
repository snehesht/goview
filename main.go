package main

import (
	"log"
	"net/http"

	_ "github.com/snehesht/goview/statik"

	"github.com/rakyll/statik/fs"
	"github.com/webview/webview"
)


func main() {
	go start()
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Create React App")
	w.SetSize(1280, 720, webview.HintNone)
	w.Navigate("http://localhost:8181/index.html")
	w.Run()
}

func start() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	http.ListenAndServe(":8181", nil)
}