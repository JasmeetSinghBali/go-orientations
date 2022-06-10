package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"key1": 5,
		"key2": 10,
		"key3": 1,
	}
	fmt.Println(mp)
	mp["key1"] = 55
	mp["key4"] = 69
	fmt.Println("adding and updating maps\n", mp)

	delete(mp, "key2")
	fmt.Println("after deleting key 2\n", mp)

	val, ok := mp["tim"]
	fmt.Println("trying to access tim key", val, ok)
}
