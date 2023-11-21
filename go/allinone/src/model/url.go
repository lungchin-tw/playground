package model

import (
	"fmt"
	"net/url"
	"playground/allinone/config"
	"strconv"
)

func BuildURL(
	host string,
	user string,
	height int,
	gender EnumGender,
	numdates int,
) string {
	return fmt.Sprintf(
		"%v?%v=%v&%v=%v&%v=%v&%v=%v",
		host,
		config.PARAM_USER, user,
		config.PARAM_HEIGHT, height,
		config.PARAM_GENDER, gender,
		config.PARAM_NUMDATES, numdates,
	)
}

func GetUserFromURL(url *url.URL) string {
	return url.Query().Get(config.PARAM_USER)
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
