package server

import (
	"go_final_project/go_final_project/pkg/api"
	//"go_final_project/go_final_project/pkg/db"
	//"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	DefaultPort = "7540"
	WebDir      = "./web"
)

func StartServer(port string, webDir string) error {
	api.Init()
	fs := CustomFileServer(webDir)
	http.Handle("/", fs)

	return http.ListenAndServe(":"+port, nil)
}

func CustomFileServer(root string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Корректируем путь к файлу
		path := filepath.Join(root, r.URL.Path)

		if r.URL.Path == "/" {
			path = filepath.Join(root, "index.html")
		}

		if _, err := os.Stat(path); os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}

		http.ServeFile(w, r, path)
	})
}
