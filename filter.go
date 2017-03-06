package dpipe

// Filter checks value of field
// and returns True/False if it passes validation
type Filter interface {
	Validate(string) bool
}
