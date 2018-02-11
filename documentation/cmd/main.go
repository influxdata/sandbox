package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/shurcooL/github_flavored_markdown"
)

var (
	port     *string
	filePath *string
)

func init() {
	gopath := os.Getenv("GOPATH")
	files := fmt.Sprintf("%v/src/github.com/jackzampolin/sandbox/documentation/", gopath)
	port = flag.String("port", ":3010", "specify port to run server")
	filePath = flag.String("filePath", files, "path where assets live")
	flag.Parse()
}

func mp(add string) string {
	return fmt.Sprintf("%v%v", *filePath, add)
}

func main() {
	log.Printf("Files in path %v\n", *filePath)
	log.Printf("Server started. Listening on port %v\n", *port)
	RunRouter()
}

// RunRouter is the router
func RunRouter() {
	mux := mux.NewRouter()
	mux.StrictSlash(true)
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/tutorials", TutorialIndexHandler)
	mux.HandleFunc("/tutorials/{article:[-a-zA-Z]+}", TemplateHandler).Methods("GET")
	mux.HandleFunc("/healthz", HealthzHandler).Methods("GET")

	// Static files
	mux.PathPrefix("/").Handler(http.FileServer(http.Dir(mp("static/"))))

	http.Handle("/", mux)
	http.ListenAndServe(*port, nil)
}

// ##################
// #    HANDLERS    #
// ##################

// RenderArticle is a helper that renders an article from templates
func RenderArticle(w http.ResponseWriter, r *http.Request, body []byte) {
	fp := mp("static/templates/article.html")
	articleTemplate, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatalf("failed to open %v", fp)
	}
	vars := map[string]interface{}{
		"Header": template.HTML(loadTemplatePart("header")),
		"Hero":   template.HTML(loadTemplatePart("hero")),
		"Body":   template.HTML(body),
		"Footer": template.HTML(loadTemplatePart("footer")),
	}
	t := template.New("article template")
	t, _ = t.Parse(string(articleTemplate))
	t.Execute(w, vars)
}

// TemplateHandler handles the /templates/:id route
func TemplateHandler(w http.ResponseWriter, r *http.Request) {
	requestVars := mux.Vars(r)
	articleName := string(requestVars["article"])
	article := loadArticle(articleName)
	if article == nil {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	RenderArticle(w, r, article.Body)
}

// HealthzHandler handles the /healthz
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

// HomeHandler handles the / route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fp := mp("static/templates/index.html")
	indexTemplate, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatalf("failed to open %v", fp)
	}
	t := template.New("index-template")
	t, _ = t.Parse(string(indexTemplate))
	t.Execute(w, indexTemplate)
}

// TutorialIndexHandler handles the /templates route
func TutorialIndexHandler(w http.ResponseWriter, r *http.Request) {
	article := loadArticle("index")
	RenderArticle(w, r, article.Body)
}

// TODO: write 404 page.
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {

	}
}

// ##################
// #    HELPERS     #
// ##################

// Article is used to build the github_flavored_markdown for pages
type Article struct {
	Body []byte
}

// buildArticlePath is a path helper for articles
func buildArticlePath(fileName string) string {
	return mp(fmt.Sprintf("static/tutorials/%v.md", fileName))
}

// loadArticle pulls the article file from disk and renders the markdown
func loadArticle(fileName string) *Article {
	fp := buildArticlePath(fileName)
	body, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatalf("failed to open %v", fp)
	}
	markdown := github_flavored_markdown.Markdown(body)
	return &Article{Body: markdown}
}

// loadTemplatePart loads the parts of the template into memory
func loadTemplatePart(part string) string {
	fp := mp(fmt.Sprintf("static/templates/%v.html", part))
	content, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatalf("failed to open %v", fp)
	}
	return string(content)
}
