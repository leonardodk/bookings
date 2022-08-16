package models

// holds data sent from handlers to template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[int]int
	FloatMap  map[float32]float32
	Data      map[string]interface{} // interface is what you put when you don't know the data type
	CSRFToken string
	Flash     string //for flash messages
	Error     string //for error messages
	Warning   string //for warning messages
}
