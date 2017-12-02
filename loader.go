package vast

import "html/template"

// TemplateLoader is an interface that provides template data to a vast Model.
// Template loaders should produce a single output template comprising all
// of the data they can load.
type TemplateLoader interface {
	LoadTemplate(baseTemplate *template.Template) (*template.Template, error)
}

// Glob is a TemplateLoader that loads templates from a set of files named by a glob
// wildcard.
type Glob string

// LoadTemplate exists to satisfy the TemplateLoader interface.
func (g Glob) LoadTemplate(baseTemplate *template.Template) (*template.Template, error) {
	return baseTemplate.ParseGlob(string(g))
}

// Files is a TemplateLoader that loads templates from a list of files.
// Wildcards are not supported.
type Files []string

// LoadTemplate exists to satisfy the TemplateLoader interface.
func (f Files) LoadTemplate(baseTemplate *template.Template) (*template.Template, error) {
	return baseTemplate.ParseFiles(f...)
}

// Strings is a TemplateLoader that loads templates from a set of strings.
type Strings []string

// LoadTemplate exists to satisfy the TemplateLoader interface.
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
