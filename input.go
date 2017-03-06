package dpipe

// Input interface must be implemented by
// all input sources
type Input interface {
	// Get receives gatherer
	// reades input data
	// iterates and write data into gatherer
	Read(Gatherer) error
}
