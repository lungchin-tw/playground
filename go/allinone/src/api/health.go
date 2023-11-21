package api

import (
	"fmt"
	"html/template"
	"net/http"
	"playground/allinone/util"
)

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println(util.CurFuncDesc())
	fmt.Println("\tRequest:", r)
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Healthy")
	} else {
		w.WriteHeader(http.StatusBadRequest)
		err := fmt.Errorf("Invalid Method:%v", r.Method)
		fmt.Fprintln(w, template.HTMLEscapeString(err.Error()))
	}
}
