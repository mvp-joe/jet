package main

import (
	"flag"
	"fmt"
	"gopkg.in/errgo.v2/errors"
	"jet"
	"log"
	"os"
)

var fileFlag = flag.String("f", "", "path to script file")

func main() {
	// parse command-line flags
	flag.Parse()
	if *fileFlag == "" {
		log.Fatal("no script file specified (use -f flag)")
	}

	// read script file
	script, err := os.ReadFile(*fileFlag)
	if err != nil {
		fmt.Println("error reading script file")
		fmt.Println(errors.Details(err))
		log.Fatal(err)
	}

	// read stdin
	input := jet.ReadStdin()

	// run script
	result, err := jet.RunJs(string(script), input)
	if err != nil {
		fmt.Println("error running script")
		fmt.Println(errors.Details(err))
		log.Fatal(err)
	}

	// print result JSON
	fmt.Println(string(result))
}
