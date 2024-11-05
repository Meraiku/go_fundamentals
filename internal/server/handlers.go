package server

import (
	"fmt"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	w.Write([]byte("Hello world"))
}

func (api *Api) handleSearch(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	search := query.Get("q")

	data, err := api.client.Get(fmt.Sprintf("/search?q=%s", search))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(data)
}
