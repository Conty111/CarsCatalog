package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Car struct {
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   int    `json:"year,omitempty"`
	Owner  People `json:"owner"`
}

type People struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
}

func getInfoHandler(w http.ResponseWriter, r *http.Request) {
	regNum := r.URL.Query().Get("regNum")
	if regNum == "" {
		http.Error(w, "regNum parameter is required", http.StatusBadRequest)
		return
	}

	// Generate or fetch car information based on regNum
	car := generateCarInfo(regNum)

	if car.RegNum == "" {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	responseJSON, err := json.Marshal(car)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func generateCarInfo(regNum string) Car {
	// Here you can implement your logic to generate or fetch car information
	// For simplicity, let's generate random data
	return Car{
		RegNum: regNum,
		Mark:   "Lada",
		Model:  "Vesta",
		Year:   2002,
		Owner: People{
			Name:       "John",
			Surname:    "Doe",
			Patronymic: "Smith",
		},
	}
}

func main() {
	http.HandleFunc("/info", getInfoHandler)

	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
