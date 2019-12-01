package handlefunc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
)

//Page struct is used for HTML templates
type Page struct {
	Name       string
	VideoNames []string
	Index      int
}

var myClient = http.Client{Timeout: 10 * time.Second}
var tmpls = template.Must(template.ParseFiles("tmpl/index.html", "tmpl/upload.html", "tmpl/watch.html", "tmpl/information.html"))

//TemplateInit used every where to initialize HTML templates
func TemplateInit(w http.ResponseWriter, templateFile string, templateData Page) {
	if err := tmpls.ExecuteTemplate(w, templateFile, templateData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//Index is the main page of the web app
func Index(w http.ResponseWriter, r *http.Request) {

	page := GetVideoList()
	TemplateInit(w, "index.html", page)
}

//Upload files handler
func Upload(w http.ResponseWriter, r *http.Request) {
	page := Page{Name: "Upload Page"}
	TemplateInit(w, "upload.html", page)
}

//Watch handles video playing
func Watch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in Watch")
	indexOfVideo := r.URL.Path[len("/watch/"):]
	indexOfVideoInt, _ := strconv.Atoi(indexOfVideo)
	listOfVideos := GetVideoList()

	page := Page{Name: strings.Replace(listOfVideos.VideoNames[indexOfVideoInt], ".mp4", "", -1), Index: indexOfVideoInt}
	TemplateInit(w, "watch.html", page)
}

//Info displays the Information Page
func Info(w http.ResponseWriter, r *http.Request) {
	page := Page{Name: "Information Page"}
	TemplateInit(w, "information.html", page)
}

//GetVideoList returns a list of videos uploaded on the server
func GetVideoList() Page {
	page := Page{}
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := myClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &page)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return page
}
