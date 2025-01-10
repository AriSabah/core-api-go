package actions

import (
	"core-api-go/internal/core/authorization/access"
)

const (
	VIEW       = "V"
	ADD        = "A"
	EDIT       = "E"
	DELETE     = "D"
	DOWNLOAD   = "W"
	UPLOAD     = "U"
	ADMIN      = "M"
	STATISTICS = "S"
	EXTRA      = "X"
)

type Action struct {
	name   string
	key    string
	access *access.Access
}

func (a *Action) GenerateKey() string {
	if a.key == "V" {
		return a.access.GenerateKey()
	}

	return a.key + a.access.GenerateKey()
}

func getNameFromKey(key string) string {
	name := "none"

	switch key {
	case "V":
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
		key = "V"
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
