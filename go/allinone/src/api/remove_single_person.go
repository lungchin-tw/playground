package api

import (
	"fmt"
	"net/http"
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

	w.WriteHeader(http.StatusOK)
}
