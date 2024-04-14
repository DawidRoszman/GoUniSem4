package main

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type FibFrequency struct {
	frequency map[int64]int64
	n         int64
}

func main() {
	nick := "DawRos"
	stripPolishLetters(&nick)
	makeLowercase(&nick)
	asciiNumbers := wordToAsciiNumbers(nick)
	fmt.Println(nick)
	fmt.Println(asciiNumbers)
	strongNumber := searchForFactorialContainingAllAscii(asciiNumbers)
	fmt.Println(strongNumber)
	fibFrequency := &FibFrequency{make(map[int64]int64), 30}
	fibonacci := calculateCallsInFibonacci(fibFrequency.n, fibFrequency)
	fmt.Println(fibonacci)
	fmt.Println(fibFrequency.frequency)
	times := make([]float64, 0)
	for i := 10; i <= 30; i++ {
		calculateFunctionTime(&times, "Fibonacci "+strconv.Itoa(i), calculateFib, int64(i))
	}
	var avg float64
	var sum float64
	for i := 0; i < len(times)-1; i++ {
		sum += times[i+1] / times[i]
	}
	avg = sum / float64(len(times))
	fmt.Printf("Avg: %f\n", avg)
	calculateFunctionTime(&times, "Fib 40", calculateFib, int64(40))
	prediction := times[20] * math.Pow(avg, float64(10))
	fmt.Printf("Predicted 40 %.2f\n", prediction)
	fmt.Printf("Accuracy: %f\n", prediction/times[21])
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

func calculateCallsInFibonacci(n int64, fibFrequency *FibFrequency) int64 {
	if n <= 1 {
		return n
	}
	fibFrequency.frequency[n-1] += 1
	fibFrequency.frequency[n-2] += 1
	return calculateCallsInFibonacci(n-1, fibFrequency) + calculateCallsInFibonacci(n-2, fibFrequency)
}

func timeTrack(start time.Time, name string, times *[]float64) {
	elapsed := time.Since(start).Nanoseconds()
	fmt.Printf("%s took %d\n", name, elapsed)
	*times = append(*times, float64(elapsed))
}

func calculateFunctionTime(times *[]float64, functionName string, f interface{}, args ...interface{}) {
	fValue := reflect.ValueOf(f)
	if fValue.Kind() != reflect.Func {
		fmt.Println("Error: f is not a function")
		return
	}

	numArgs := fValue.Type().NumIn()
	if len(args) != numArgs {
		fmt.Printf("Error: Expected %d arguments, but got %d\n", numArgs, len(args))
		return
	}

	argValues := make([]reflect.Value, len(args))
	for i, arg := range args {
		argValues[i] = reflect.ValueOf(arg)
	}

	defer timeTrack(time.Now(), functionName, times)
	fValue.Call(argValues)
}
