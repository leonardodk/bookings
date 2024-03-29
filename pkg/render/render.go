package render

//this is just for rendering templates

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/leonardodk/bookings/pkg/config"
	"github.com/leonardodk/bookings/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)

	return td
}

// Renders a template
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
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

	td = AddDefaultData(td, r)

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

	//get all of the files named *.page.tmpl
	// these are the names of the files that will build the pages
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

		//return a new template and name it after the filename
		ts := template.New(name)

		// Parsefiles parses the file(s) and associates them with the template Parsefiles is called on.
		// could have said ts, err := template.New(name).Parsefiles(page) as it does it all in one step

		ts, err = ts.ParseFiles(page) // build the template

		if err != nil {
			return myCache, err
		}

		// check to see if there are any layout template files
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
		// create an entry in the slice myCache "name" and associate the template ts with it
		myCache[name] = ts
	}

	return myCache, nil
}
