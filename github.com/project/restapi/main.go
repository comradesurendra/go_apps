package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Banner struct {
	Id       int64  `json:"id"`
	Image    string `json:"image"`
	Redirect string `json:"redirect"`
}

type PersonalLoan struct {
	Id       int64   `json:"id"`
	Logo     string  `json:"logo"`
	Rate     float64 `json:"rate"`
	Tenure   string  `json:"tenure"`
	LoanAmt  int64   `json:"amt"`
	ViewMore string  `json:"viewmore"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home page.\n")
}

func getBanner(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: /banner")
	json.NewEncoder(w).Encode(Banners)
}

func getPersonalLoan(w http.ResponseWriter, r *http.Request) {
}

func getHomeLoan(w http.ResponseWriter, r *http.Request) {
}

var Banners []Banner

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/api/banner", getBanner)
	http.HandleFunc("/api/per_loan", getPersonalLoan)
	http.HandleFunc("/api/home_loan", getHomeLoan)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	Banners = []Banner{
		Banner{Id: 1, Image: "images1", Redirect: "http://google.com"},
		Banner{Id: 2, Image: "images2", Redirect: "http://facebook.com"},
		Banner{Id: 3, Image: "images3", Redirect: "http://apple.com.com"},
	}
	handleRequests()
}
