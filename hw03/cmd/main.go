package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	inCh, sqCh, dbCh := make(chan int), make(chan int), make(chan int)

	wg.Add(4)
	go func() {
		defer wg.Done()
		defer close(inCh)

		var input string

		for {
			fmt.Print("enter a number or 'stop' to exit app: ")
			_, err := fmt.Scan(&input)
			if err != nil {
				fmt.Printf("bad input: %s\n", err)
				continue
			}
			if input == "stop" {
				return
			}
			num, err := strconv.Atoi(input)
			if err != nil {
				fmt.Printf("bad input: %s\n", err)
				continue
			}

			inCh <- num
			time.Sleep(time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(sqCh)

		for v := range inCh {
			sqCh <- v * v
		}
	}()

	go func() {
		defer wg.Done()
		defer close(dbCh)

		for v := range sqCh {
			dbCh <- 2 * v
		}
	}()

	go func() {
		defer wg.Done()

		for v := range dbCh {
			fmt.Println("Result is:", v)
		}
	}()

	wg.Wait()
}
