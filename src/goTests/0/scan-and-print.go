/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Fri May 20 2022
 */

/*
 * Run `go run <file-name>.go`
 * REFERENCE: https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs
 */

// NOTE: Use golang for CP: https://www.codementor.io/@tucnak/using-golang-for-competitive-programming-h8lhvxzt3

package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {

	defer writer.Flush()

	fmt.Println("Enter two number for sum for testing: ")

	var a, b int

	scanf("%d %d\n", &a, &b)

	printf("Sum: %d\n", a+b)
}
