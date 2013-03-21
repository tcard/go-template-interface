	import "github.com/tcard/go-template-interface"
	import "github.com/tcard/go-template-interface/html"
	import "github.com/tcard/go-template-interface/text"

Package template provides a common interface for `text/template#Template` and `html/template#Template`. 
Wrappers for both types which implement this interface are at `/html` and `/text`.

Objects of types which implement this interface must care about its underlying type only
when they are created. So we can turn this:

	import (
		textTemplate "text/template"
		htmlTemplate "html/template"
	)

	// ...

	if weWantHTML {
		tpl, err := htmlTemplate.New("myTemplate").Parse(myHTML)
		// ... do lots of stuff with tpl ...
		tpl.Execute(wr, commonData)
	} else if weWantText {
		tpl, err := textTemplate.New("myTemplate").Parse(myText)
		// ... do lots of stuff with tpl ...
		tpl.Execute(wr, commonData)
	}

	// ...

Into this:

	import (
		textTemplate "github.com/tcard/go-template-interface/text"
		htmlTemplate "github.com/tcard/go-template-interface/html"
		"github.com/tcard/go-template-interface"
	)

	// ...

	var tpl template.Template
	if weWantHTML {
		tpl, err = htmlTemplate.New("myTemplate").Parse(myHTML)
	} else if weWantText {
		tpl, err = textTemplate.New("myTemplate").Parse(myText)
	}
	// ... do lots of stuff with tpl ...
	tpl.Execute(wr, commonData)

	// ...
