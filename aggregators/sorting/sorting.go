package sorting

import (
	"errors"
	"sort"

	"github.com/maksadbek/dpipe"
	"github.com/maksadbek/dpipe/aggregators"
)

type sortFunc func(h1, h2 *dpipe.Hotel) bool

// sorting functions by each field
var sortFuncs = map[string]sortFunc{
	"name": func(h1, h2 *dpipe.Hotel) bool {
		return h1.Name < h2.Name
	},

	"stars": func(h1, h2 *dpipe.Hotel) bool {
		return h1.Stars < h2.Stars
	},

	"phone": func(h1, h2 *dpipe.Hotel) bool {
		return h1.Phone < h2.Phone
	},
}

// Sorting provides sorting aggregation for hotels data
type Sorting struct {
	data []dpipe.Hotel
	by   sortFunc
}

func (s *Sorting) Add(h dpipe.Hotel) error {
	s.data = append(s.data, h)
	return nil
}

// Do makes a sort by given field
// then returns slice of hotels
func (s *Sorting) Do(field string) ([]dpipe.Hotel, error) {
	var ok bool
	if s.by, ok = sortFuncs[field]; ok {
		sort.Sort(s)
		return s.data, nil
	} else {
		return nil, errors.New("unsupported field to sort")
	}

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
