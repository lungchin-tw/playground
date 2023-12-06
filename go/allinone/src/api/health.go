package api

import (
	"fmt"
	"net/http"
	"playground/allinone/util"
)

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println(util.CurFuncDesc())
	fmt.Println("\t Request:", r)

	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Healthy")
	} else {
		w.WriteHeader(http.StatusBadRequest)
		err := fmt.Errorf("Invalid Method:%v", r.Method)
		fmt.Fprintln(w, err.Error())
	}
}
