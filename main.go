package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// Get app folder
func getAppDir() string {
	// Get current folder or die
	dir, patherr := filepath.Abs(filepath.Dir(os.Args[0]))
	if patherr != nil {
		log.Fatal(patherr)
	}
	return dir
}

// Settings
var (
	port       = "8080"
	dir        = getAppDir()
	staticpath = dir + "/static/"
	imgpath    = staticpath + "img/"
	csspath    = staticpath + "css/"
)

// Uncomment here and comment indexTemplate at index() to cache template
// indexTemplate is the HTML template we use to present the index page.
//var (
//	indexTemplate = template.Must(template.ParseFiles("index.min.html"))
//)

// index serves the index page
func index(w http.ResponseWriter, r *http.Request) *appError {
	// This check prevents the "/" handler from handling all requests by default
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return nil
	}

	// Fill in the missing fields in index.html
	var data = struct {
		Title, PageHeader string
	}{"MyTitle", "Bootstrap starter template"}

	// Uncomment here when you want cache.
	// remember to uncomment indexTemplate above.
	indexTemplate := template.Must(template.ParseFiles("index.min.html"))

	// Render and serve the HTML
	err := indexTemplate.Execute(w, data)
	if err != nil {
		log.Println("error rendering template:", err)
		return &appError{err, "Error rendering template", 500}
	}
	return nil
}

// appHandler is to be used in error handling
type appHandler func(http.ResponseWriter, *http.Request) *appError

type appError struct {
	Err     error
	Message string
	Code    int
}

// serveHTTP formats and passes up an error
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		log.Println(e.Err)
		http.Error(w, e.Message, e.Code)
	}
}

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

func handleFavicon(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("static", "favicon.ico")
	http.ServeFile(w, r, fp)
}

// Here everything starts
func main() {
	// Serve the index.html page
	http.Handle("/", appHandler(index))

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticpath))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir(imgpath))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir(csspath))))

	serveSingle("/favicon.ico", "./static/favicon.ico")

	// Starting server
	fmt.Println("Starting server on port: " + port + "...")
	err := http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
