package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

type Photo struct {
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
}

type PhotoResponse struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Photos  []Photo `json:"photos"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/html/index.html")
}

func getJS(w http.ResponseWriter, r *http.Request) {
	filename := strings.Split(r.URL.Path, "/")[2]
	http.ServeFile(w, r, "web/js/"+filename)
}

func getCSS(w http.ResponseWriter, r *http.Request) {
	filename := strings.Split(r.URL.Path, "/")[2]
	http.ServeFile(w, r, "web/css/"+filename)
}

func getImg(w http.ResponseWriter, r *http.Request) {
	filename := strings.Split(r.URL.Path, "/")[2]
	http.ServeFile(w, r, "web/img/"+filename)
}

// From https://stackoverflow.com/questions/55300117/how-do-i-find-all-files-that-have-a-certain-extension-in-go-regardless-of-depth
func WalkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func getPhotoList(w http.ResponseWriter, r *http.Request) {
	var ret PhotoResponse

	err := filepath.Walk("./images/photos/",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err)
				return err
			}
			var nextPhoto Photo
			nextPhoto.Filename = info.Name()
			nextPhoto.Size = info.Size()
			ret.Photos = append(ret.Photos, nextPhoto)
			return nil
		})
	
	files, err := WalkMatch("./images/photos", "*.jpeg")

	if err != nil {
		error_string := fmt.Sprint(err)
		ret.Status = "fail"
		ret.Message = error_string
	} else {
		for _, f in range(files){
			var nextPhoto Photo
			nextPhoto.Filename = info.Name()
			nextPhoto.Size = info.Size()
			ret.Photos = append(ret.Photos, nextPhoto)
		}
		ret.Status = "ok"
	}

	json.NewEncoder(w).Encode(&ret)
}

func RunWebserver() {
	// Setup and run all routes
	fmt.Println("Starting web server")

	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/js/{filename}", getJS)
	r.HandleFunc("/css/{filename}", getCSS)
	r.HandleFunc("/img/{filename}", getImg)

	r.HandleFunc("/api/photos", getPhotoList)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
