package document

import (
	"strings"

	"github.com/matthewmueller/golly/js"
)

// Element struct
// js:"element,omit"
type Element struct {
	nodeType int
	nodeName string
}

// Body is a reference to the document body
var Body = js.Rewrite("document.body").(*Element)

// NodeName gets the node name
func (e *Element) NodeName() string {
	js.Rewrite("$<.nodeName", e)
	return e.nodeName
}

// InnerHTML fn
// TODO: go implementation
func (e *Element) InnerHTML() string {
	js.Rewrite("$<.innerHTML")
	// lower := strings.ToLower(e.nodeName)
	return ""
}

// OuterHTML fn
// TODO: go implementation
func (e *Element) OuterHTML() string {
	js.Rewrite("$<.outerHTML")
	lower := strings.ToLower(e.nodeName)
	return "<" + lower + ">" + "</" + lower + ">"
}

// CreateElement creates an element
func CreateElement(element string) *Element {
	js.Rewrite("document.createElement($1)", element)
	return &Element{
		nodeType: 1,
		nodeName: strings.ToUpper(element),
	}
}
