package util

import (
	"embed"
	"os"
	"text/template"
)

var TemplateFs embed.FS

func PopulateFiles(path string, templateName string, templatePath string, data interface{}) {
	file, _ := os.Create(path)

	tmpl := template.New(templateName).Delims("[[", "]]")

	tmpl, _ = tmpl.ParseFS(TemplateFs, "templates/"+templatePath+"/"+templateName)
	tmpl.Execute(file, data)

	file.Close()
}
