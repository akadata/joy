package window

import "github.com/matthewmueller/golly/js"

var _ Event = (*MessageEvent)(nil)

// NewMessageEvent fn
func NewMessageEvent(typearg string, eventinitdict *MessageEventInit) *MessageEvent {
	js.Rewrite("MessageEvent")
	return &MessageEvent{}
}

// MessageEvent struct
// js:"MessageEvent,omit"
type MessageEvent struct {
}

// InitMessageEvent fn
// js:"initMessageEvent"
func (*MessageEvent) InitMessageEvent(typeArg string, canBubbleArg bool, cancelableArg bool, dataArg interface{}, originArg string, lastEventIdArg string, sourceArg *Window) {
	js.Rewrite("$_.initMessageEvent($1, $2, $3, $4, $5, $6, $7)", typeArg, canBubbleArg, cancelableArg, dataArg, originArg, lastEventIdArg, sourceArg)
}

// InitEvent fn
// js:"initEvent"
func (*MessageEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*MessageEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*MessageEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*MessageEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// Data prop
// js:"data"
func (*MessageEvent) Data() (data interface{}) {
	js.Rewrite("$_.data")
	return data
}

// Origin prop
// js:"origin"
func (*MessageEvent) Origin() (origin string) {
	js.Rewrite("$_.origin")
	return origin
}

// Ports prop
// js:"ports"
func (*MessageEvent) Ports() (ports interface{}) {
	js.Rewrite("$_.ports")
	return ports
}

// Source prop
// js:"source"
func (*MessageEvent) Source() (source *Window) {
	js.Rewrite("$_.source")
	return source
}

// Bubbles prop
// js:"bubbles"
func (*MessageEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*MessageEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*MessageEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*MessageEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*MessageEvent) CurrentTarget() (currentTarget EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*MessageEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*MessageEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*MessageEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*MessageEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*MessageEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*MessageEvent) SrcElement() (srcElement Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*MessageEvent) Target() (target EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*MessageEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*MessageEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
