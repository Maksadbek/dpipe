package dpipe

// Gatherer receives Hotel struct
// from inputs
type Gatherer interface {
	Write(Hotel)
}
