package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Item struct {
	ShortDescription    string `json:"shortDescription"`
	Price 				string `json:"price"`
}

type Receipt struct {
	Retailer 			string `json:"retailer"`
	PurchaseDate		string `json:"purchaseDate"`
	PurchaseTime		string `json:"purchaseTime"`
	Items				[]Item `json:"items"`
	Total				string	`json:"total"`
}

type ProcessResponse struct {
	ID string `json:"id"`
}

type PointsResponse struct {
	Points int `json:"points"`
}

var receiptStore = make(map[string]Receipt)
var pointsStore = make(map[string]int)

func processReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	receiptStore[id] = receipt
	pointsStore[id] = calculatePoints(receipt)

	response := ProcessResponse{ID: id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func calculatePoints(

	// write function here
)