package main

import (
	"fmt"
)

func main() {
	// bool (true/false)
	fmt.Printf("Tupe: %T - Value: %v\n", true, true)

	// string - sequência de bytes
	fmt.Printf("Tupe: %T - Value: %v\n", "levi", "levi")
	fmt.Printf("Tupe: %T - Value: %v\n", "1", "1")

	// int
	fmt.Printf("Tupe: %T - Value: %v\n", 1, 1)

	// float (float64/float32) - decimal
	fmt.Printf("Tupe: %T - Value: %v\n", 1.233, 1.233)
}

// Tipos:
// bool (true/false)
// string - sequência de bytes
// int
// float (float64/float32) - decimal