// package main

// import "fmt"

// func main() {
// 	weather := [3]string{"summer", "winter", "rainy"}
// 	for i := 0; i < len(weather); i++ {
// 		fmt.Println(weather[i])
// 	}
// }

package main

import "fmt"

func main() {
	// Arrays
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	// Slices
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// Maps
	myMap := make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	fmt.Println(myMap)
}
