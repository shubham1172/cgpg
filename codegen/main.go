package main

import (
	"log"
	"os"
	"path"

	"github.com/shubham1172/cgpg/codegen/generators"
)

const generatedDir = "../generated"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("required: path to the directory containing .json files")
	}

	jsonMap, err := readJsonFiles(args[0])
	if err != nil {
		log.Fatal(err)
	}

	generators := []generators.Generator{
		generators.NewPythonGenerator(),
		generators.NewGoGenerator(),
	}

	for _, generator := range generators {
		fileMap, err := generator.GenerateCode(jsonMap)
		if err != nil {
			log.Fatal(err)
		}

		for fileName, fileContents := range fileMap {
			filePath := path.Join(generatedDir, generator.GetName(), fileName)

			err = os.MkdirAll(path.Dir(filePath), 0755)
			if err != nil {
				log.Fatal(err)
			}

			err = os.WriteFile(filePath, []byte(fileContents), 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
