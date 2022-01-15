package main

import (
	"fmt"
	"strings"
)

func main() {
	// // การประกาศ Map เปล่า
	// // map[string]string => map[key จะเป็น type อาราย]value จะเป็น type อาราย ** เป็น struc ก็ได้นะ
	emptyMap := map[string]string{}

	if len(emptyMap) == 0 {
		fmt.Println("emptyMap is empty")
	}
	emptyMap["a"] = "A"
	fmt.Println("emptyMap", emptyMap)

	// // ประกาศ Map แบบ Preassign

	m := map[string]string{
		"a": "A",
		"b": "B",
	}
	fmt.Println("m", m)

	input := "Apple Banana Apple Banana apple"
	mapWordCoubt := wordCount(input)
	fmt.Println("mapWordCoubt", mapWordCoubt)
}

func wordCount(s string) map[string]int {
	results := map[string]int{}
	words := strings.Fields(s)

	for i := 0; i < len(words); i++ {
		currentWord := words[i]
		// _, found := results[currentWord] // traditional style

		if _, found := results[currentWord]; !found { // go style
			results[currentWord] = 1
		} else {
			results[currentWord] += 1
		}
	}

	return results
}
