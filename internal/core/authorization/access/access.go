package access

import "strconv"

const (
	OWN_DATA int = iota
	OWN_CITY
	CITY_GROUP
	ALL_CITIES
	ACCESS_GROUP
)

type Access struct {
	name string
	key  int
	ids  []int
}

func (a *Access) GenerateKey() string {
	key := strconv.Itoa(a.key)

	for i := range a.ids {
		if i > 0 {
			key += "."
		}

		key += strconv.Itoa(a.ids[i])
	}

	return key
}

func New(key int, ids []int) *Access {
	name := "none"

	switch key {
	case 0:
		name = "own-data"
	case 1:
		name = "own-city"
	case 2:
		name = "city-group"
	case 3:
		name = "all-cities"
	case 4:
		name = "access-group"
	}

	return &Access{
		name: name,
		key:  key,
		ids:  ids,
	}
}
