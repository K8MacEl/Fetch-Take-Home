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

func calculatePoints(receipt Receipt) int {
	//start points at 0
	points := 0
	 //One point for every alphanumeric character in the retailer name
	for _, char := range receipt.Retailer {
		if ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || ('0' <= char && char <= '9') {
			points++
		} 
	}
	// 50 pints if the total is a round dollar amount with no cents
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err == nil && total == float64(int(total)) {
		points += 50
	}
	//25 points if teh totla is a multiple of .25
	if total*100 == float64(int(total*100)) && int(total*100)%25 == 0 {
	points +=25
	}
}
