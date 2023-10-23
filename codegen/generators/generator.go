package generators

type Generator interface {
	// GenerateCode generates code from the given JSON map.
	// The returned string is the generated code.
	GenerateCode(jsonMap map[string]interface{}) (string, error)
}
