package sorting

import (
	"errors"
	"sort"

	"github.com/maksadbek/dpipe"
	"github.com/maksadbek/dpipe/aggregators"
)

// Sorting provides sorting aggregation for hotels data
type Sorting struct {
	data []dpipe.Hotel
	by   func(h1, h2 *dpipe.Hotel) bool
}

// sorting functions by each field
var (
	byName = func(h1, h2 *dpipe.Hotel) bool {
		return h1.Name < h2.Name
	}

	byStars = func(h1, h2 *dpipe.Hotel) bool {
		return h1.Stars < h2.Stars
	}

	byPhone = func(h1, h2 *dpipe.Hotel) bool {
		return h1.Phone < h2.Phone
	}
)

func (s *Sorting) Add(h dpipe.Hotel) error {
	s.data = append(s.data, h)
	return nil
}

// Do makes a sort by given field
// then returns slice of hotels
func (s *Sorting) Do(field string) ([]dpipe.Hotel, error) {
	switch field {
	case "name":
		s.by = byName
		sort.Sort(s)
	case "stars":
		s.by = byStars
		sort.Sort(s)
	case "phone":
		s.by = byPhone
		sort.Sort(s)
	default:
		return nil, errors.New("unsupported field to sort")
	}

	return s.data, nil
}

// implement sorting methods
func (s *Sorting) Len() int {
	return len(s.data)
}

func (s *Sorting) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}

func (s *Sorting) Less(i, j int) bool {
	return s.by(&s.data[i], &s.data[j])
}

func init() {
	aggregators.Add("sorting", &Sorting{})
}
