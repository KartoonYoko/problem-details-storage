package http_server

import (
	"net/http"
	"path"
	"path/filepath"
)

func handleWebApp(w http.ResponseWriter, r *http.Request) {
	webPath := "../../web/dist/web/browser"
	dir, file := path.Split(r.RequestURI)
	ext := filepath.Ext(file)

	if file == "" || ext == "" {
		http.ServeFile(w, r, webPath+"/index.html")
	} else {
		http.ServeFile(w, r, webPath+path.Join(dir, file))
	}
}
