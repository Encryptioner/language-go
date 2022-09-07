/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Fri May 20 2022
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	defer writer.Flush()

	var s string

	fmt.Println("Enter a string:")

	scanf("%s\n", &s)

	s = regexp.MustCompile("(WUB)+").ReplaceAllLiteralString(s, " ")
	s = strings.Trim(s, " ")
	printf("%s\n", s)
}
