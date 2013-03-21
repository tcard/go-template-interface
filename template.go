// Package template provides a common interface for `text/template#Template` and `html/template#Template`. 
// Wrappers for both types which implement this interface are at `/html` and `/text`.
//
// Objects of types which implement this interface must care about its underlying type only
// when they are created. So we can turn this:
//
// 	import (
// 		textTemplate "text/template"
// 		htmlTemplate "html/template"
// 	)
//
// 	// ...
//
// 	if weWantHTML {
// 		tpl, err := htmlTemplate.New("myTemplate").Parse(myHTML)
// 		// ... do lots of stuff with tpl ...
// 		tpl.Execute(wr, commonData)
// 	} else if weWantText {
// 		tpl, err := textTemplate.New("myTemplate").Parse(myText)
// 		// ... do lots of stuff with tpl ...
// 		tpl.Execute(wr, commonData)
// 	}
//
// 	// ...
//
// Into this:
//
// 	import (
// 		textTemplate "github.com/tcard/go-template-interface/text"
// 		htmlTemplate "github.com/tcard/go-template-interface/html"
// 		"github.com/tcard/go-template-interface"
// 	)
//
// 	// ...
//
// 	var tpl template.Template
// 	if weWantHTML {
// 		tpl, err = htmlTemplate.New("myTemplate").Parse(myHTML)
// 	} else if weWantText {
// 		tpl, err = textTemplate.New("myTemplate").Parse(myText)
// 	}
// 	// ... do lots of stuff with tpl ...
// 	tpl.Execute(wr, commonData)
//
// 	// ...
package template

import (
	"io"
	"text/template/parse"
)

type FuncMap map[string]interface{}

type Template interface {
	AddParseTree(name string, tree *parse.Tree) (Template, error)
	Clone() (Template, error)
	Delims(left, right string) Template
	Execute(wr io.Writer, data interface{}) (err error)
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	Funcs(funcMap FuncMap) Template
	Lookup(name string) Template
	Name() string
	New(name string) Template
	Parse(text string) (Template, error)
	ParseFiles(filenames ...string) (Template, error)
	ParseGlob(pattern string) (Template, error)
	Templates() []Template
}

// func Must(t *Template, err error) *Template {
// 	switch t.Type {
// 	case HTML:
// 		return &Template{Type: HTML, HtmlTemplate: htmlTemplate.Must(t.HtmlTemplate, err)}
// 	case Text:
// 		fallthrough
// 	default:
// 		return &Template{Type: Text, TextTemplate: textTemplate.Must(t.HtmlTemplate, err)}
// 	}
// }

// func New(tplType uint, name string) *Template {
// 	switch tplType {
// 	case HTML:
// 		return &Template{Type: HTML, HtmlTemplate: htmlTemplate.New(name)}
// 	case Text:
// 		fallthrough
// 	default:
// 		return &Template{Type: Text, TextTemplate: textTemplate.New(name)}
// 	}
// }

// func ParseFiles(tplType uint, filenames ...string) (*Template, error) {
// 	switch tplType {
// 	case HTML:
// 		return &Template{Type: HTML, HtmlTemplate: htmlTemplate.ParseFiles(filenames...)}
// 	case Text:
// 		fallthrough
// 	default:
// 		return &Template{Type: Text, TextTemplate: textTemplate.ParseFiles(filenames...)}
// 	}
// }

// func ParseGlob(tplType uint, pattern string) (*Template, error) {
// 	switch tplType {
// 	case HTML:
// 		return &Template{Type: HTML, HtmlTemplate: htmlTemplate.ParseGlob(pattern)}
// 	case Text:
// 		fallthrough
// 	default:
// 		return &Template{Type: Text, TextTemplate: textTemplate.ParseGlob(pattern)}
// 	}
// }

// func (t *Template) AddParseTree(name string, tree *parse.Tree) (*Template, error) {
// 	var err error
// 	switch t.Type {
// 	case HTML:
// 		t.HtmlTemplate, err = t.HtmlTemplate.AddParseTree(name, tree)
// 	case Text:
// 		fallthrough
// 	default:
// 		t.TextTemplate, err = t.TextTemplate.AddParseTree(name, tree)
// 	}
// 	return t, err
// }

// func (t *Template) Clone() (*Template, error) {
// 	var err error
// 	switch t.Type {
// 	case HTML:
// 		t.HtmlTemplate, err = t.HtmlTemplate.Clone()
// 	case Text:
// 		fallthrough
// 	default:
// 		t.TextTemplate, err = t.TextTemplate.Clone()

// 	}
// 	return t, err
// }

// func (t *Template) Delims(left, right string) *Template {
// 	switch t.Type {
// 	case HTML:
// 		t.HtmlTemplate = t.HtmlTemplate.Delims(left, right)
// 	case Text:
// 		fallthrough
// 	default:
// 		t.TextTemplate = t.TextTemplate.Delims(left, right)
// 	}
// 	return t
// }

// func (t *Template) Execute(wr io.Writer, data interface{}) (err error) {
// 	switch t.Type {
// 	case HTML:
// 		err = t.HtmlTemplate.Execute(wr, data)
// 	case Text:
// 		fallthrough
// 	default:
// 		err = t.TextTemplate.Execute(wr, data)
// 	}
// 	return
// }

// func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
// 	switch t.Type {
// 	case HTML:
// 		return t.HtmlTemplate.Execute(wr, name, data)
// 	case Text:
// 		fallthrough
// 	default:
// 		return t.TextTemplate.Execute(wr, name, data)
// 	}
// }

// func (t *Template) Funcs(funcMap FuncMap) *Template {
// 	switch t.Type {
// 	case HTML:
// 		t.HtmlTemplate = t.HtmlTemplate.Funcs(funcMap)
// 	case Text:
// 		fallthrough
// 	default:
// 		t.TextTemplate = t.TextTemplate.Funcs(funcMap)
// 	}
// 	return t
// }

// func (t *Template) Lookup(name string) *Template {
// 	switch t.Type {
// 	case HTML:
// 		t.HtmlTemplate = t.HtmlTemplate.Lookup(name)
// 	case Text:
// 		fallthrough
// 	default:
// 		t.TextTemplate = t.TextTemplate.Lookup(name)
// 	}
// 	return t
// }

// func (t *Template) Name() string {
// 	switch t.Type {
// 	case HTML:
// 		return t.HtmlTemplate.Name()
// 	case Text:
// 		fallthrough
// 	default:
// 		return t.TextTemplate.Name()
// 	}
// }

// func (t *Template) New(name string) *Template {
// 	switch t.Type {
// 	case HTML:
// 		t.HtmlTemplate = t.HtmlTemplate.New(name)
// 	case Text:
// 		fallthrough
// 	default:
// 		t.TextTemplate = t.TextTemplate.New(name)
// 	}
// 	return t
// }

// func (t *Template) Parse(text string) (*Template, error) {
// 	var err error
// 	switch t.Type {
// 	case HTML:
// 		t.HtmlTemplate, err = t.HtmlTemplate.Parse(text)
// 	case Text:
// 		fallthrough
// 	default:
// 		t.TextTemplate, err = t.TextTemplate.Parse(text)
// 	}
// 	return t, err
// }

// func (t *Template) ParseFiles(filenames ...string) (*Template, error) {
// 	var err error
// 	switch t.Type {
// 	case HTML:
// 		t.HtmlTemplate, err = t.HtmlTemplate.ParseFiles(filenames...)
// 	case Text:
// 		fallthrough
// 	default:
// 		t.TextTemplate, err = t.TextTemplate.ParseFiles(filenames...)
// 	}
// 	return t, err
// }

// func (t *Template) ParseGlob(pattern string) (*Template, error) {
// 	var err error
// 	switch t.Type {
// 	case HTML:
// 		t.HtmlTemplate, err = t.HtmlTemplate.ParseGlob(pattern)
// 	case Text:
// 		fallthrough
// 	default:
// 		t.TextTemplate, err = t.TextTemplate.ParseGlob(pattern)
// 	}
// 	return t, err
// }

// func (t *Template) Templates() []*Template {
// 	var retTpls []*Template
// 	switch t.Type {
// 	case HTML:
// 		tpls := t.HtmlTemplate.Templates()
// 		retTpls = make([]*Template, len(tpls))
// 		for i, v := range tpls {
// 			retTpls[i] = v
// 		}
// 	case Text:
// 		fallthrough
// 	default:
// 		tpls := t.TextTemplate.Templates()
// 		retTpls = make([]*Template, len(tpls))
// 		for i, v := range tpls {
// 			retTpls[i] = v
// 		}
// 	}
// 	return retTpls
// }
