package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	nick := "DawRos"
	stripPolishLetters(&nick)
	makeLowercase(&nick)
	asciiNumbers := wordToAsciiNumbers(nick)
	fmt.Println(nick)
	fmt.Println(asciiNumbers)
	strongNumber := searchForFactorialContainingAllAscii(asciiNumbers)
	fmt.Println(strongNumber)
	weakNumberCache := calculateFibWithCache(strongNumber, make(map[int64]*big.Int))
	fmt.Println(weakNumberCache)
	weakNumber := calculateFib(strongNumber)
	fmt.Println(weakNumber)
}

func stripPolishLetters(nick *string) {
	polishLetters := map[string]string{
		"ą": "a",
		"ć": "c",
		"ę": "e",
		"ł": "l",
		"ń": "n",
		"ó": "o",
		"ś": "s",
		"ź": "z",
		"ż": "z",
	}
	for key, value := range polishLetters {
		*nick = strings.ReplaceAll(*nick, key, value)
	}
}

func makeLowercase(nick *string) {
	*nick = strings.ToLower(*nick)
}

func letterToAsciiNumber(letter string) int {
	return int(letter[0])
}

func wordToAsciiNumbers(word string) []int {
	asciiNumbers := make([]int, len(word))
	for i, letter := range word {
		asciiNumbers[i] = letterToAsciiNumber(string(letter))
	}
	return asciiNumbers
}

func calculateFactorial(n int64) *big.Int {
	result, _ := new(big.Int).SetString("1", 10)
	for i := int64(1); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result
}

func checkIfAllAsciiAreInFactorial(asciiNumbers []int, fac *big.Int) bool {
	for _, ascii := range asciiNumbers {
		if !strings.Contains(fac.String(), strconv.Itoa(ascii)) {
			return false
		}
	}
	return true
}

func searchForFactorialContainingAllAscii(asciiNumbers []int) int64 {
	running := true
	var index int64 = 0
	for running {
		index++
		factorial := calculateFactorial(index)

		running = !checkIfAllAsciiAreInFactorial(asciiNumbers, factorial)
	}
	return index
}

func calculateFib(n int64) *big.Int {
	if n <= 1 {
		return big.NewInt(n)
	}
	return new(big.Int).Add(calculateFib(n-1), calculateFib(n-2))
}

func calculateFibWithCache(n int64, cache map[int64]*big.Int) *big.Int {
	if n <= 1 {
		return big.NewInt(n)
	}
	if _, ok := cache[n]; !ok {
		cache[n] = new(big.Int).Add(calculateFibWithCache(n-1, cache), calculateFibWithCache(n-2, cache))
	}
	return cache[n]
}
