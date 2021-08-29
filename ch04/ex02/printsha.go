package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
	"regexp"
)

const (
	SHA256 = 256
	SHA384 = 384
	SHA512 = 512
)

func main() {
	var sha = flag.Int("sha", SHA256, "encription flag")
	flag.Parse()

	if !isSHA(*sha) {
		fmt.Printf("sorry, not support SHA%d\n", *sha)
		os.Exit(1)
	}

	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Please Enter any key or keyword:")
		sc.Scan()
		input := sc.Text()

		if isGoodBye(input) {
			fmt.Println("Ok, ByeBye")
			os.Exit(1)
		}

		fmt.Print("your input: ", input)
		switch *sha {
		case SHA256:
			fmt.Printf("your input(SHA256): %x\n", sha256.Sum256([]byte(input)))
		case SHA384:
			fmt.Printf("your input(SHA384): %x\n", sha512.Sum384([]byte(input)))
		case SHA512:
			fmt.Printf("your input(SHA512): %x\n", sha512.Sum512([]byte(input)))
		}
	}
}

func isSHA(sha int) bool {
	return sha == SHA256 || sha == SHA384 || sha == SHA512
}

func isGoodBye(input string) bool {
	r := regexp.MustCompile(`realy good bye`)
	return r.MatchString(input)
}
