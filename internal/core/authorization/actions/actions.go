package actions

import (
	"core-api-go/internal/core/authorization/access"
)

const (
	VIEW       = "view"
	ADD        = "add"
	EDIT       = "edit"
	DELETE     = "delete"
	DOWNLOAD   = "download"
	UPLOAD     = "upload"
	ADMIN      = "admin"
	STATISTICS = "statistics"
	EXTRA      = "extra"
)

type Action struct {
	name   string
	key    string
	access *access.Access
}

func (a *Action) GenerateKey() string {
	return a.key + a.access.GenerateKey()
}

func getNameFromKey(key string) string {
	name := "none"

	switch key {
	case "":
		name = "view"
	case "A":
		name = "add"
	case "E":
		name = "edit"
	case "D":
		name = "delete"
	case "W":
		name = "download"
	case "U":
		name = "upload"
	case "M":
		name = "admin"
	case "S":
		name = "statistics"
	case "X":
		name = "extra"
	}

	return name
}

func getKeyFromName(name string) string {
	key := "none"

	switch name {
	case "view":
		key = ""
	case "add":
		key = "A"
	case "edit":
		key = "E"
	case "delete":
		key = "D"
	case "download":
		key = "W"
	case "upload":
		key = "U"
	case "admin":
		key = "M"
	case "statistics":
		key = "S"
	case "extra":
		key = "X"
	}

	return key
}

func New(action string, access *access.Access) *Action {
	if len(action) > 1 {
		return &Action{
			name:   action,
			key:    getKeyFromName(action),
			access: access,
		}
	}

	return &Action{
		name:   getNameFromKey(action),
		key:    action,
		access: access,
	}
}
