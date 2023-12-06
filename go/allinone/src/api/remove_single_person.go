package api

import (
	"fmt"
	"net/http"
	"playground/allinone/model"
	"playground/allinone/util"
)

func RemoveSinglePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println(util.CurFuncDesc())
	fmt.Println("\t Query:", r.URL.RawQuery)

	if r.Method != http.MethodPost {
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

	model.DefaultMatchService().DeleteUser(user_name)
	w.WriteHeader(http.StatusOK)
}
