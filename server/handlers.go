package server

import (
	"fmt"
	"net/http"

	"github.com/unrolled/render"
)

func errorHandlerJSON(err error, w http.ResponseWriter, status int) {
	r := render.New(render.Options{
		IndentJSON: true,
	})

	json := map[string]string{
		"error": fmt.Sprintf("%s", err),
	}
	r.JSON(w, status, json)
}

func errorHandler(err error, w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404 image not available. ", err)
	}
}
