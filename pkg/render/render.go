package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
	"html/template"
	"github.com/AMSteffensen/bookings/pkg/config"
	"github.com/AMSteffensen/bookings/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders a template using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	var err error

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Fatal(err)
		}
	}

	// get the template cache from app config
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := make(map[string]*template.Template)

	// get all of the files named *.page.tmpl
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
