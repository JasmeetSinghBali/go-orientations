package main

import "fmt"

func main() {
	/*a map m map[] with keys as string map[string] & values as string*/
	var m map[string]string
	// default value of map is nil
	fmt.Println(m == nil)
	/*📝way -1 to instantiate a map*/
	m = map[string]string{} // instantiate map object i.e any empty map
	// no longer nil
	fmt.Println(m == nil)
	fmt.Println(m)
	fmt.Println(len(m))
	/*📝way -2 to instantiate a map*/
	m = make(map[string]string, 5) // with preallocate space as 5
	fmt.Println(len(m))
	/*📝way -3 to instantiate a map with predefined elements*/
	m = map[string]string{"john": "Manager"}
	fmt.Println(m)
	fmt.Println(len(m))
	/*📝Add new key:values to map*/
	m["lucy"] = "Not a Manager"
	fmt.Println(m)
	/*📝Modify existing values for key in map*/
	m["lucy"] = "Manager"
	fmt.Println(m)
	/*📝access key in map*/
	fmt.Println(m["john"])

	delete(m, "lucy")
	fmt.Println(m)
	m["jack"] = "Programmer"
	m["jack"] += " Admin"
	fmt.Println(m)

	for name, title := range m {
		fmt.Println(name, title)
	}

	/*📝check wheather the key exist in map*/
	title, ok := m["jack"]
	if ok {
		fmt.Println(title)
	} else {
		fmt.Println("didnt found jack")
	}
}
