package service

import (
	""
)

type (
	User type {

	}
)

var (
	Users = make(map[uint32]User, 256)
)