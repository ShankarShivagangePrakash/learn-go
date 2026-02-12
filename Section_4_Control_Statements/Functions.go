package main

import "fmt"

type Product struct {
	id int
	name string
	price float64
}

func search(products []Product, productId int){
	for _, product := range products {
		if product.id == productId {
			fmt.Println("Product Found:", product)
			return
		}
	}
	fmt.Println("Product Not Found")
}

func main(){

	products := []Product{
		{1, "Laptop", 999.99},
		{2, "Smartphone", 499.99},
	}

	fmt.Print("Enter product ID to search:")
	var productId int
	fmt.Scanln(&productId)

	search(products[:], productId)

}
