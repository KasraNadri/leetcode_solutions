package main

import "fmt"

func main() {
	fmt.Println(isValid("([])"))
}

func isValid(s string) bool {
	symbols := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	slot := string(s[0])

	for index1, value1 := range s[1:] {
		for index2, value2 := range slot {
			if value
		}
	}

}
