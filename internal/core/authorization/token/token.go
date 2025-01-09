package token

import (
	"core-api-go/internal/core/authorization"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type token struct {
	key string
}

func (t *token) Parse() (*authorization.AuthorizationMap, error) {
	authorizationData := authorization.New()

	before, after, found := strings.Cut(t.key, "|")

	if !found {
		return nil, errors.New("invalid token key")
	}

	var indexs []string

	if len(before) > 0 {
		indexs = strings.Split(before, ",")
	} else {
		indexs = make([]string, 0)
	}

	startIndex := 0
	endIndex := 0

	for i := -1; i < len(indexs); i++ {
		if i == -1 {
			startIndex = 0
		} else {
			index, err := strconv.Atoi(indexs[i])

			if err != nil {
				return nil, errors.New("invalid token key, expected int in mapping portion of the token")
			}

			startIndex = index
		}

		if i == len(indexs)-1 {
			endIndex = len(after)
		} else {
			index, err := strconv.Atoi(indexs[i+1])

			if err != nil {
				return nil, errors.New("invalid token key, expected int in mapping portion of the token")
			}

			endIndex = index
		}

		accessData := make([]authorization.AccessData, 0, 1)
		action := ""

		for i := startIndex + 1; i < endIndex; {
			if i != startIndex+1 {
				action = string(after[i])
			}

			accessType, err := strconv.Atoi(string(after[i+1]))

			if err != nil {
				return nil, errors.New(fmt.Sprintf("invalid access type, expected int in %s [%s]", after[i-1:i+2], after))
			}

			indexShift := 2
			var AccessDataIds []int = nil

			if i+2 < len(after) && !unicode.IsUpper(rune(after[i+2])) {
				for f := 0; f < len(after); f++ {
					if i+2+f == len(after) || unicode.IsUpper(rune(after[i+2+f])) {
						indexShift += f
						break
					}
				}

				strIds := strings.Split(after[i+2:i+indexShift], ".")

				for _, strId := range strIds {
					id, err := strconv.Atoi(strId)

					if err != nil {
						return nil, errors.New(fmt.Sprintf("invalid access type, expected int in %s [%s]", after[i-1:i+2], after))
					}

					AccessDataIds = append(AccessDataIds, id)
				}
			}

			accessData = append(accessData, authorization.AccessData{
				Action:        string(action),
				AccessType:    accessType,
				AccessDataIds: AccessDataIds,
			})

			i += indexShift
		}

		authorizationData.SetAccess(after[startIndex:startIndex+2], accessData)
	}

	return &authorizationData, nil
}

func New(key string) *token {
	return &token{key: key}
}
