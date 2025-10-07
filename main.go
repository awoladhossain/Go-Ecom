package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World! bahi ami go language sikhsi")
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"image_url"`
}

var productList []Product

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	w.WriteHeader(http.StatusCreated) // 201 Created naile cors error dibe
	encoder := json.NewEncoder(w)
	encoder.Encode(newProduct)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", helloWorldHandler)
	mux.HandleFunc("/products", getProductsHandler)
	mux.HandleFunc("/createProduct", createProductHandler)
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
func init() {
	pd1 := Product{
		ID:          1,
		Title:       "Product 1",
		Description: "Description 1",
		Price:       100.0,
		ImageUrl:    "http://example.com/product1.jpg",
	}
	pd2 := Product{
		ID:          2,
		Title:       "Product 2",
		Description: "Description 2",
		Price:       200.0,
		ImageUrl:    "http://example.com/product2.jpg",
	}
	pd3 := Product{
		ID:          3,
		Title:       "Product 3",
		Description: "Description 3",
		Price:       300.0,
		ImageUrl:    "http://example.com/product3.jpg",
	}
	pd4 := Product{
		ID:          4,
		Title:       "Product 4",
		Description: "Description 4",
		Price:       400.0,
		ImageUrl:    "http://example.com/product4.jpg",
	}
	pd5 := Product{
		ID:          5,
		Title:       "Product 5",
		Description: "Description 5",
		Price:       500.0,
		ImageUrl:    "http://example.com/product5.jpg",
	}
	productList = append(productList, pd1)
	productList = append(productList, pd2)
	productList = append(productList, pd3)
	productList = append(productList, pd4)
	productList = append(productList, pd5)
}
