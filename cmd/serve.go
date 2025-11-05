package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(middleware.Logger)
	mux := http.NewServeMux()

	// mux.Handle("GET /abc", manager.With(middleware.Logger)(http.HandlerFunc(handlers.GetProductsHandler)))

	// mux.Handle("GET /abc", manager.With(http.HandlerFunc(handlers.GetProductsHandler), middleware.Logger))
	mux.Handle("GET /abc", manager.With(http.HandlerFunc(handlers.GetProductsHandler)))
	// mux.Handle("GET /products", http.HandlerFunc(getProductsHandler))
	// loggerMiddle = middleware.Logger()
	// mux.Handle("GET /products", middleware.Logger(http.HandlerFunc(handlers.GetProductsHandler)))
	mux.Handle("GET /produts", manager.With(http.HandlerFunc(handlers.GetProductsHandler)))
	// mux.Handle("POST /createProduct", http.HandlerFunc(handlers.CreateProductHandler))
	mux.Handle("POST /createProduct", manager.With(http.HandlerFunc(handlers.CreateProductHandler)))
	// mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductsByID))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProductsByID)))

	fmt.Println("Server is running on http://localhost:8080")
	globalR := global_router.GlobalRouters(mux)

	err := http.ListenAndServe(":8080", globalR)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// RESTApi => representational state transfer means api is stateless and data is stateful
//  from backend we wants resource
