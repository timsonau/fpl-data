package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const BASE_FPL_URL = "https://fantasy.premierleague.com/api/"

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/general", getGeneralInfoHandler).Methods("GET")
	router.HandleFunc("/manager", getMyManagerProfileHandler).Methods("GET")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getGeneralInfoHandler(w http.ResponseWriter, r *http.Request) {
	data, err := getGeneralInfo()
	if err != nil {
		http.Error(w, "Failed to fetch API data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(data)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func getGeneralInfo() (map[string]interface{}, error) {
	resp, err := http.Get(BASE_FPL_URL + "/bootstrap-static")
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

func getMyManagerProfileHandler(w http.ResponseWriter, r *http.Request) {
	data, err := getMyManagerProfile()
	if err != nil {
		http.Error(w, "Failed to fetch API data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(data)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
func getMyManagerProfile() (map[string]interface{}, error) {
	managerId := os.Getenv("FPL_MANAGER_ID")
	url := BASE_FPL_URL + fmt.Sprintf("entry/%s/history", managerId)
	fmt.Print(url)
	resp, err := http.Get(url)
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
