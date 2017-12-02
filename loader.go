package vast

import "html/template"

// TemplateLoader is an interface that provides template data to a vast Model.
// Template loaders should produce a single output template comprising all
// of the data they can load.
type TemplateLoader interface {
	LoadTemplate(baseTemplate *template.Template) (*template.Template, error)
}

type Glob string

func (g Glob) LoadTemplate(baseTemplate *template.Template) (*template.Template, error) {
	return baseTemplate.ParseGlob(string(g))
}

type Files []string

func (f Files) LoadTemplate(baseTemplate *template.Template) (*template.Template, error) {
	return baseTemplate.ParseFiles(f...)
}

type Strings []string

func (s Strings) LoadTemplate(baseTemplate *template.Template) (*template.Template, error) {
	var err error
	for _, v := range s {
		baseTemplate, err = baseTemplate.Parse(v)
		if err != nil {
			break
		}
	}
	return baseTemplate, err
}
