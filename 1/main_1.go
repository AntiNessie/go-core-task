package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	a := []any{numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum}

	for _, v := range a {
		fmt.Printf("Тип: %T\n", v)
	}

	b := toString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println(b)

	fmt.Println(toRune(b))

	runes := toRune(b)
	hashed := hashWithSalt(runes)
	fmt.Println("SHA256", hashed)

}

func toString(vars ...any) string {
	result := ""

	for _, v := range vars {
		result += fmt.Sprint(v)
	}
	return result

}

func toRune(b string) []rune {
	return []rune(b)

}

func hashWithSalt(runes []rune) string {

	mid := len(runes) / 2
	runes = append(runes[:mid], append([]rune("go-2024"), runes[mid:]...)...)
	str := string(runes)
	hash := sha256.New()
	hash.Write([]byte(str))
	hashed := hash.Sum(nil)

	return fmt.Sprintf("%x", hashed)

}
