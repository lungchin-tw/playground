package api

import (
	"fmt"
	"net/http"
	"playground/allinone/util"
)

func RemoveSinglePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println(util.CurFuncDesc())
	fmt.Println("\t Query:", r.URL.RawQuery)

	w.WriteHeader(http.StatusOK)
}
