package main

import "fmt"

func main() {
	hashTable := make(map[string]int)

	// create
	hashTable["Asus"] = 100
	hashTable["Mac"] = 200

	// get
	value := hashTable["Asus"]
	fmt.Println("Value for 'Asus':", value)

	value, exists := hashTable["Lenovo"]
	if exists {
		fmt.Println("Value for 'Lenovo':", value)
	} else {
		fmt.Println("'Lenovo' does not exists in the hash table")
	}

	// delete key
	delete(hashTable, "Mac")
}
