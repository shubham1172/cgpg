package generators

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

type pythonGenerator struct{}

var pythonTemplateText = `
class {{.Name}}:
{{range $key, $value := .Data}}
	{{- $key | indent | ToUpper}} = "{{$value}}"
{{end}}
`

func NewPythonGenerator() Generator {
	return &pythonGenerator{}
}

func (g *pythonGenerator) GetName() string {
	return "python"
}

func (g *pythonGenerator) GenerateCode(jsonMap map[string]interface{}) (map[string]string, error) {
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

	funcMap := template.FuncMap{
		"ToUpper": strings.ToUpper,
		"indent": func(s string) string {
			return "\t" + s
		},
	}
	tmpl, err := template.New("python").Funcs(funcMap).Parse(pythonTemplateText)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	for _, t := range ts {
		buf := new(bytes.Buffer)
		err = tmpl.Execute(buf, t)
		if err != nil {
			return nil, fmt.Errorf("failed to execute template: %w", err)
		}

		files[t.Name+".py"] = buf.String()
	}

	return files, nil
}
