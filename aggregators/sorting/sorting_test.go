package sorting

import (
	"testing"

	"github.com/maksadbek/dpipe"
	"github.com/stretchr/testify/require"
)

func TestSorting(t *testing.T) {
	testHotels := []dpipe.Hotel{
		dpipe.Hotel{
			Name:    "The Gibson",
			Address: "63847 Lowe Knoll, East Maxine, WA 97030-4876",
			Stars:   1,
			Contact: "Dr. Sinda Wyman",
			Phone:   "1-270-665-9933x1626",
			URI:     "http://thegibson.com",
		},
		dpipe.Hotel{
			Name:    "Martini Cattaneo",
			Address: "Stretto Bernardi 004, Quarto Mietta nell'emilia, 07958 Torino (OG)",
			Stars:   1,
			Contact: "Rosalino Marchetti",
			Phone:   "+39 627 68225719",
			URI:     "http://www.farina.org/blog/categories/tags/about.html",
		},
		dpipe.Hotel{
			Name:    "Apartment Dörr",
			Address: "Bolzmannweg 451, 05116 Hannover",
			Stars:   2,
			Contact: "Scarlet Kusch-Linke",
			Phone:   "08177354570",
			URI:     "http://www.garden.com/list/home.html",
		},
		dpipe.Hotel{
			Name:    "Henck Schleich",
			Address: "Jesselstraße 31, 82544 Rochlitz",
			Stars:   2,
			Contact: "Klarissa Etzold",
			Phone:   "(09891) 58482",
			URI:     "http://reichmann.de/main/",
		},
		dpipe.Hotel{
			Name:    "The Rolland",
			Address: "56, chemin de Bertin, 02035 Gros",
			Stars:   3,
			Contact: "Paulette Maury",
			Phone:   "+33 1 39 31 91 77",
			URI:     "http://www.rousseau.fr/",
		},
		dpipe.Hotel{
			Name:    "Lagarde Comfort Inn",
			Address: "76, rue Célina Durand, 68 773 BoutinVille",
			Stars:   4,
			Contact: "Alex Henry",
			Phone:   "+33 3 60 79 55 08",
			URI:     "http://the.com/register/",
		},
		dpipe.Hotel{
			Name:    "Diaz",
			Address: "41, avenue de Marin, 91 255 Morvan",
			Stars:   5,
			Contact: "Clémence Hoarau",
			Phone:   "0602634745",
			URI:     "http://vaillant.com/list/app/faq/",
		},
	}

	s := Sorting{}

	for _, h := range testHotels {
		s.Add(h)
	}

	// sort and check hotels by stars
	expectedStars := []int{1, 1, 2, 2, 3, 4, 5}
	sortedHotels, err := s.Do("stars")
	require.NoError(t, err)

	sortedStars := []int{}
	for _, h := range sortedHotels {
		sortedStars = append(sortedStars, h.Stars)
	}
	require.Equal(t, expectedStars, sortedStars)

	// sort and check hotels by name
	expectedNames := []string{
		"Apartment Dörr",
		"Diaz",
		"Henck Schleich",
		"Lagarde Comfort Inn",
		"Martini Cattaneo",
		"The Gibson",
		"The Rolland",
	}
	sortedHotels, err = s.Do("name")
	require.NoError(t, err)

	sortedNames := []string{}
	for _, h := range sortedHotels {
		sortedNames = append(sortedNames, h.Name)
	}

	require.Equal(t, expectedNames, sortedNames)
}
