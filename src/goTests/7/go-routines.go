/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Sun May 22 2022
 */

// REFERENCE: https://www.sololearn.com/learning/1164/4888/12671/1

package main

import (
	"fmt"
	"time"
)

func out(from int, to int, ch chan bool) {
	for i := from; i <= to; i++ {
		time.Sleep((50 * time.Millisecond))
		fmt.Println(i)
	}

	// send data to the channel
	ch <- true
}

func evenSum(from, to int, ch chan int) {
	result := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			result += i
		}
	}
	fmt.Println("evenSum = ", result)
	ch <- result
}

func squareSum(from, to int, ch chan int) {
	result := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			result += i * i
		}
	}
	fmt.Println("squareSum = ", result)
	ch <- result
}

// PROBLEM: calculate and output the sum of all even numbers in a given range plus the sum of their squares and output the result
func calculateSum() {
	evenCh := make(chan int)
	sqCh := make(chan int)

	go evenSum(0, 100, evenCh)
	go squareSum(0, 100, sqCh)

	// fmt.Println("output : ", <-evenCh + <-sqCh)

	fmt.Println("Starting infinite for loop")
	for {
		select {
		case x := <-evenCh:
			fmt.Println("evenCh = ", x)
			close(evenCh)
			return
		case y := <-sqCh:
			fmt.Println("sqCh = ", y)
			close(sqCh)
			return
		default:
			const sleepTime = 1
			fmt.Println("Nothing available after ", sleepTime, " milliSeconds")
			time.Sleep((sleepTime * time.Millisecond))
		}
	}
}

func download(val int, ch chan int) {
	sum := (val * (val + 1)) / 2
	ch <- sum
}

func main() {

	// channel

	// The type after the chan keyword indicates the type of the data we will send through the channel.
	ch := make(chan bool)

	go out(0, 5, ch)
	go out(6, 10, ch)

	// adding sleep to finish executing both go function before main gets exit
	// not needed when using channel
	// time.Sleep((500 * time.Millisecond))

	// receive data from the channel in a variable
	// value := <-ch

	// If we do not need the value as a variable
	<-ch

	/*
	 * NOTE:
	 * The receive operation blocks the code until and unless some data is sent by the send operation.
	 * If no data is received, a deadlock will occur, blocking the code from executing.
	 */

	// If you do not need to send data to a channel anymore, you can close it using close(ch), where ch is the name of the channel. This is done in the sender.
	close(ch)

	calculateSum()
}
