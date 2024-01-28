package models

import "github.com/mneumantic/wilkinson-innovations/internal/forms"

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
	Form       *forms.Form
}

type Product struct {
	Id      string
	ImgFile string
	Caption string
}

type Testimonial struct {
	Name    string
	ImgFile string
	ImgAlt  string
	Quote   string
	Hidden  bool
}

type Contact struct {
	Name    string
	Title   string
	Company string
	Email   string
	Phone   string
	Message string
}
