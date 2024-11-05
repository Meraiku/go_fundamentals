package server

import (
	"fmt"
	"net/http"
	"strings"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		write(w, []byte("Method not allowed"))
		return
	}

	write(w, []byte("Hello world"))
}

func (api *Api) handleSearch(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	search := strings.ReplaceAll(query.Get("q"), " ", "+")

	data, err := api.client.Get(fmt.Sprintf("/search?q=%s", search))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		write(w, []byte(err.Error()))
		return
	}

	write(w, data)
}
