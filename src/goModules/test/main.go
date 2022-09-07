/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Fri May 20 2022
 */

/*
 * Run `go run <file-name>.go`
 * Reference: https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs
 */
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

	fmt.Println("Enter two number for sum")

	defer writer.Flush()

	var a, b int

	scanf("%d %d\n", &a, &b)

	printf("Sum: %d\n", a+b)
}
