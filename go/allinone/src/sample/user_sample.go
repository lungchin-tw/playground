package sample

import "playground/allinone/model"

func GetSampleMaleUsers() []*model.User {
	return []*model.User{
		model.NewUser("male01", 170, model.EG_MALE, 1),
		model.NewUser("male02", 175, model.EG_MALE, 2),
		model.NewUser("male03", 180, model.EG_MALE, 3),
		model.NewUser("male04", 185, model.EG_MALE, 4),
		model.NewUser("male05", 190, model.EG_MALE, 5),
	}
}

func GetSampleFemaleUsers() []*model.User {
	return []*model.User{
		model.NewUser("female01", 160, model.EG_FEMALE, 1),
		model.NewUser("female02", 165, model.EG_FEMALE, 2),
		model.NewUser("female03", 170, model.EG_FEMALE, 3),
		model.NewUser("female04", 175, model.EG_FEMALE, 4),
		model.NewUser("female05", 180, model.EG_FEMALE, 5),
		model.NewUser("female06", 190, model.EG_FEMALE, 6),
	}
}
