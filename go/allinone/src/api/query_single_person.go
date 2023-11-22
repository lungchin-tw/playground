package api

import (
	"fmt"
	"net/http"
	"playground/allinone/model"
	"playground/allinone/util"
)

func QuerySinglePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println(util.CurFuncDesc())
	fmt.Println("\t Query:", r.URL.RawQuery)

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		err := fmt.Errorf("Invalid Method:%v", r.Method)
		fmt.Fprintln(w, err.Error())
		return
	}

	user_name, err := model.GetUserFromURL(r.URL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err.Error())
		return
	}

	result := model.DefaultMatchService().FindPossiblePeopleByName(user_name)
	if len(result) > 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, result)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Can't find any available person")
	}

}
