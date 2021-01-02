package domain

import "httpTools/src/infrastructure/event/vo"

type IEvent interface {
	Register()
	CreateHandler() vo.HFunc
}
