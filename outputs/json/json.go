package json

import (
	"encoding/json"
	"io"
	"os"
	"sync"

	"github.com/maksadbek/dpipe"
	"github.com/maksadbek/dpipe/outputs"
	"github.com/spf13/viper"
)

// JSON implements dpipe.Output interface
// reads dpipe.Hotel data
type JSON struct {
	filePath string
	file     *os.File
	once     *sync.Once
	encoder  *json.Encoder
}

type taggedHotel struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Stars   int    `json:"stars"`
	Contact string `json:"contact"`
	Phone   string `json:"phone"`
	URI     string `json:"uri"`
}

func (j *JSON) Write(h dpipe.Hotel) error {
	hotel := taggedHotel{
		Name:    h.Name,
		Address: h.Address,
		Stars:   h.Stars,
		Contact: h.Contact,
		Phone:   h.Phone,
		URI:     h.URI,
	}

	j.encoder.Encode(hotel)
	j.file.WriteString(",")

	return nil
}

func (j *JSON) LoadConf(conf *viper.Viper) error {
	j.filePath = conf.GetString("file")
	return j.open()
}

func (j *JSON) Close() error {
	// set the offset to last element
	// it is ',' symbol
	// the next write replaces it
	j.file.Seek(-1, io.SeekCurrent)
	j.file.WriteString("]")
	return j.file.Close()
}

// open creates a file and json encoder
func (j *JSON) open() error {
	var err error

	j.file, err = os.Create(j.filePath)
	if err != nil {
		return err
	}

	j.file.WriteString("[")
	j.encoder = json.NewEncoder(j.file)
	return nil
}

func init() {
	outputs.Add("json", &JSON{})
}
