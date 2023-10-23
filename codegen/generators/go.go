package generators

import "fmt"

type goGenerator struct{}

func NewGoGenerator() Generator {
	return &goGenerator{}
}

func (g *goGenerator) GenerateCode(jsonMap map[string]interface{}) (string, error) {
	fmt.Println("Generating code...")
	fmt.Printf("Received JSON map: %v\n", jsonMap)
	return "", nil
}
