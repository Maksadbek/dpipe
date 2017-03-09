package xml

import (
	"encoding/xml"
	"os"

	"github.com/maksadbek/dpipe"
	"github.com/maksadbek/dpipe/outputs"
	"github.com/spf13/viper"
)

// XML implements dpipe.Output interface
type XML struct {
	outputFile string
	file       *os.File
	encoder    *xml.Encoder
}

type taggedHotel struct {
	XMLName xml.Name `xml:"hotel"`
	Name    string   `xml:"name"`
	Address string   `xml:"address"`
	Stars   int      `xml:"stars"`
	Contact string   `xml:"contact"`
	Phone   string   `xml:"phone"`
	URI     string   `xml:"uri"`
}

func (x *XML) Write(h dpipe.Hotel) error {
	hotel := taggedHotel{
		Name:    h.Name,
		Address: h.Address,
		Stars:   h.Stars,
		Contact: h.Contact,
		Phone:   h.Phone,
		URI:     h.URI,
	}

	return x.encoder.Encode(hotel)
}

func (x *XML) Close() error {
	x.encoder.Indent("", "  ")
	x.file.WriteString("</hotels>")
	return x.file.Close()
}

func (x *XML) LoadConf(v *viper.Viper) error {
	x.outputFile = v.GetString("file")
	return x.open()
}

func (x *XML) open() error {
	var err error

	x.file, err = os.Create(x.outputFile)
	if err != nil {
		return err
	}

	x.encoder = xml.NewEncoder(x.file)
	x.file.WriteString("<hotels>")
	x.encoder.Indent("  ", "  ")
	return nil
}

func init() {
	outputs.Add("xml", &XML{})
}
