package generators

import (
	"bytes"
	"fmt"
	"text/template"
)

type goGenerator struct{}

var goTemplateText = `
package generated

type {{.Name}} struct {
	Code string
	Message string
}

func (e *{{.Name}}) GetMessage() string {
	return e.Message
}

func (e *{{.Name}}) GetCode() string {
	return e.Code
}

var (
	{{range $key, $value := .Data}}
	{{$key}} = &{{$.Name}}{
		Code: "{{$key}}",
		Message: "{{$value}}",
	}
	{{end}}
)`

func NewGoGenerator() Generator {
	return &goGenerator{}
}

func (g *goGenerator) GetName() string {
	return "go"
}

func (g *goGenerator) GenerateCode(jsonMap map[string]interface{}) (map[string]string, error) {
	ts := make([]tplData, 0)
	files := make(map[string]string)

	for fileName, fileContent := range jsonMap {
		fileContentMap, ok := fileContent.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid JSON: %s", fileName)
		}

		ts = append(ts, tplData{
			Name: fileName,
			Data: fileContentMap,
		})
	}

	tmpl, err := template.New("go").Parse(goTemplateText)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	for _, t := range ts {
		buf := new(bytes.Buffer)
		err = tmpl.Execute(buf, t)
		if err != nil {
			return nil, fmt.Errorf("failed to execute template: %w", err)
		}

		files[t.Name+".go"] = buf.String()
	}

	return files, nil
}
