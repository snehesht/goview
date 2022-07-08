package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/webview/webview"
)

//go:embed app/build/*
var content embed.FS

func main() {
	go start()
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Create React App")
	w.SetSize(1280, 720, webview.HintNone)
	w.Navigate("http://localhost:8181")
	w.Run()
}

func httpHandler() http.Handler {
	fsys := fs.FS(content)
	html, _ := fs.Sub(fsys, "app/build")
	return http.FileServer(http.FS(html))
}

func start() {
	mux := http.NewServeMux()
	mux.Handle("/", httpHandler())

	http.ListenAndServe(":8181", mux)
}