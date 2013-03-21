// Package textTemplate provides a wrapper type around text/template#Template which implements
// github.com/tcard/go-template-interface#Template.
//
// See text/template for docs on methods and github.com/tcard/go-template-interface for rationale.
package textTemplate

import (
	"github.com/tcard/go-template-interface"
	"io"
	oldTemplate "text/template"
	"text/template/parse"
)

type TextTemplate struct {
	*oldTemplate.Template
}

func Must(t template.Template, err error) template.Template {
	return TextTemplate{oldTemplate.Must(t.(TextTemplate).Template, err)}
}

func New(name string) template.Template {
	return TextTemplate{oldTemplate.New(name)}
}

func ParseFiles(filenames ...string) (template.Template, error) {
	t, err := oldTemplate.ParseFiles(filenames...)
	return TextTemplate{t}, err
}

func ParseGlob(pattern string) (template.Template, error) {
	t, err := oldTemplate.ParseGlob(pattern)
	return TextTemplate{t}, err
}

func (t TextTemplate) AddParseTree(name string, tree *parse.Tree) (template.Template, error) {
	ret, err := t.Template.AddParseTree(name, tree)
	return TextTemplate{ret}, err
}

func (t TextTemplate) Clone() (template.Template, error) {
	ret, err := t.Template.Clone()
	return TextTemplate{ret}, err
}

func (t TextTemplate) Delims(left, right string) template.Template {
	ret := t.Template.Delims(left, right)
	return TextTemplate{ret}
}

func (t TextTemplate) Execute(wr io.Writer, data interface{}) (err error) {
	return t.Template.Execute(wr, data)
}

func (t TextTemplate) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return t.Template.ExecuteTemplate(wr, name, data)
}

func (t TextTemplate) Funcs(funcMap template.FuncMap) template.Template {
	ret := t.Template.Funcs(oldTemplate.FuncMap(funcMap))
	return TextTemplate{ret}
}

func (t TextTemplate) Lookup(name string) template.Template {
	ret := t.Template.Lookup(name)
	return TextTemplate{ret}
}

func (t TextTemplate) Name() string {
	return t.Template.Name()
}

func (t TextTemplate) New(name string) template.Template {
	ret := t.Template.New(name)
	return TextTemplate{ret}
}

func (t TextTemplate) Parse(text string) (template.Template, error) {
	ret, err := t.Template.Parse(text)
	return TextTemplate{ret}, err
}

func (t TextTemplate) ParseFiles(filenames ...string) (template.Template, error) {
	ret, err := t.Template.ParseFiles(filenames...)
	return TextTemplate{ret}, err
}

func (t TextTemplate) ParseGlob(pattern string) (template.Template, error) {
	ret, err := t.Template.ParseGlob(pattern)
	return TextTemplate{ret}, err
}

func (t TextTemplate) Templates() []template.Template {
	oldTpls := t.Template.Templates()
	tpls := make([]template.Template, len(oldTpls))
	for i, v := range oldTpls {
		tpls[i] = TextTemplate{v}
	}
	return tpls
}
