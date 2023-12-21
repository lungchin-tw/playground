package model

import (
	"fmt"
	"net/url"
	"playground/allinone/config"
	"strconv"
)

func BuildURL_AddSinglePersonAndMatch(
	path string,
	user string,
	height int,
	gender EnumGender,
	numdates int,
) string {
	return fmt.Sprintf(
		"%v?%v=%v&%v=%v&%v=%v&%v=%v",
		path,
		config.PARAM_USER, user,
		config.PARAM_HEIGHT, height,
		config.PARAM_GENDER, gender,
		config.PARAM_NUMDATES, numdates,
	)
}

func BuildURL_QuerySinglePerson(
	host string,
	user string,
) string {
	return fmt.Sprintf(
		"%v?%v=%v",
		host,
		config.PARAM_USER, user,
	)
}

func BuildURL_RemoveSinglePerson(
	host string,
	user string,
) string {
	return fmt.Sprintf(
		"%v?%v=%v",
		host,
		config.PARAM_USER, user,
	)
}

func GetUserFromURL(url *url.URL) (string, error) {
	v := url.Query().Get(config.PARAM_USER)
	if len(v) == 0 {
		return "", fmt.Errorf("User Not Found in Query")
	}

	return v, nil
}

func GetHeightFromURL(url *url.URL) (int, error) {
	v := url.Query().Get(config.PARAM_HEIGHT)
	return strconv.Atoi(v)
}

func GetGenderFromURL(url *url.URL) (EnumGender, error) {
	v := url.Query().Get(config.PARAM_GENDER)
	int_value, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}

	enum_value := EnumGender(int_value)
	return enum_value, enum_value.CheckValid()
}

func GetNumDatesFromURL(url *url.URL) (int, error) {
	v := url.Query().Get(config.PARAM_NUMDATES)
	return strconv.Atoi(v)
}
