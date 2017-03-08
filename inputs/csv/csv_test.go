package csv

import (
	"testing"

	"github.com/maksadbek/dpipe"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

type mockGatherer struct {
	expectedHotelsCount int
	hotelsCount         int
}

func (mg *mockGatherer) Write(h dpipe.Hotel) {
	mg.hotelsCount++
}

func TestRead(t *testing.T) {
	mg := &mockGatherer{
		expectedHotelsCount: 8,
	}
	csvParser := &CSV{}

	config := viper.New()
	config.Set("file", "./testdata/hotels.csv")

	csvParser.LoadConf(config)
	err := csvParser.Read(mg)
	require.NoError(t, err)

	require.Equal(t, mg.expectedHotelsCount, mg.hotelsCount)
}
