package render

import (
	"bytes"
	"github.com/mneumantic/wilkinson-innovations/internal/config"
	"github.com/mneumantic/wilkinson-innovations/internal/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.Template, r *http.Request) *models.Template {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.Template) {
	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	_ = t.Execute(buf, td)

	// render template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// get all of the pages
	pages, err := filepath.Glob("./templates/pages/*.gohtml")

	if err != nil {
		return cache, err
	}

	// range throught pages
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("./templates/layouts/*.gohtml")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/layouts/*.gohtml")
			if err != nil {
				return cache, err
			}
		}

		partials, err := filepath.Glob("./templates/partials/*.gohtml")
		if err != nil {
			return cache, err
		}

		if len(partials) > 0 {
			ts, err = ts.ParseGlob("./templates/partials/*.gohtml")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = ts
	}

	return cache, nil
}
