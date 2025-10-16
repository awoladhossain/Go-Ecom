package handlers

import (
	"ecommerce/product"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct product.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	newProduct.ID = len(product.ProductList) + 1
	product.ProductList = append(product.ProductList, newProduct)

	// w.WriteHeader(http.StatusCreated) // 201 Created naile cors error dibe
	// encoder := json.NewEncoder(w)
	// encoder.Encode(newProduct)
	utils.SendData(w, newProduct, http.StatusCreated)
}
