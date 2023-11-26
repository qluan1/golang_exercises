package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	fs := flag.NewFlagSet("Prints out SHA256/SHA384/SHA512 hash of stdin string.", flag.ExitOnError)
	fs.SetOutput(os.Stdout)

	useSHA384 := fs.Bool("sha384", false, "use SHA384 hash function")
	useSHA512 := fs.Bool("sha512", false, "use SHA512 hash function")
	fs.Usage = func() {
		fmt.Println("Usage: [argument] [flag]")
		fmt.Println("Available flags:")
		fs.PrintDefaults()
	}

	fs.Parse(os.Args[1:])

	if *useSHA384 && *useSHA512 {
		fmt.Fprintln(os.Stderr, "pick no more than one flag")
		return
	}

	args := fs.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "no input received")
		return
	}
	if len(args) > 1 {
		fmt.Fprintln(os.Stderr, "more than one input received")
		return
	}

	if *useSHA384 {
		fmt.Printf("%x\n", sha512.Sum384([]byte(args[0])))
		return
	}

	if *useSHA512 {
		fmt.Printf("%x\n", sha512.Sum512([]byte(args[0])))
		return
	}

	fmt.Printf("%x\n", sha256.Sum256([]byte(args[0])))
}