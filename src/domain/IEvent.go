package domain

import "httpTools/src/infrastructure/event/vo"

type IEvent interface {
	Emit()
	Register()
	CreateHandler() vo.HFunc
}
