package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := getGeneralInfo()
		if err != nil {
			http.Error(w, "Failed to fetch API data", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data) // Send JSON response
	})

	http.ListenAndServe(":8000", router)
}

func getGeneralInfo() (map[string]interface{}, error) {
	resp, err := http.Get("https://fantasy.premierleague.com/api/bootstrap-static/")
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data) // Convert JSON to Go map
	if err != nil {
		return nil, err
	}

	return data, nil
}
