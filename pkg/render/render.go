package render

//this is just for rendering templates

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/leonardodk/bookings/pkg/config"
	"github.com/leonardodk/bookings/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// Renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// create a template cache
	//get the template cache from app.config
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("Could not get template from template cache...")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)

	if err != nil {
		log.Println(err)
	}
	//render the teamplate
	_, err = buf.WriteTo(w)

	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	// could also do make(map[string]*template.Template)

	//get all of the files named *page.tmpl
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	// element number and value
	// here we loop through slice of strings returned by Glob, we name it and then check for layouts
	// if there are any layouts we add them to the page template
	// this is then added to the template cache via its name
	for _, page := range pages {
		// Base gets us the last elemeent of the filepath ie the name
		name := filepath.Base(page)
		//parse the template and then name it after itself
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// check to see if there are any layouts
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		// if there are any layouts then set the template to the template but also with layouts added
		// basically x += template
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
