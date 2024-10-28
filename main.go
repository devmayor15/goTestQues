package main

import (
	"fmt"
	"math"
	"strings"
	"errors"
)

// SOLVED QUESTIONS IN FUNC
func sum(a int, b int) int {
	return a + b
}

// return true if even or otherwise
func isEven(n int) bool {
	return n%2 == 0
}

// calculate factorial
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// convert temp
func toFahrenheit(celcius float64) float64 {
	return (celcius * 9 / 5) + 32
}
func toCelcius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

// find maximum value
func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// calculate fibonacci sequence
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// check if a number is prime
func isPrime(n int) bool {
	if n < 1 {
		return false
	}
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// count vowels in a string
func countVowels(s string) int {
	s = strings.ToLower(s)
	count := 0
	for _, char := range s {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			count++

		}
	}
	return count
}

//function returning another function
func adder(x int) func(int) int{
	return func(y int) int {
		return x + y
	} 
}

//basic calculator
func calculator(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("cannot divide by zero")
		}
		return a / b, nil
	default:
		return 0, errors.New("unsupported operation")
	}
}


//SLICES
//sum  of elements
func sumSlice(nums []int) int{
	sum := 0
	for _, num := range nums{
		sum += num
	}
	return sum
}

//finding maximum value with map
func highestScore(scores map[string]int) (string, int, error) {
	if len(scores) == 0 {
		return "", 0, errors.New("no scores available")
	}


    maxScore := -1
    topStudent := ""
    for student, score := range scores {
        if score > maxScore {
            maxScore = score
            topStudent = student
        }
    }
    return topStudent, maxScore, nil
}

//high order function where func is another param
func operate(a, b int, op func(int, int) int) int{
	return op(a, b)
}
func add(x, y int) int {
    return x + y
}

//closures
func makeCounter() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}
func main() {
	

	fmt.Println(sum(1, 2))

	fmt.Println(isEven(2))
	fmt.Println(isEven(7))

	fmt.Println(factorial(5))
	fmt.Println(factorial(3))

	fmt.Println(toFahrenheit(32))
	fmt.Println(toCelcius(32))

	fmt.Println(findMax(10, 20))

	fmt.Println(fibonacci(10))

	fmt.Println(isPrime(3))

	fmt.Println(countVowels("haseiouuuuu"))

	fmt.Println(adder(5)(3))

	// Test cases for the calculator function
    result, err := calculator(10, 5, "+")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("10 + 5 =", result) // Output: 10 + 5 = 15
    }

    result, err = calculator(10, 5, "-")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("10 - 5 =", result) // Output: 10 - 5 = 5
    }

    result, err = calculator(10, 5, "*")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("10 * 5 =", result) // Output: 10 * 5 = 50
    }

    result, err = calculator(10, 0, "/")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("10 / 5 =", result) // Output: 10 / 5 = 2
    }

	nums := []int{1, 2, 3, 4, 5}
    resultSlice := sumSlice(nums)
    fmt.Println("Sum of slice elements:", resultSlice)

	//highest value map
	scores := map[string]int{
        "Alice": 85,
        "Bob":   92,
        "Cathy": 78,
        "David": 90,
    }
	student, score, err := highestScore(scores)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Top student:", student, score)
	}


	//high order
	resulting := operate(2, 3, add)
	fmt.Println(resulting)

	//closures
	count := makeCounter()
	fmt.Println(count())
	fmt.Println(count())
}
