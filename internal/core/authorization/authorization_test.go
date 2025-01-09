package authorization_test

import (
	"core-api-go/internal/core/authorization"
	"core-api-go/internal/core/authorization/access"
	"core-api-go/internal/core/authorization/actions"
	"core-api-go/internal/core/authorization/sections"
	"core-api-go/internal/core/authorization/token"
	"testing"
)

const (
	authorizationKey  = "10,17,45|RI3A21.2E1LE1E1X3WL3A21.2E21.2D0W0U0M21.2S3X3IN21.2A1E41.3"
	authorizationKey2 = "29,34|RI41151.1152A41174.1151E41151LE3E1WL21.2"
	authorizationKey3 = "|RI1A1E1"
)

func TestGenerateAuthenticationToken(t *testing.T) {
	authorizationData1 := authorization.New()

	authorizationData1.SetAccess(sections.RIDERS, []authorization.AccessData{
		{
			Action:        actions.VIEW,
			AccessType:    access.ALL_CITIES,
			AccessDataIds: nil,
		},
		{
			Action:        actions.ADD,
			AccessType:    access.CITY_GROUP,
			AccessDataIds: []int{1, 2},
		},
		{
			Action:        actions.EDIT,
			AccessType:    access.OWN_CITY,
			AccessDataIds: nil,
		},
	})

	authorizationData1.SetAccess(sections.LEAVES, []authorization.AccessData{
		{
			Action:        actions.VIEW,
			AccessType:    access.OWN_CITY,
			AccessDataIds: nil,
		},
		{
			Action:        actions.EDIT,
			AccessType:    access.OWN_CITY,
			AccessDataIds: nil,
		},
		{
			Action:        actions.EXTRA,
			AccessType:    access.ALL_CITIES,
			AccessDataIds: nil,
		},
	})

	authorizationData1.SetAccess(sections.WAITING_LIST, []authorization.AccessData{
		{
			Action:        actions.VIEW,
			AccessType:    access.ALL_CITIES,
			AccessDataIds: nil,
		},
		{
			Action:        actions.ADD,
			AccessType:    access.CITY_GROUP,
			AccessDataIds: []int{1, 2},
		},
		{
			Action:        actions.EDIT,
			AccessType:    access.CITY_GROUP,
			AccessDataIds: []int{1, 2},
		},
		{
			Action:        actions.DELETE,
			AccessType:    access.OWN_DATA,
			AccessDataIds: nil,
		},
		{
			Action:        actions.DOWNLOAD,
			AccessType:    access.OWN_DATA,
			AccessDataIds: nil,
		},
		{
			Action:        actions.UPLOAD,
			AccessType:    access.OWN_DATA,
			AccessDataIds: nil,
		},
		{
			Action:        actions.ADMIN,
			AccessType:    access.CITY_GROUP,
			AccessDataIds: []int{1, 2},
		},
		{
			Action:        actions.STATISTICS,
			AccessType:    access.ALL_CITIES,
			AccessDataIds: nil,
		},
		{
			Action:        actions.EXTRA,
			AccessType:    access.ALL_CITIES,
			AccessDataIds: nil,
		},
	})

	authorizationData1.SetAccess(sections.INTERVIEW, []authorization.AccessData{
		{
			Action:        actions.VIEW,
			AccessType:    access.CITY_GROUP,
			AccessDataIds: []int{1, 2},
		},
		{
			Action:        actions.ADD,
			AccessType:    access.OWN_CITY,
			AccessDataIds: nil,
		},
		{
			Action:        actions.EDIT,
			AccessType:    access.ACCESS_GROUP,
			AccessDataIds: []int{1, 3},
		},
	})

	generatedToken := authorizationData1.GenerateToken()

	if generatedToken != authorizationKey {
		t.Fatalf(`
			The generated token from authorizationData1 is not correct.
			generated token: %s
			expected token:  %s
		`, generatedToken, authorizationKey)
	}

	authorizationData2 := authorization.New()

	authorizationData2.SetAccess(sections.RIDERS, []authorization.AccessData{
		{
			Action:        actions.VIEW,
			AccessType:    access.ACCESS_GROUP,
			AccessDataIds: []int{1151, 1152},
		},
		{
			Action:        actions.ADD,
			AccessType:    access.ACCESS_GROUP,
			AccessDataIds: []int{1174, 1151},
		},
		{
			Action:        actions.EDIT,
			AccessType:    access.ACCESS_GROUP,
			AccessDataIds: []int{1151},
		},
	})

	authorizationData2.SetAccess(sections.LEAVES, []authorization.AccessData{
		{
			Action:        actions.VIEW,
			AccessType:    access.ALL_CITIES,
			AccessDataIds: nil,
		},
		{
			Action:        actions.EDIT,
			AccessType:    access.OWN_CITY,
			AccessDataIds: nil,
		},
	})

	authorizationData2.SetAccess(sections.WAITING_LIST, []authorization.AccessData{
		{
			Action:        actions.VIEW,
			AccessType:    access.CITY_GROUP,
			AccessDataIds: []int{1, 2},
		},
	})

	generatedToken = authorizationData2.GenerateToken()

	if generatedToken != authorizationKey2 {
		t.Fatalf(`
			The generated token from authorizationData2 is not correct.
			generated token: %s
			expected token:  %s
		`, generatedToken, authorizationKey2)
	}

	authorizationData3 := authorization.New()

	authorizationData3.SetAccess(sections.RIDERS, []authorization.AccessData{
		{
			Action:        actions.VIEW,
			AccessType:    access.OWN_CITY,
			AccessDataIds: nil,
		},
		{
			Action:        actions.ADD,
			AccessType:    access.OWN_CITY,
			AccessDataIds: nil,
		},
		{
			Action:        actions.EDIT,
			AccessType:    access.OWN_CITY,
			AccessDataIds: nil,
		},
	})

	generatedToken = authorizationData3.GenerateToken()

	if generatedToken != authorizationKey3 {
		t.Fatalf(`
			The generated token from authorizationData3 is not correct.
			generated token: %s
			expected token:  %s
		`, generatedToken, authorizationKey3)
	}
}

func TestParseAuthenticationToken(t *testing.T) {
	token1 := token.New(authorizationKey)

	authorizationData1, err := token1.Parse()

	if err != nil {
		t.Fatalf("error while parsing token1, error: %s", err.Error())
	}

	generatedAuthorizationKey := authorizationData1.GenerateToken()

	if generatedAuthorizationKey != authorizationKey {
		t.Fatalf(`
			The parsed token from authorizationKey is not correct.
			generated token: %s
			expected token:  %s
		`, generatedAuthorizationKey, authorizationKey)
	}

	token2 := token.New(authorizationKey2)

	authorizationData2, err := token2.Parse()

	if err != nil {
		t.Fatalf("error while parsing token2, error: %s", err.Error())
	}

	generatedAuthorizationKey2 := authorizationData2.GenerateToken()

	if generatedAuthorizationKey2 != authorizationKey2 {
		t.Fatalf(`
			The parsed token from authorizationKey is not correct.
			generated token: %s
			expected token:  %s
		`, generatedAuthorizationKey2, authorizationKey2)
	}

	token3 := token.New(authorizationKey3)

	authorizationData3, err := token3.Parse()

	if err != nil {
		t.Fatalf("error while parsing token3, error: %s", err.Error())
	}

	generatedAuthorizationKey3 := authorizationData3.GenerateToken()

	if generatedAuthorizationKey3 != authorizationKey3 {
		t.Fatalf(`
			The parsed token from authorizationKey is not correct.
			generated token: %s
			expected token:  %s
		`, generatedAuthorizationKey3, authorizationKey3)
	}
}
