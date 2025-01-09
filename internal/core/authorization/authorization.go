package authorization

import (
	"core-api-go/internal/core/authorization/access"
	"core-api-go/internal/core/authorization/actions"
	"strconv"
)

type AccessData struct {
	Action        string
	AccessType    int
	AccessDataIds []int
}

type AuthorizationMap map[string][]*actions.Action

func (auth AuthorizationMap) SetAccess(section string, accesses []AccessData) {
	_, ok := auth[section]

	if !ok {
		auth[section] = make([]*actions.Action, 0, 1)
	}

	for _, data := range accesses {
		auth[section] = append(auth[section], actions.New(data.Action, access.New(data.AccessType, data.AccessDataIds)))
	}
}

func (auth AuthorizationMap) GenerateToken() string {
	tokenIndex := ""
	tokenBody := ""

	remain := len(auth)

	for section := range auth {
		remain--

		bodyLen := len(tokenBody)

		if bodyLen > 0 {
			tokenIndex += strconv.Itoa(bodyLen)

			if remain > 0 {
				tokenIndex += ","
			}
		}

		tokenBody += section

		for _, action := range auth[section] {
			tokenBody += action.GenerateKey()
		}
	}

	return tokenIndex + "|" + tokenBody
}

func New() AuthorizationMap {
	return make(AuthorizationMap)
}
