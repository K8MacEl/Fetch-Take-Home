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
	ShortDescription string `json:"shortDescription"`
	Price string `json:"price"`
}