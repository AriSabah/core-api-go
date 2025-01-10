package token

import (
	"core-api-go/internal/core/authorization"
	"core-api-go/internal/core/authorization/actions"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func getSection(token string, section string) (sectionToken string, isExist bool) {
	splitterIndex := strings.Index(token, "|")

	if splitterIndex == -1 {
		return "", false
	}

	mapIndex := 0
	startIndex := splitterIndex + 1
	endIndex := 0

	for {
		if mapIndex < splitterIndex {
			nextCommaIndex := strings.Index(token[mapIndex:splitterIndex], ",")

			if nextCommaIndex == -1 {
				nextCommaIndex = splitterIndex - mapIndex
			}

			nextSectionIndex, err := strconv.Atoi(token[mapIndex : mapIndex+nextCommaIndex])

			if err != nil {
				return "", false
			}

			endIndex = splitterIndex + nextSectionIndex + 1
			mapIndex += nextCommaIndex + 1
		} else {
			endIndex = len(token)
		}

		if token[startIndex:startIndex+2] == section {
			return token[startIndex:endIndex], true
		}

		if endIndex == len(token) {
			break
		}

		startIndex = endIndex
	}

	return "", false
}

func getNextLetterIndex(s string, from int) int {
	for i := from; i < len(s); i++ {
		if unicode.IsLetter(rune(s[i])) {
			return i
		}
	}

	return -1
}

func getActionAccess(sectionToken string, action string) (accessType int, accessData []int, isExist bool) {
	currentIndex := 1

	for {
		if currentIndex >= len(sectionToken) {
			break
		}

		nextActionIndex := getNextLetterIndex(sectionToken, currentIndex+1)

		if nextActionIndex == -1 {
			nextActionIndex = len(sectionToken)
		}

		if (currentIndex == 1 && action == actions.VIEW) || string(sectionToken[currentIndex]) == action {
			accessType, err := strconv.Atoi(string(sectionToken[currentIndex+1]))

			if err != nil {
				break
			}

			if currentIndex+2 != nextActionIndex {
				ids := strings.Split(sectionToken[currentIndex+2:nextActionIndex], ".")

				for _, id := range ids {
					numId, err := strconv.Atoi(id)

					if err != nil {
						return -1, nil, false
					}

					accessData = append(accessData, numId)
				}
			}

			return accessType, accessData, true
		}

		currentIndex = nextActionIndex
	}

	return -1, nil, false
}

func HasAccess(token string, section string, action string) bool {
	sectionToken, isExist := getSection(token, section)
	_, _, isExist = getActionAccess(sectionToken, action)

	return isExist
}

type Access struct {
	Section string
	Action  string
}

func HasAccessGroupAll(token string, accesses []Access) bool {
	for _, access := range accesses {
		if !HasAccess(token, access.Section, access.Action) {
			return false
		}
	}

	return true
}

func HasAccessGroupAny(token string, accesses []Access) bool {
	for _, access := range accesses {
		if HasAccess(token, access.Section, access.Action) {
			return true
		}
	}

	return false
}

func hasAccessToData() {}

func Parse(token string) (*authorization.AuthorizationMap, error) {
	authorizationData := authorization.New()

	before, after, found := strings.Cut(token, "|")

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
