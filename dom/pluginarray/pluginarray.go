package pluginarray

import (
	"github.com/matthewmueller/golly/dom/mimetype"
	"github.com/matthewmueller/golly/js"
)

// PluginArray struct
// js:"PluginArray,omit"
type PluginArray struct {
}

// Item fn
// js:"item"
func (*PluginArray) Item(index uint) (m *mimetype.Plugin) {
	js.Rewrite("$_.item($1)", index)
	return m
}

// NamedItem fn
// js:"namedItem"
func (*PluginArray) NamedItem(name string) (m *mimetype.Plugin) {
	js.Rewrite("$_.namedItem($1)", name)
	return m
}

// Refresh fn
// js:"refresh"
func (*PluginArray) Refresh(reload *bool) {
	js.Rewrite("$_.refresh($1)", reload)
}

// Length prop
// js:"length"
func (*PluginArray) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}
