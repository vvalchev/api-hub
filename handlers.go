package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
)

var VERSION_RE = regexp.MustCompile("^(?:(\\d+)\\.)?(?:(\\d+)\\.)?(\\*|\\d+)(-\\w+)?$")
var ID_RE = regexp.MustCompile("^[a-z]+([a-z_\\-])?[a-z]+")

// CreateDoc - Registers a new API with the system.
func CreateDoc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	id := r.MultipartForm.Value["id"][0]
	version := r.MultipartForm.Value["version"][0]
	name := r.MultipartForm.Value["name"][0]
	file, header, err := r.FormFile("document")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	defer file.Close()
	ext := filepath.Ext(header.Filename)

	// validate parameters
	if !validate(id, version) {
		http.Error(w, "Invalid parameters", 400)
		return
	}

	// create the document
	err = ServiceCreateDoc(id, version, name, file, ext)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	ok(&w, true)
}

// DeleteDoc - Gets the Swagger/OpenAPI for a particular entry.
func DeleteDoc(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// validate parameters
	id := params["id"]
	ver := params["version"]
	if !validate(id, ver) {
		http.Error(w, "Invalid parameters", 400)
		return
	}

	// delete the documnet
	err := ServiceDeleteDoc(id, ver)
	if err == nil {
		ok(&w, true)
	} else {
		http.NotFound(w, r)
	}
}

// GetDoc - Gets the Swagger/OpenAPI for a particular entry.
func GetDoc(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// validate parameters
	id := params["id"]
	ver := params["version"]
	if !validate(id, ver) {
		http.Error(w, "Invalid parameters", 400)
		return
	}

	// get the document and server it
	mime, file, err := ServiceGetDoc(id, ver)
	if err == nil {
		// calculare baseurl & replace in data file
		baseurl := fmt.Sprintf("http://%s", r.Host)
		if r.TLS != nil {
			baseurl = fmt.Sprintf("https://%s", r.Host)
		}
		data := strings.Replace(string(file), "<<baseurl>>", baseurl, -1)
		// write file
		w.Header().Set("Content-Type", mime)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(data))
	} else {
		http.NotFound(w, r)
	}
}

// GetDocs - Retrieve a summary of available documentation
func GetDocs(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(DB_FILE)
	if err == nil {
		defer file.Close()
		ok(&w, false)
		io.Copy(w, file)
	} else {
		http.NotFound(w, r)
	}
}

func validate(id string, version string) bool {
	ok := len(id) <= 30 && ID_RE.MatchString(id)
	return ok && len(version) < 256 && VERSION_RE.MatchString(version)
}

func ok(w *http.ResponseWriter, empty bool) {
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
	if empty {
		(*w).Write([]byte("{}"))
	}
}
