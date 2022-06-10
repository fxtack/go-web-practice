package templates

import "html/template"

var Template *template.Template

func LoadTemplate(templateFile string) {
	Template = template.Must(template.ParseGlob(templateFile))
}
