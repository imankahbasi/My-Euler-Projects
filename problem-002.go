/*
Even Fibonacci numbers
Problem 2
Each new term in the Fibonacci sequence is generated by adding the previous two terms.
By starting with 1 and 2, the first 10 terms will be:
1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
By considering the terms in the Fibonacci sequence whose values do not exceed four million,
find the sum of the even-valued terms.
*/

package main

import "fmt"

func main() {
	var fir, sec, sum int
	fir = 1
	sec = 2
	sum = sec
	for i := 0; sec <= 4000000; i++ {
		fir, sec = sec, fir+sec
		if sec%2 == 0 {
			sum += sec
		}
	}
	fmt.Println(sum)
}