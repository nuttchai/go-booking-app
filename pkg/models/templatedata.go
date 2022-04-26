package models

// NOTE: If we create its models as a sparate file, it can avoid IMPORT CYCLE problem

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} // NOTE: when we are not sure about its datatype, we can pass its type as interface{}
	CSRFToken string                 // NOTE: It is a security token that is used to handle the forum post; it called cross site request forgery token
	Flash     string                 // NOTE: message to the user
	Warning   string
	Error     string
}
