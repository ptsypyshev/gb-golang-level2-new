package conveyor

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)


type conveyor struct {
	inCh, sqCh, dbCh chan int
	wg               sync.WaitGroup
}

type input struct {
	value int
	exit bool
}

func New() *conveyor {
	return &conveyor{
		inCh: make(chan int),
		sqCh: make(chan int),
		dbCh: make(chan int),
	}
}

func (c *conveyor) Run(ctx context.Context) {
	c.wg.Add(4)
	go c.Read(ctx)
	go c.SquareNum()
	go c.DoubleNum()
	go c.Write()
	c.wg.Wait()
}

func (c *conveyor) Read(ctx context.Context) {
	defer c.wg.Done()
	defer close(c.inCh)

	var consoleInput string
	readCh := make(chan input)

	go func() {
		for {
			time.Sleep(time.Millisecond)
			fmt.Print("enter a number or 'stop' to exit app: ")
			
			_, err := fmt.Scan(&consoleInput)
			if err != nil {
				fmt.Printf("bad input: %s\n", err)
				continue
			}
			if consoleInput == "stop" {
				readCh <- input{exit: true}
				close(readCh)
				return
			}
			num, err := strconv.Atoi(consoleInput)
			if err != nil {
				fmt.Printf("bad input: %s\n", err)
				continue
			}

			readCh <- input{value: num}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case v := <- readCh:
			if v.exit {
				fmt.Println("Graceful shutdown by 'stop' command!")
				return
			}
			c.inCh <- v.value
		}
	}
}

func (c *conveyor) SquareNum() {
	defer c.wg.Done()
	defer close(c.sqCh)

	for v := range c.inCh {
		c.sqCh <- v * v
	}
}

func (c *conveyor) DoubleNum() {
	defer c.wg.Done()
	defer close(c.dbCh)

	for v := range c.sqCh {
		c.dbCh <- 2 * v
	}
}

func (c *conveyor) Write() {
	defer c.wg.Done()

	for v := range c.dbCh {
		fmt.Println("Result is:", v)
	}
}
