package api

import (
	"fmt"
	"html/template"
	"net/http"
	"playground/allinone/model"
	"playground/allinone/util"
)

func AddSinglePersonAndMatch(w http.ResponseWriter, r *http.Request) {
	fmt.Println(util.CurFuncDesc())
	fmt.Println("\t Query:", r.URL.RawQuery)

	user_name := model.GetUserFromURL(r.URL)
	height, err := model.GetHeightFromURL(r.URL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, template.HTMLEscapeString(err.Error()))
		return
	}

	gender, err := model.GetGenderFromURL(r.URL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, template.HTMLEscapeString(err.Error()))
		return
	}

	numDates, err := model.GetNumDatesFromURL(r.URL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, template.HTMLEscapeString(err.Error()))
		return
	}

	source := model.NewUser(
		user_name,
		height,
		gender,
		numDates,
	)

	fmt.Println("\t Source:", source)

	model.DefaultMatchService().AddUser(source)
	w.WriteHeader(http.StatusOK)
	if target := model.DefaultMatchService().RandomMatch(source); target != nil {
		fmt.Fprintln(w, target)
	} else {
		fmt.Fprintln(w, "Can't find any available person")
	}
}
