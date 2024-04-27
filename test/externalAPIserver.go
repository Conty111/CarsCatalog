package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-faker/faker/v3"
	"log"
	"math/rand"
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

var (
	marks = []string{"Lada", "BMW", "Mercedes", "Volga", "Toyota"}
)

func getInfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL, r.Method, r.UserAgent())

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
	return Car{
		RegNum: regNum,
		Mark:   marks[rand.Intn(len(marks))],
		Model:  "some model",
		Year:   rand.Intn(30) + 1990, // случайный год от 1990 до 2019
		Owner: People{
			Name:       faker.Name(),
			Surname:    faker.LastName(),
			Patronymic: faker.FirstName(),
		},
	}
}

func main() {
	http.HandleFunc("/info", getInfoHandler)

	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
