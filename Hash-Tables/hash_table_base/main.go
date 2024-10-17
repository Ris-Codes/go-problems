package main

import "fmt"

func main() {
	hashTable := Initialize(7)

	hashTable.Insert("asus", 5)
	hashTable.Insert("lenovo", 3)
	hashTable.Insert("mac", 7)

	if value, found := hashTable.Search("asus"); found {
		fmt.Println("Asus:", value)
	} else {
		fmt.Println("Asus not found")
	}

	hashTable.Delete("mac")
	if _, found := hashTable.Search("mac"); !found {
		fmt.Println("Mac deleted")
	}
}

type HashEntry struct {
	key   string
	value int
	next  *HashEntry // Handling collisions(chaining)
}

type HashTable struct {
	buckets []*HashEntry
	size    int
}

func Initialize(size int) *HashTable {
	return &HashTable{
		buckets: make([]*HashEntry, size),
		size:    size,
	}
}

// hash function
func (ht *HashTable) hash(key string) int {
	var sum int
	for _, v := range key {
		sum += int(v)
	}
	return sum % ht.size
}

// Insert
func (ht *HashTable) Insert(key string, value int) {
	idx := ht.hash(key)
	entry := &HashEntry{key: key, value: value}

	if ht.buckets[idx] == nil {
		ht.buckets[idx] = entry
	} else {
		// Handle collision using chaining
		current := ht.buckets[idx]
		for current.next != nil {
			// update value if key already exists
			if current.key == key {
				current.value = value
				return
			}
			current = current.next
		}
		// Add a new entry to the end of the chain
		current.next = entry
	}
}

// Search
func (ht *HashTable) Search(key string) (int, bool) {
	idx := ht.hash(key)
	current := ht.buckets[idx]

	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}
	return 0, false
}

// Delete
func (ht *HashTable) Delete(key string) {
	idx := ht.hash(key)
	current := ht.buckets[idx]
	var prev *HashEntry

	for current != nil {
		if current.key == key {
			if prev == nil {
				ht.buckets[idx] = current.next
			} else {
				prev.next = current.next
			}
			return
		}
		prev = current
		current = current.next
	}
}
