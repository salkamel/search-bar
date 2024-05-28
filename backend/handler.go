package backend

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if !Exist("homePage.html") {
		w.WriteHeader(http.StatusInternalServerError)
		tmplt := template.New("error template")
		tmplt, _ = tmplt.Parse("Internal server error page missing: {{.}}")
		tmplt.Execute(w, "500")
		return
	}
	switch r.URL.Path {
	case "/":
		t, _ := template.ParseFiles("frontend/html/homePage.html")
		t.Execute(w, artists)
	default:
		errPage := ErrorPage{ErrStatus: `Error 404`, ErrMsg: `Page Not Found`}
		t, _ := template.ParseFiles("frontend/html/errorPage.html")
		w.WriteHeader(404)
		t.Execute(w, errPage)
		return
	}
}

func ArtistsPage(w http.ResponseWriter, r *http.Request) {
	if !Exist("artist.html") {
		w.WriteHeader(http.StatusInternalServerError)
		tmplt := template.New("error template")
		tmplt, _ = tmplt.Parse("Internal server error page missing: {{.}}")
		tmplt.Execute(w, "500")
		return
	}
	rawURL := strings.TrimSuffix(r.URL.Path, "/")
	if strings.Count(rawURL, "/") != 2 { fmt.Println("Invalid URL: ", rawURL)
		errPage := ErrorPage{ErrStatus: `Error 404`, ErrMsg: `Invalid URL`}
		t, _ := template.ParseFiles("frontend/html/errorPage.html")
		w.WriteHeader(404)
		t.Execute(w, errPage)
		return
	} else {
		_, artistId := path.Split(rawURL); // fmt.Printf("artistId = %v\n", artistId)
		artistIndex, _ := strconv.Atoi(artistId); artistIndex = artistIndex - 1
		if artistIndex >= 0 && artistIndex < len(artists) {
			t, _ := template.ParseFiles("frontend/html/artist.html")
			t.Execute(w, artists[artistIndex])
		} else {
			fmt.Println("Invalid artist index: ", artistIndex)
			errPage := ErrorPage{ErrStatus: `Error 404`, ErrMsg: `Invalid artist ID`}
			t, _ := template.ParseFiles("frontend/html/errorPage.html")
			w.WriteHeader(404)
			t.Execute(w, errPage)
		}
	}
}

func Exist(str string) bool {
	path := "frontend/html/" + str
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
