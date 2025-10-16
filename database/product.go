package database


var ProductList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"image_url"`
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
	pd4 :=Product{
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
	ProductList = append(ProductList, pd1)
	ProductList = append(ProductList, pd2)
	ProductList = append(ProductList, pd3)
	ProductList = append(ProductList, pd4)
ProductList = append(ProductList, pd5)
}
