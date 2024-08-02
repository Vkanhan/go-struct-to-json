package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// User struct holds info about the user
type User struct {
	Name  string
	Email string
	Age   int
}

// Discovery struct holds the info about the discovery of the user
type Discovery struct {
	Name  string
	Year  int
	Prize string
}

// PrintAsJSON prints any struct as a formatted JSON string
func PrintAsJSON(data interface{}) {
	//Return reflect.value which represents the runtime value of data
	v := reflect.ValueOf(data)

	// Check if we have a struct
	if v.Kind() != reflect.Struct {
		fmt.Println("PrintAsJSON only accepts structs")
		return
	}

	// Convert struct to JSON
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Error converting to JSON: %v\n", err)
		return
	}

	// Print the JSON string
	fmt.Println(string(jsonBytes))
}

func main() {
	// Create a User struct instance
	user := User{Name: "Einstein", Email: "einstein@gmail.com", Age: 43}

	//Create a Product struct instance
	product := Discovery{Name: "Photoelectric effect", Year: 1905, Prize: "Nobel Prize"}

	//Print User struct as JSON
	fmt.Println("User struct as JSON:")
	PrintAsJSON(user)

	//Print Discovery struct as JSON
	fmt.Println("\nDiscovery as JSON:")
	PrintAsJSON(product)
}
