package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Shreeyash-Naik/Hotel-Booking/pkg/config"
	"github.com/Shreeyash-Naik/Hotel-Booking/pkg/models"
)

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.html")
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
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}

func RenderTemplate(w *http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Printf("%s template does not exist", tmpl)
	}

	td = AddDefaultData(td)
	t.Execute(*w, td)

	// parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	// if err := parsedTemplate.Execute(*w, parsedTemplate); err != nil {
	// 	fmt.Println("error parsing template: ", err)
	// 	return
	// }
}
