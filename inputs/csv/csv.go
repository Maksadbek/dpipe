package csv

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/maksadbek/dpipe"
	"github.com/maksadbek/dpipe/inputs"
	"github.com/spf13/viper"
)

// CSV reads csv data from file
// creates a dpipe.Hotel data
// writes the hotel data into gatherer
type CSV struct {
	file string // file path
}

func (c *CSV) Read(g dpipe.Gatherer) error {
	f, err := os.Open(c.file)
	if err != nil {
		return err
	}

	r := csv.NewReader(f)

	// read headers
	_, err = r.Read()
	if err != nil {
		return err
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		stars, _ := strconv.Atoi(record[2])

		h := dpipe.Hotel{
			Name:    record[0],
			Address: record[1],
			Stars:   stars,
			Contact: record[3],
			Phone:   record[4],
			URI:     record[5],
		}

		g.Write(h)
	}

	return nil
}

func (c *CSV) LoadConf(v *viper.Viper) {
	// get filename from config
	c.file = v.GetString("file")
}

func init() {
	inputs.Add("csv", &CSV{})
}
