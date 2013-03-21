// Package htmlTemplate provides a wrapper type around html/template#Template which implements
// github.com/tcard/go-template-interface#Template.
//
// See text/template for docs on methods and github.com/tcard/go-template-interface for rationale.
package htmlTemplate

import (
	"github.com/tcard/go-template-interface"
	oldTemplate "html/template"
	"io"
	"text/template/parse"
)

type HtmlTemplate struct {
	*oldTemplate.Template
}

func Must(t template.Template, err error) template.Template {
	return HtmlTemplate{oldTemplate.Must(t.(HtmlTemplate).Template, err)}
}

func New(name string) template.Template {
	return HtmlTemplate{oldTemplate.New(name)}
}

func ParseFiles(filenames ...string) (template.Template, error) {
	t, err := oldTemplate.ParseFiles(filenames...)
	return HtmlTemplate{t}, err
}

func ParseGlob(pattern string) (template.Template, error) {
	t, err := oldTemplate.ParseGlob(pattern)
	return HtmlTemplate{t}, err
}

func (t HtmlTemplate) AddParseTree(name string, tree *parse.Tree) (template.Template, error) {
	ret, err := t.Template.AddParseTree(name, tree)
	return HtmlTemplate{ret}, err
}

func (t HtmlTemplate) Clone() (template.Template, error) {
	ret, err := t.Template.Clone()
	return HtmlTemplate{ret}, err
}

func (t HtmlTemplate) Delims(left, right string) template.Template {
	ret := t.Template.Delims(left, right)
	return HtmlTemplate{ret}
}

func (t HtmlTemplate) Execute(wr io.Writer, data interface{}) (err error) {
	return t.Template.Execute(wr, data)
}

func (t HtmlTemplate) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return t.Template.ExecuteTemplate(wr, name, data)
}

func (t HtmlTemplate) Funcs(funcMap template.FuncMap) template.Template {
	ret := t.Template.Funcs(oldTemplate.FuncMap(funcMap))
	return HtmlTemplate{ret}
}

func (t HtmlTemplate) Lookup(name string) template.Template {
	ret := t.Template.Lookup(name)
	return HtmlTemplate{ret}
}

func (t HtmlTemplate) Name() string {
	return t.Template.Name()
}

func (t HtmlTemplate) New(name string) template.Template {
	ret := t.Template.New(name)
	return HtmlTemplate{ret}
}

func (t HtmlTemplate) Parse(text string) (template.Template, error) {
	ret, err := t.Template.Parse(text)
	return HtmlTemplate{ret}, err
}

func (t HtmlTemplate) ParseFiles(filenames ...string) (template.Template, error) {
	ret, err := t.Template.ParseFiles(filenames...)
	return HtmlTemplate{ret}, err
}

func (t HtmlTemplate) ParseGlob(pattern string) (template.Template, error) {
	ret, err := t.Template.ParseGlob(pattern)
	return HtmlTemplate{ret}, err
}

func (t HtmlTemplate) Templates() []template.Template {
	oldTpls := t.Template.Templates()
	tpls := make([]template.Template, len(oldTpls))
	for i, v := range oldTpls {
		tpls[i] = HtmlTemplate{v}
	}
	return tpls
}
