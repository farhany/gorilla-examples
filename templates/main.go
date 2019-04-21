package main

import (
	"fmt"
	"net/http"

	"text/template"

	"github.com/gorilla/mux"
)

type servers struct {
	Name string
	IP   string
}

// Render a template given a model
func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	t := template.Must(template.ParseGlob(tmpl))
	// t := template.Must()
	err := t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	s := []servers{
		{Name: "ubu1", IP: "192.168.4.139"},
		{Name: "ubu2", IP: "192.168.4.133"},
		{Name: "ubu3", IP: "192.168.4.113"},
	}

	for _, v := range s {
		fmt.Println(v.Name, v.IP)
	}
	varMap := map[string]interface{}{
		"Title":   "Home",
		"Servers": s,
	}
	renderTemplate(w, "templates/servers.html", &varMap)
}

func main() {
	router := mux.NewRouter().StrictSlash(false)
	router.Path("/").HandlerFunc(homeHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}
