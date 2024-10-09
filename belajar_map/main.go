package main

import "fmt"

func main() {
	m := make(map[string]string) // create empty map
	m["warna"] = "merah"         // set value to map, key & value
	fmt.Println("warna:", m["warna"])

	m["bentuk"] = "kotak"
	fmt.Println("bentuk:", m["bentuk"])

	fmt.Println("debug", m) // print all
	fmt.Println("==========")

	// literal init map
	mapBaru := map[string]float64{
		"pi":           3.14,
		"euler_number": 2.7182,
		"satu":         9,
		"hapus_ini":    -1,

		// ascii char value
		"A": 65, "B": 66, "C": 67, "D": 68, "E": 69, // not recommended write code like this
	}

	pi := mapBaru["pi"]          // get value from map
	mapBaru["satu"] = 1          // set value to map
	delete(mapBaru, "hapus_ini") // delete value from map

	fmt.Println("debug pi =", pi)
	fmt.Println("tes gaada key:", mapBaru["gaada key-nya"]) // output zero value
	fmt.Println("==========")

	satu, ok := mapBaru["satu"] // get 2 value from map; variable ok => type bool
	if ok {
		fmt.Println("satu =", satu)
	}

	if gaada, ok := mapBaru["gaada"]; ok { // shortcut one-line with if block
		// variable gaada punya scope local disini
		fmt.Println("gaada:", gaada)
	} else {
		fmt.Println("gaada key nya ya gaada :)")
	}
	fmt.Println("==========")

	// Range over Map
	// beware: it's random order, try run program few times to prove
	for key, value := range mapBaru {
		fmt.Println(key, " = ", value) //print each keys & values in map
	}
}
