package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"log"
	"net/http"
	"strconv"
)

func GetProductsByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	id, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	for _, product := range database.ProductList {
		// log.Println(_, product)
		log.Println("product-id", id)
		if product.ID == id {
			utils.SendData(w, product, http.StatusOK)
			return
		}
	}
	utils.SendData(w, "Product not found", http.StatusNotFound)
}
