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
	// Format nickname
	name := "Dawid"
	surname := "Roszman"
	nick := name[:3] + surname[:3]
	stripPolishLetters(&nick)
	makeLowercase(&nick)
	asciiNumbers := wordToAsciiNumbers(nick)
	fmt.Println(nick)
	fmt.Println(asciiNumbers)

	// Calculate Strong number
	strongNumber := searchForFactorialContainingAllAscii(asciiNumbers)
	fmt.Println(strongNumber)

	// Calculate Fibonacci Frequency
	fibFrequency := &FibFrequency{make(map[int64]int64), 30}
	fibonacci := calculateCallsInFibonacci(fibFrequency.n, fibFrequency)
	fmt.Println(fibonacci)
	fmt.Println(fibFrequency.frequency)

	times := make([]float64, 0)
	start, end := 10, 30
	fibAvg := calculateExponentailRateFib(&times, start, end)
	makePredictions(40, fibAvg, int64(end), &times)
	makePredictions(100, fibAvg, int64(end), &times)
	makePredictions(strongNumber, fibAvg, int64(end), &times)
	result := ackermann(3, 4)
	times = make([]float64, 0)
	ackAvg := calculateExponentailRateAck(&times, 1, 3)
	fmt.Println("Ackermann(3, 4) =", result)
	fmt.Println(ackAvg)
	makePredictions(10, ackAvg, 3, &times)
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

func formatDuration(seconds float64) string {
	years := math.Floor(seconds / (365 * 24 * 60 * 60))
	seconds -= years * 365 * 24 * 60 * 60
	months := math.Floor(seconds / (30 * 24 * 60 * 60))
	seconds -= months * 30 * 24 * 60 * 60
	days := math.Floor(seconds / (24 * 60 * 60))
	seconds -= days * 24 * 60 * 60
	hours := math.Floor(seconds / (60 * 60))
	seconds -= hours * 60 * 60
	minutes := math.Floor(seconds / 60)
	seconds -= minutes * 60

	formatted := ""
	if years > 0 {
		formatted += fmt.Sprintf("%.0f years ", years)
	}
	if months > 0 {
		formatted += fmt.Sprintf("%.0f months ", months)
	}
	if days > 0 {
		formatted += fmt.Sprintf("%.0f days ", days)
	}
	if hours > 0 {
		formatted += fmt.Sprintf("%.0f hours ", hours)
	}
	if minutes > 0 {
		formatted += fmt.Sprintf("%.0f minutes ", minutes)
	}
	if seconds > 0 {
		formatted += fmt.Sprintf("%.0f seconds", seconds)
	}

	return formatted
}

func calculateExponentailRateFib(times *[]float64, start, end int) float64 {
	for i := start; i <= end; i++ {
		calculateFunctionTime(times, "Fibonacci "+strconv.Itoa(i), calculateFib, int64(i))
	}
	var sum float64
	for i := 0; i < len(*times)-1; i++ {
		sum += (*times)[i+1] / (*times)[i]
	}
	return math.Ceil((sum/float64(len(*times)))*10) / 10
}

func makePredictions(predicting int64, avg float64, end int64, times *[]float64) {
	fmt.Printf("Avg: %.1f\n", avg)
	prediction := (*times)[len(*times)-1] * math.Pow(avg, float64(predicting-end))
	fmt.Printf("Predicted %d %s\n", predicting, formatDuration(prediction/1_000_000_000))
}

func ackermann(m, n int) int {
	if m == 0 {
		return n + 1
	} else if m > 0 && n == 0 {
		return ackermann(m-1, 1)
	} else if m > 0 && n > 0 {
		return ackermann(m-1, ackermann(m, n-1))
	}
	return -1 // Invalid input
}

func calculateExponentailRateAck(times *[]float64, start, end int) float64 {
	for i := start; i <= end; i++ {
		calculateFunctionTime(times, "Ackermann "+strconv.Itoa(i)+" "+strconv.Itoa(i+1), ackermann, i, i+1)
	}
	var sum float64
	for i := 0; i < len(*times)-1; i++ {
		sum += (*times)[i+1] / (*times)[i]
	}
	return math.Ceil((sum/float64(len(*times)))*10) / 10
}
