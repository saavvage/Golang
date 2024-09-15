package main
import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func productToJSON(p Product) string {
	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatalf("Error converting to JSON: %v", err)
	}
	return string(jsonData)
}

func jsonToProduct(jsonStr string) Product {
	var p Product
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	return p
}

func main() {
	product := Product{Name: "Laptop", Price: 500, Quantity: 10}

	jsonString := productToJSON(product)
	fmt.Println("JSON String:", jsonString)

	decodedProduct := jsonToProduct(jsonString)
	fmt.Printf("Decoded Product: %+v\n", decodedProduct)
}
