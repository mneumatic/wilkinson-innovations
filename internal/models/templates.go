package models

type Template struct {
	StringMap  map[string]string
	IntMap     map[string]int
	FloatMap   map[string]float32
	Data       map[string]interface{}
	CSRFToken  string
	Flash      string
	Warning    string
	Error      string
	Title      string
	Production bool
	Other      interface{}
	Products   interface{}
}
