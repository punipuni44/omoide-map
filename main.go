package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func main() {
	// 静的ファイルの配信
	fs := http.FileServer(http.Dir(filepath.Join(".", "static")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// ルートパスで index.html を返す
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, filepath.Join(".", "static", "index.html"))
	})

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
