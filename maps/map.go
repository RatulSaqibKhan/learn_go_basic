package main

import ("fmt")

func main() {
	_stock := map[string]float64 {
		"amazon" : 223.4,
		"facebook" : 1234.4,
	}

	// Number of items in a map
	fmt.Println(len(_stock))

	// Get zero value if not found
	fmt.Println(_stock["tesla"]) // 0

	// Use two value to see if found
	value, ok := _stock["tesla"]
	
	if !ok {
		fmt.Println("tesla not found")
	} else {
		fmt.Printf("tesla found %f \n", value)
	}

	// Set
	_stock["tesla"] = 234.2

	fmt.Println(_stock)

	// delete a key
	delete(_stock, "amazon")

	fmt.Println(_stock)
	
	// Single value "for" on keys
	for key := range _stock {
		fmt.Println(key)
	}

	// Double value "for" on key, value pairs
	for key, val := range _stock {
		fmt.Printf("%v => %f\n", key, val)
	}
}