package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"` // Adding json tags for better JSON representation
	Image       string  `json:"image"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

var ProductsList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != http.MethodGet {
		http.Error(w, "Get method required", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(ProductsList)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")             // Setting content type to JSON
	w.Header().Set("Access-Control-Allow-Origin", "*")             // Allowing CORS for all origins
	w.Header().Set("Access-Control-Allow-Methods", "POST")         // Allowing only POST method
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type") // Allowing Content-Type header

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Post method required", 400)
		return
	}

	deocder := json.NewDecoder(r.Body)
	var newProduct Product
	err := deocder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request payload", 400)
		return
	}

	newProduct.ID = len(ProductsList) + 1

	w.WriteHeader(201)
	ProductsList = append(ProductsList, newProduct)
	json.NewEncoder(w).Encode(newProduct)
}

func init() {
	fmt.Println("Server is starting...")
	p1 := Product{
		ID:          1,
		Image:       "image1.jpg",
		Name:        "Product 1",
		Description: "Description for product 1",
		Price:       19.99,
	}
	p2 := Product{
		ID:          2,
		Image:       "image2.jpg",
		Name:        "Product 2",
		Description: "Description for product 2",
		Price:       29.99,
	}
	p3 := Product{
		ID:          3,
		Image:       "image3.jpg",
		Name:        "Product 3",
		Description: "Description for product 3",
		Price:       39.99,
	}
	p4 := Product{
		ID:          4,
		Image:       "image4.jpg",
		Name:        "Product 4",
		Description: "Description for product 4",
		Price:       49.99,
	}
	ProductsList = []Product{p1, p2, p3, p4}
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/getProducts", getProducts)
	mux.HandleFunc("/createProduct", createProduct)

	err := http.ListenAndServe(":80", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
