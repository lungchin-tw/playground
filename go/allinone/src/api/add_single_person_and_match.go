package api

import (
	"fmt"
	"net/http"
	"playground/allinone/model"
	"playground/allinone/util"
)

func AddSinglePersonAndMatch(w http.ResponseWriter, r *http.Request) {
	fmt.Println(util.CurFuncDesc())
	fmt.Println("\t Query:", r.URL.RawQuery)

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		err := fmt.Errorf("Invalid Method:%v", r.Method)
		fmt.Fprintln(w, err.Error())
		return
	}

	user_name := model.GetUserFromURL(r.URL)
	height, err := model.GetHeightFromURL(r.URL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err.Error())
		return
	}

	gender, err := model.GetGenderFromURL(r.URL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err.Error())
		return
	}

	numDates, err := model.GetNumDatesFromURL(r.URL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err.Error())
		return
	}

	source := model.NewUser(
		user_name,
		height,
		gender,
		numDates,
	)

	fmt.Println("\t Source:", source)

	matchService := model.DefaultMatchService()
	matchService.AddUser(source)
	if target := matchService.RandomMatch(source); target != nil {
		w.WriteHeader(http.StatusOK)
		matchService.DecreaseNumDatesAndRemove(source.Name())
		matchService.DecreaseNumDatesAndRemove(target.Name())
		fmt.Fprintln(w, target)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Can't find any available person")
	}
}
