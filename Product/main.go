package main

import (
	"fmt"
)

// Product struct with attributes Name, Price, and Stock
type Product struct {
	Name  string
	Price float64
	Stock int
}

// Function to update the product's price
func (p *Product) updatePrice(newPrice float64) {
	p.Price = newPrice
}

// Function to reduce the product's stock when sold
func (p *Product) updateStock(sold int) {
	if sold > p.Stock {
		fmt.Println("Insufficient stock")
	} else {
		p.Stock -= sold
	}
}

func main() {
	//1. Creating example products
	product1 := Product{Name: "Laptop", Price: 15000.0, Stock: 10}
	product2 := Product{Name: "Headphone", Price: 500.0, Stock: 25}

	// Displaying initial product data
	fmt.Printf("Product 1: %+v\n", product1)
	fmt.Printf("Product 2: %+v\n", product2)

	//2. Updating product prices
	product1.updatePrice(14000.0)
	product2.updatePrice(450.0)

	//3. Displaying updated product prices
	fmt.Printf("Price of %s: %f\n", product1.Name, product1.Price)
	fmt.Printf("Price of %s: %f\n", product2.Name, product2.Price)

	//4. Updating product stocks
	product1.updateStock(3)
	product2.updateStock(5)

	//5. Displaying updated product stocks
	fmt.Printf("Stock of %s: %d\n", product1.Name, product1.Stock)
	fmt.Printf("Stock of %s: %d\n", product2.Name, product2.Stock)
}
