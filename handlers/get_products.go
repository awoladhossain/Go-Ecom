package handlers

import (
	"ecommerce/product"
	"ecommerce/utils"
	"net/http"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "GET" {
	// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }
	// encoder := json.NewEncoder(w)
	// encoder.Encode(productList)
	utils.SendData(w, product.ProductList, http.StatusOK)
}
