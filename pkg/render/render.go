package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/nuttchai/WebApp-Golang/pkg/config"
	"github.com/nuttchai/WebApp-Golang/pkg/models"
)

// IMPORTANT: all the files for a package must exist in the same directory

var functions = template.FuncMap{}
var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

/* NOTE: All the files in the same directory are the same package
So, we need to name RenderTemplate as capital letter to export the function outside the package */

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		// NOTE: Newer Version, reading template from memory
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// NOTE: ok variable tells whether the key is found in the given map or not
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	// NOTE: we are going to put the parsed template that is in memory into some bytes
	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	_ = t.Execute(buf, td)
	// NOTE: this would write to response writer, and that should write everything to the template
	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

	/* Old Version, reading template from disk
	// NOTE: template.ParseFiles will parse the HTML file to be rendered
	parseTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parseTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	} */
}

// CreateTemplateCache is used to parse all of out templates, including the layouts, and store them in the variable, myCache
// It creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	/* NOTE:
	key: string in this case is a name of the template like home.page.html
	value: it is a pointer to that ready to use template which is used to render the site */
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			// NOTE: ParseGlob returns a value that you're throwing away, but you need to keep it; it's the template object that you need to call
			ts, err = ts.ParseGlob("./templates/*.layout.html")

			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
