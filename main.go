package main

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/product"
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World! bahi ami go language sikhsi")
}

func main() {
	mux := http.NewServeMux()

	// mux.HandleFunc("/", helloWorldHandler)
	mux.Handle("GET /", http.HandlerFunc(helloWorldHandler))

	// mux.Handle("GET /products", http.HandlerFunc(getProductsHandler))

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProductsHandler))

	mux.Handle("OPTIONS /products", http.HandlerFunc(handlers.GetProductsHandler))
	mux.Handle("POST /createProduct", http.HandlerFunc(handlers.CreateProductHandler))
	mux.Handle("OPTIONS /createProduct", http.HandlerFunc(handlers.CreateProductHandler))
	fmt.Println("Server is running on http://localhost:8080")
	globalR := global_router.GlobalRouters(mux)
	err := http.ListenAndServe(":8080", globalR)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
func init() {
	pd1 := product.Product{
		ID:          1,
		Title:       "Product 1",
		Description: "Description 1",
		Price:       100.0,
		ImageUrl:    "http://example.com/product1.jpg",
	}
	pd2 := product.Product{
		ID:          2,
		Title:       "Product 2",
		Description: "Description 2",
		Price:       200.0,
		ImageUrl:    "http://example.com/product2.jpg",
	}
	pd3 := product.Product{
		ID:          3,
		Title:       "Product 3",
		Description: "Description 3",
		Price:       300.0,
		ImageUrl:    "http://example.com/product3.jpg",
	}
	pd4 := product.Product{
		ID:          4,
		Title:       "Product 4",
		Description: "Description 4",
		Price:       400.0,
		ImageUrl:    "http://example.com/product4.jpg",
	}
	pd5 := product.Product{
		ID:          5,
		Title:       "Product 5",
		Description: "Description 5",
		Price:       500.0,
		ImageUrl:    "http://example.com/product5.jpg",
	}
	product.ProductList = append(product.ProductList, pd1)
	product.ProductList = append(product.ProductList, pd2)
	product.ProductList = append(product.ProductList, pd3)
	product.ProductList = append(product.ProductList, pd4)
	product.ProductList = append(product.ProductList, pd5)
}
