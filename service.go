package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

const DATA_DIR = "data"
const DB_FILE = "data/api-urls.json"

func ServiceCreateDoc(id string, version string, name string, file io.Reader, ext string) error {
	// fix short extensions
	if ext == ".yml" {
		ext = ".yaml"
	}
	// copy to storage
	os.Mkdir(filepath.Join(DATA_DIR, id), 0700)
	filename := fmt.Sprintf("%s%s", version, ext)
	filename = filepath.Join(DATA_DIR, id, filename)
	dest, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer dest.Close()
	_, err = io.Copy(dest, file)
	if err != nil {
		return err
	}
	// copy to index
	url := fmt.Sprintf("<<baseurl>>/api/v1/docs/%s/%s", id, version)
	entry := ApiUrl{
		Url:     url,
		Name:    name,
		Id:      id,
		Version: version,
	}

	data, err := ServiceLoadDb()
	if err != nil {
		return err
	}
	exists := false
	for i := 0; i < len(data); i++ {
		if data[i].Id == id && data[i].Version == version {
			data[i] = entry
			exists = true
			break
		}
	}
	if !exists {
		data = append(data, entry)
	}
	return ServiceSaveDb(data)
}

func ServiceDeleteDoc(id string, version string) error {
	// remove file
	name := fmt.Sprintf("%s.json", version)
	name = filepath.Join(DATA_DIR, id, name)
	if file_exists(name) {
		err := os.Remove(name)
		if err != nil {
			return err
		}
	}
	name = fmt.Sprintf("%s.yaml", version)
	name = filepath.Join(DATA_DIR, id, name)
	err := os.Remove(name)
	if err != nil {
		return err
	}

	// load the index
	data, err := ServiceLoadDb()
	if err != nil {
		return err
	}
	// remove from the index and save it back
	for i := 0; i < len(data); i++ {
		if data[i].Id == id && data[i].Version == version {
			data = remove(data, i)
			ServiceSaveDb(data)
			break
		}
	}
	return nil
}

func ServiceGetDoc(id string, version string) (string, []byte, error) {
	name := fmt.Sprintf("%s.json", version)
	name = filepath.Join(DATA_DIR, id, name)
	if file_exists(name) {
		data, err := ioutil.ReadFile(name)
		return "application/json; charset=UTF-8", data, err
	}
	name = fmt.Sprintf("%s.yaml", version)
	name = filepath.Join(DATA_DIR, id, name)
	data, err := ioutil.ReadFile(name)
	return "text/vnd.yaml; charset=UTF-8", data, err
}

func ServiceLoadDb() ([]ApiUrl, error) {
	data := make([]ApiUrl, 0)
	file, err := ioutil.ReadFile(DB_FILE)
	if err == nil {
		err = json.Unmarshal([]byte(file), &data)
	}
	return data, err
}

func ServiceSaveDb(data []ApiUrl) error {
	// sort the api first
	sort.Slice(data, func(i, j int) bool {
		return data[i].Name < data[j].Name
	})
	// save it
	file, _ := json.MarshalIndent(data, "", " ")
	return ioutil.WriteFile(DB_FILE, file, 0644)
}

func remove(s []ApiUrl, i int) []ApiUrl {
	s[i] = s[0]
	return s[1:]
}

func file_exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
