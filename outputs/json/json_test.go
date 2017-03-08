package json

import (
	jsonEncoder "encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/maksadbek/dpipe"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestWrite(t *testing.T) {
	conf := viper.New()
	conf.Set("file", "hotels.json")

	json := &JSON{}

	err := json.LoadConf(conf)
	require.NoError(t, err)

	defer os.Remove(conf.GetString("file"))

	expectedHotels := []dpipe.Hotel{
		dpipe.Hotel{
			Name:    "Evrard Breton",
			Address: "23, rue Stéphanie Legrand, 07 124 Leconte ",
			Stars:   1,
			Contact: "Roland Bazin",
			Phone:   "04 01 64 99 70",
			URI:     "http://www.begue.fr/search/register",
		},
	}

	for _, h := range expectedHotels {
		err = json.Write(h)
		require.NoError(t, err)
	}
	json.Close()

	c, err := ioutil.ReadFile(conf.GetString("file"))
	require.NoError(t, err)

	currentHotels := []taggedHotel{}

	err = jsonEncoder.Unmarshal(c, &currentHotels)
	require.NoError(t, err)

	require.Equal(t, len(expectedHotels), len(currentHotels))

}
