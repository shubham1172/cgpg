package generators

type tplData struct {
	Name string
	Data map[string]interface{}
}

type Generator interface {
	// GetName returns the name of the generator.
	GetName() string

	// GenerateCode generates code from the given JSON map.
	// The returned string is the generated code.
	GenerateCode(jsonMap map[string]interface{}) (map[string]string, error)
}
