package main

import (
	"log"
	"os"

	"github.com/shubham1172/cgpg/codegen/generators"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("required: path to the directory containing .json files")
	}

	jsonMap, err := readJsonFiles(args[0])
	if err != nil {
		log.Fatal(err)
	}

	goGen := generators.NewGoGenerator()
	_, err = goGen.GenerateCode(jsonMap)
	if err != nil {
		log.Fatal(err)
	}
}
