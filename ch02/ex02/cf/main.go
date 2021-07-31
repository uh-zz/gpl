package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("No command line arguments, So enter a number or Please exit Program.")

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			t, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			ArgsToTempConv(t)
			ArgsToLenConv(t)
		}
	}
	for _, arg := range flag.Args() {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		ArgsToTempConv(t)
		ArgsToLenConv(t)
	}
}
