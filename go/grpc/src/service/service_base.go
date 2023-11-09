package service

import "log"

type ServiceBase struct {
}

func (s *ServiceBase) logError(err error) error {
	if err != nil {
		log.Println(err)
	}

	return err
}
