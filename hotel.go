package dpipe

// Hotel wraps hotel info
type Hotel struct {
	Name    string
	Address string
	Stars   int
	Contact string
	Phone   string
	URI     string
}

// GetFieldValue can be used to
// Hotel structs fields by name
func (h Hotel) GetFieldValue(field string) interface{} {
	switch field {
	case "name":
		return h.Name
	case "address":
		return h.Address
	case "stars":
		return h.Stars
	case "contact":
		return h.Contact
	case "phone":
		return h.Phone
	case "uri":
		return h.URI
	default:
		return nil
	}
}
