package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {

    // primero solo descomenten executionWithoutConcurrency y corranlo
	fmt.Println("Execution without concurrency")
    executionWithoutConcurrency()
    // luego solo descomenten executionWithConcurrency y sientan la concurrencia en su salsa!
	fmt.Println("Execution with concurrency and waitgroups")
    executionWithConcurrency()
	fmt.Println("Execution with concurrency and channels")
    executionWithConcurrencyUsingChannels()
}


func executionWithConcurrencyUsingChannels() {

    // TIP (más a la derecha... VER SOLO si no se te cae una idea.................................................. seguro que querés ver? tampoco es la gran cosa... .................................... bueno, necesitamos un canal para valores enteros  `var channel chan int = make(chan int)`)

	var ch chan int = make(chan int)

    go func() {
        sum := doAnExpensiveSum()
		ch <- sum
    }()
    go func() {
        number := getNumberFromStdin()
		ch <- number
    }()

	var totalSum, inputNum int
	for i := 0; i < 2; i++ {
		total := <- ch	
		if i==0 {
			totalSum = total
		} 
		inputNum = total
	}

    fmt.Println("Suma: ", totalSum)
    fmt.Println("Input: ", inputNum)
    // lo que quiero es imprimir aca tanto `sum` e `number`
}

func executionWithConcurrency() {
    var wg sync.WaitGroup

    wg.Add(2)

    go func() {
        sum := doAnExpensiveSum()
        fmt.Println(sum)
        wg.Done()
    }()
    go func() {
        input := getNumberFromStdin()
        fmt.Println(input)
        wg.Done()
    }()

    wg.Wait()
}

func executionWithoutConcurrency() {
    sum := doAnExpensiveSum()
    fmt.Println(sum)

    input := getNumberFromStdin()
    fmt.Println("Input: ", input)
}

func doAnExpensiveSum() int {
    ac := 0
	start := time.Now()
    for i := 1; i < 3000000000; i++ {
		ac += i
    }
	duration := time.Since(start)
	fmt.Println(duration)
    return ac
}

func getNumberFromStdin() int {
    scanner := bufio.NewScanner(os.Stdin)
    for true {
        fmt.Println("Enter valid number (or press CTR+C abort):")
        hasInput := scanner.Scan()
        if !hasInput {
            break
        }
        text := scanner.Text()
        number, err := strconv.Atoi(text)
        if err == nil {
            return number
        }
    }
    return 0
}
