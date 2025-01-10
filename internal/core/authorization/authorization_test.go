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

	authorizationData1, err := token.Parse(authorizationKey)

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

	authorizationData2, err := token.Parse(authorizationKey2)

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

	authorizationData3, err := token.Parse(authorizationKey3)

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

func TestHasAccessTokenCheck(t *testing.T) {
	hasAccess1 := token.HasAccess(authorizationKey, sections.WAITING_LIST, actions.ADMIN)

	if !hasAccess1 {
		t.Fatal("authorizationKey should have access to waitin list admin access")
	}

	hasAccess2 := token.HasAccess(authorizationKey2, sections.PERFORMANCE, actions.VIEW)

	if hasAccess2 {
		t.Fatal("authorizationKey2 should not have access to performance view")
	}

	hasAccess3 := token.HasAccess(authorizationKey3, sections.RIDERS, actions.EDIT)

	if !hasAccess3 {
		t.Fatal("authorizationKey3 should have access to rides edit access")
	}
}

func TestHasAccessGroupTokenCheck(t *testing.T) {
	hasAccessGroupAll1 := token.HasAccessGroupAll(authorizationKey, []token.Access{
		{
			Section: sections.RIDERS,
			Action:  actions.EDIT,
		},
		{
			Section: sections.WAITING_LIST,
			Action:  actions.EXTRA,
		},
		{
			Section: sections.INTERVIEW,
			Action:  actions.ADD,
		},
	})

	if !hasAccessGroupAll1 {
		t.Fatal("authorizationKey should have access to hasAccessGroupAll1")
	}

	hasAccessGroupAll2 := token.HasAccessGroupAll(authorizationKey, []token.Access{
		{
			Section: sections.RIDERS,
			Action:  actions.EDIT,
		},
		{
			Section: sections.PERFORMANCE,
			Action:  actions.ADD,
		},
	})

	if hasAccessGroupAll2 {
		t.Fatal("authorizationKey3 should not have access to hasAccessGroupAll2")
	}

	hasAccessGroupAny1 := token.HasAccessGroupAny(authorizationKey, []token.Access{
		{
			Section: sections.WAREHOUSE,
			Action:  actions.VIEW,
		},
		{
			Section: sections.PAYROLL,
			Action:  actions.UPLOAD,
		},
		{
			Section: sections.INTERVIEW,
			Action:  actions.ADD,
		},
	})

	if !hasAccessGroupAny1 {
		t.Fatal("authorizationKey should have access to hasAccessGroupAny1")
	}

	hasAccessGroupAny2 := token.HasAccessGroupAny(authorizationKey3, []token.Access{
		{
			Section: sections.WAREHOUSE,
			Action:  actions.VIEW,
		},
		{
			Section: sections.PAYROLL,
			Action:  actions.UPLOAD,
		},
	})

	if hasAccessGroupAny2 {
		t.Fatal("authorizationKey3 should not have access to hasAccessGroupAny2")
	}
}
